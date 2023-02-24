package make

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CmdMakeSeeder = &cobra.Command{
	Use:   "seeder",
	Short: "Create seeder file, example:  make seeder user",
	Run:   runMakeSeeder,
	Args:  cobra.MinimumNArgs(1), //参数数量 1-模型名称 2-创建数量
}

func runMakeSeeder(cmd *cobra.Command, args []string) {

	// 格式化模型名称 返回一个model 对象
	model := makeModelFromString(args[0])

	// 定义创建数据数量,默认为10
	variables := map[string]string{
		"createCount": "10",
	}
	if len(args) > 1 {
		variables["createCount"] = args[1]
	}

	//拼接目标文件路径
	filepath := fmt.Sprintf("database/seeders/%s_seeder.go", model.TableName)

	// 基于模板创建文件 (做好标量替换)
	createFileFromStub(filepath, "seeder", model, variables)
}
