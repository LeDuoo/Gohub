package auth

import (
	"Gohub/pkg/config"
	"Gohub/pkg/helpers"
	"Gohub/pkg/logger"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//返回错误方法
func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}

//EncodeBase64Upload 解码base64编码并上传图片
func EncodeBase64Upload(c *gin.Context, Base64Data []string) ([]string, error) {
	//创建存放图片url切片
	imageUrlData := make([]string, 0)
	if len(Base64Data) > 0 {
		for _, item := range Base64Data {
			//图片拼接规则 当前时间 + 随机字符
			nowTime := time.Now().Format("2006-01-02")
			randString := helpers.RandomString(8)
			//拼接存储本地地址
			localPath := "/Users/bai/img/"

			//正则匹配base64数据,获得img数据与图片格式
			imageData, imageType := Base64RemovePrefix(item)

			//拼接图片名称
			imgName := nowTime + randString + "." + imageType

			//完整存储路径
			localPath = localPath + imgName

			//解码base64数据
			data, err := base64.StdEncoding.DecodeString(imageData)

			if err != nil {
				logger.Fatal(c.Request.URL.Path+"解码Base64失败",
					zap.Time("time", time.Now()),
					zap.Any("error", err),
					zap.String("request", string("解码数据"+imageData)),
				)
			}

			//解码成功将图片写入本地
			//打开本地文件		O_RDWR可读可写权限    O_CREATE如果文件不存在则创建   os.ModePerm文件权限 0777
			f, err := os.OpenFile(localPath, os.O_RDWR|os.O_CREATE, os.ModePerm)
			if err != nil {
				logger.LogWarnIf(err)
				panic(err)
			}
			//写入指定文件内
			f.Write(data)

			//写入成功后上传至阿里云Oss
			imageUrl, err := UploadFileByLocal(localPath, imgName)

			if err != nil {
				logger.Fatal(c.Request.URL.Path+"上传阿里云失败",
					zap.Time("time", time.Now()),
					zap.Any("error", err),
					zap.String("request", string("本地存储地址"+localPath)),
				)
			}

			//上传成功后先关闭文件再删除本地图片
			defer func() {
				RemoveErr := os.Remove(localPath)
				if RemoveErr != nil {
					logger.Fatal(c.Request.URL.Path+"删除本地图片失败",
						zap.Time("time", time.Now()),
						zap.Any("error", err),
						zap.String("request", string("删除文件地址"+localPath)),
					)
				}
			}()

			//关闭文件
			defer f.Close()

			//将图片url添加至存储图片切片中
			imageUrlData = append(imageUrlData, imageUrl)

		}
	} else {
		//通过 errors.New 传入错误的描述信息，就可以创建一个 error 接口类型的变量。
		return imageUrlData, errors.New("base64数据有误")
	}
	return imageUrlData, nil
}

//Base64RemovePrefix 处理base64数据
func Base64RemovePrefix(imgData string) (string, string) {
	// imgData = strings.Replace(imgData, "data:image/jpeg;base64,", "", 1)
	// imgData = strings.Replace(imgData, "data:image/png;base64,", "", 1)
	// imgData = strings.Replace(imgData, "data:image/gif;base64,", "", 1)//Base64RemovePrefix 处理base64数据
	// imgData = strings.Replace(imgData, "data:image/svg+xml;base64,", "", 1)
	imgType := ""
	//设置匹配规则,Compile会返回匹配的内容
	rule := "^data:image/(\\w+)(\\+.*?)?;base64,"
	reg, _ := regexp.Compile(rule)
	subMatch := reg.FindStringSubmatch(imgData)

	// subMath值 (string) (len=26) "data:image/svg+xml;base64,",
	// (string) (len=3) "svg",
	// (string) (len=4) "+xml"
	if len(subMatch) == 0 {
		return "", ""
	}
	//转换为string格式,第二位参数不为空为替换内容
	imgData = reg.ReplaceAllString(imgData, "")
	imgType = subMatch[1]
	//转换为小写
	imgType = strings.ToLower(imgType)
	return imgData, imgType

}

//UploadFileByLocal 上传图片至阿里云
func UploadFileByLocal(localPath string, imgName string) (string, error) {
	//上传地址
	endpoint := config.GetString("oss.end_point")
	// 阿里云账号AccessKey
	accessKeyId := config.GetString("oss.access_key_id")
	// 阿里云账号AccessKeySecret
	accessKeySecret := config.GetString("oss.access_key_secret")
	// 阿里云存储桶名称
	bucketName := config.GetString("oss.bucket_name")
	// 上传文件到OSS时需要指定包含文件后缀在内的完整路径，例如go/upload-test/2022-06-08uYuvCCCn.jpeg。
	objectName := "go/upload-test/" + imgName
	// 由本地文件路径加文件名包括后缀组成，例如/Users/bai/img/2022-06-08uYuvCCCn.jpeg。
	localFileName := localPath

	//创建OSSClient实例
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)

	if err != nil {
		logger.LogWarnIf(err)
		handleError(err)
	}

	// 获取存储空间
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		logger.LogWarnIf(err)
		handleError(err)
	}

	//上传文件
	err = bucket.PutObjectFromFile(objectName, localFileName)
	if err != nil {
		logger.LogWarnIf(err)
		handleError(err)
	}

	//拼接存储阿里云图片路径
	imageUrl := "https://" + bucketName + "." + endpoint + "/" + objectName
	return imageUrl, nil
}
