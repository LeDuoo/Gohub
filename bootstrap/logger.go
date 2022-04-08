package bootstrap

import (
	"Gohub/pkg/config"
    "Gohub/pkg/logger"
)

// Setuplogger 初始化Logger
func Setuplogger(){

	logger.InitLogger(
		config.GetString("log.filename"),
        config.GetInt("log.max_size"),
        config.GetInt("log.max_backup"),
        config.GetInt("log.max_age"),
        config.GetBool("log.compress"),
        config.GetString("log.type"),
        config.GetString("log.level"),
	)
}