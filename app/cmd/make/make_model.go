package make

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var CmdMakeModel = &cobra.Command{
	Use:   "model",
	Short: "Crate model file, example: make model user",
	Run:   runMakeModel,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数

}

func runMakeModel(cmd *cobra.Command, arg []string) {

	//格式化模型名称,返回一个Model对象
	model := makeModelFromString(arg[0])

	// 确保模型的目录存在, 例如 `app/models/user`
	dir := fmt.Sprintf("app/model/%s/", model.PackageName)
	//os.MkdirAll 会确保父目录和子目录都会创建,第二个参数是目录权限 使用0777
	os.MkdirAll(dir, os.ModePerm)

	//替换变量
	createFileFromStub(dir+model.PackageName+"_model.go", "model/model", model)
	createFileFromStub(dir+model.PackageName+"util.go", "model/model_util", model)
	createFileFromStub(dir+model.PackageName+"_hooks.go", "model/model_hooks", model)

}
