package cmd

import (
	"Gohub/pkg/console"
	"Gohub/pkg/helpers"

	"github.com/spf13/cobra"
)

var CmdKey = &cobra.Command{
	Use:   "key",
	Short: "Generate App Key, will print the generated Key",
	Run:   runKeyGenerate,
	Args:  cobra.NoArgs, //不允许传参
	/*Arg参数
	NoArgs - 如果存在任何位置参数，该命令将报错
	ArbitraryArgs - 该命令会接受任何位置参数
	OnlyValidArgs - 如果有任何位置参数不在命令的 ValidArgs 字段中，该命令将报错
	MinimumNArgs(int) - 至少要有 N 个位置参数，否则报错
	MaximumNArgs(int) - 如果位置参数超过 N 个将报错
	ExactArgs(int) - 必须有 N 个位置参数，否则报错
	`ExactValidArgs (int) 必须有 N 个位置参数，且都在命令的 ValidArgs 字段中，否则报错
	RangeArgs(min, max) - 如果位置参数的个数不在区间 min 和 max 之中，报错
	MatchAll(pargs ...PositionalArgs) - 支持使用以上的多个验证器
	*/
}

func runKeyGenerate(cmd *cobra.Command, args []string) {
	console.Success("----")
	console.Success("App Key:")
	console.Success(helpers.RandomString(32))
	console.Success("----")
	console.Warning("please go to .env file to change the APP_KEY option")

}
