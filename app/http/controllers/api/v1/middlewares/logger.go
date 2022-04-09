package middlewares

import (
    "bytes"
    "Gohub/pkg/helpers"
    "Gohub/pkg/logger"
    "io/ioutil"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/spf13/cast"
    "go.uber.org/zap"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

//Logger 记录请求日志
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {

		//获取 response 内容
		w := &responseBodyWriter{ResponseWriter: c.Writer, body: &bytes.Buffer{}}
		c.Writer = w

		// 获取请求数据
		var requestBody []byte
		if c.Request.Body != nil {
			// c.Request.Body 是一个 buffer 对象,只能读一次
			requestBody, _ = ioutil.ReadAll(c.Request.Body)
			//读取后重新赋值 c.Resuest.Body, 以供后续的其他操作
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))
		}

		//设置开始时间
		start := time.Now()

		//Multiplexer根据URL将请求路由给指定的Handler。Handler用于处理请求并给予响应。更严格地说，用来读取请求体、并将请求对应的响应字段(respones header)写入ResponseWriter中，然后返回。
		//详细handler解析☞ https://www.cnblogs.com/f-ck-need-u/p/10020951.html
		//c.Next() 之前的操作是在 Handler 执行之前就执行；  之前的操作一般用来做验证处理，访问是否允许之类的。
		//c.Next() 之后的操作是在 Handler 执行之后再执行；  之后的操作一般是用来做总结处理，比如格式化输出、响应结束时间，响应时长计算之类的。
		c.Next()

		// 开始记录日志的逻辑
		cost := time.Since(start) //time.Since可以计算出一个事件戳到当前的时间差
		responStatus := c.Writer.Status()

		logFields := []zap.Field{
			zap.Int("status", responStatus),
			zap.String("request", c.Request.Method+" "+c.Request.URL.String()),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.String("time", helpers.MicrosecondsStr(cost)),
		}
		if c.Request.Method == "POST" || c.Request.Method =="PUT" || c.Request.Method == "DELETE"{
			//请求的内容
			logFields = append(logFields, zap.String("Request Body",string(requestBody)))

			//响应的内容
			logFields = append(logFields, zap.String("Respone Body",w.body.String()))
		}

		if responStatus > 400 && responStatus <= 499{
			// 除了 StatusBadRequest 以外，warning 提示一下，常见的有 403 404，开发时都要注意
			logger.Warn("HTTP Warning" + cast.ToString(responStatus),logFields...)// ...函数有多个不定参数的情况，可以接受多个不确定数量的参数

		} else if responStatus >= 500 && responStatus <= 599 {
			//除了内部错误,记录error
			logger.Error("HTTP Error" +cast.ToString(responStatus),logFields...)
		}else{
			logger.Debug("HTTP Access log", logFields...)
		}
	}
}
