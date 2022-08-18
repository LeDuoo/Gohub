package cmd

import (
	"Gohub/database/seeders"
	"Gohub/pkg/console"
	"Gohub/pkg/seed"

	"github.com/spf13/cobra"
)

//cmd命令配置
var CmdDBSeed = &cobra.Command{
	Use:   "seed",
	Short: "Insert fake data to the database",
	Run:   runSeeders,
	Args:  cobra.MaximumNArgs(1), //只允许一个参数
}

func runSeeders(cmd *cobra.Command, args []string) {
	//执行seed初始化
	seeders.Initalize()

	if len(args) > 0 {
		//有传入参数
		name := args[0]
		seeder := seed.GetSeeder(name)
		if len(seeder.Name) > 0 {
			//执行指定Seeder2
			seed.RunSeeder(name)
		} else {
			console.Error("Seeder not found: " + name)
		}
	} else { //默认执行所有Seeder
		seed.RunAll()
		console.Success("Done seeding")
	}
}
