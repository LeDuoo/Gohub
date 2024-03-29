package make

import (
	"Gohub/pkg/app"
	"Gohub/pkg/console"
	"fmt"

	"github.com/spf13/cobra"
)

var CmdMakeMigration = &cobra.Command{
	Use:   "migration",
	Short: "Create a migration file, example: make migration add_users_table",
	Run:   runMakeMigration,
	Args:  cobra.ExactArgs(2), // 参数数量 1-迁移模型名称 模型内变量名
}

func runMakeMigration(cmd *cobra.Command, args []string) {

	//日期格式化
	timeStr := app.TimenowInTimezone().Format("2006_01_02_150405")

	model := makeModelFromString(args[0])
	//获取新增表名 结构体名称做生成文件的变量替换
	tableStructName := makeModelFromString(args[1])
	model.StructName = tableStructName.StructName
	fileName := timeStr + "_" + model.PackageName
	filePath := fmt.Sprintf("database/migrations/%s.go", fileName)
	createFileFromStub(filePath, "migration", model, map[string]string{"{{FileName}}": fileName})
	console.Success("Migration file created,after modify it, use `migrate up` to migrate database.")
}
