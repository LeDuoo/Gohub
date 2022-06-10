package config

import "Gohub/pkg/config"

func init() {
	config.Add("oss", func() map[string]interface{} {
		return map[string]interface{}{

			// 请求地址
			"end_point": "oss-cn-shenzhen.aliyuncs.com",

			// 阿里云账号AccessKey
			"access_key_id": "LTAI5tEoRWPTtvG6KQkaW4jA",

			// 阿里云账号AccessKeySecret
			"access_key_secret": "dl6oM67AsT9R5mfsr9E81DlrbffFPJ",

			// 阿里云存储桶名称
			"bucket_name": "ayng-local",
		}
	})
}
