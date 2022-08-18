package main

import (
	"Gohub/app/cmd"
	"Gohub/app/cmd/make"
	"Gohub/bootstrap"
	btsConfig "Gohub/config"
	"Gohub/pkg/config"
	"Gohub/pkg/console"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	// 加载 config 目录下的配置信息
	btsConfig.Initialize()
}

func main() {

	//应用的主入口, 默认调用cmd.CmdServe命令
	var rootCmd = &cobra.Command{
		Use:   config.Get("app.name"),
		Short: "A simple forum project",
		Long:  `Default will run "serve" command, you can use "-h" flag to see all subcommands`,

		// rootCmd 的所有子命令都会执行以下代码
		PersistentPreRun: func(command *cobra.Command, args []string) {
			// 配置初始化，依赖命令行 --env 参数
			config.InitConfig(cmd.Env)

			// 初始化 Logger
			bootstrap.Setuplogger()

			// 初始化数据库
			bootstrap.SetupDB()

			// 初始化 Redis
			bootstrap.SetupRedis()

			// 初始化缓存
		},
	}

	// 注册子命令
	rootCmd.AddCommand(
		cmd.CmdServe,
		cmd.CmdKey,
		cmd.CmdPlay,
		cmd.CmdMigrate,
		cmd.CmdDBSeed,
		make.CmdMake,
	)

	// 配置默认运行 Web 服务
	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)

	// 注册全局参数，--env
	cmd.RegisterGlobalFlags(rootCmd)

	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}

	/*
		// 配置初始化，依赖命令行 --env 参数
		var env string
		flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
		flag.Parse()
		config.InitConfig(env)

		//初始化 Logger
		bootstrap.Setuplogger()

		// 设置 gin 的运行模式，支持 debug, release, test
		// release 会屏蔽调试信息，官方建议生产环境中使用
		// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
		// 故此设置为 release，有特殊情况手动改为 debug 即可
		gin.SetMode(gin.ReleaseMode) //我们可以有两种方式来设置 gin 的调试模式，一种是通过设置环境变量 GIN_MODE ，第二种是在代码中调用 gin.SetMode(gin.ReleaseMode) 。

		// new 一个 Gin Engine 实例
		router := gin.New()

		// 初始化 DB
		bootstrap.SetupDB()

		// 初始化 Redis
		bootstrap.SetupRedis()
		// 初始化路由绑定
		bootstrap.SetupRoute(router)

		// 运行服务
		err := router.Run(":" + config.Get("app.port"))
		if err != nil {
			// 错误处理，端口被占用了或者其他错误
			fmt.Println(err.Error())
		}
	*/
}
