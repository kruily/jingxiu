/**
* @file: create.go ==>
* @package: main
* @author: jingxiu
* @since: 2022/12/18
* @desc: //TODO
 */

package main

import (
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func init() {
	registerCommand(&cli.Command{
		Name:    "create",
		Aliases: []string{"c"},
		Usage:   "创建应用控制器",
		Action:  create,
	})
}

// Create 定义生成应用控制器命令：
// jingxiu create {$app} {$APIHandler}
// $app 应用名称 $APIHandler 实现的接口
func create(ctx *cli.Context) error {
	// read config from yaml
	if err := Read("./etc/mapping.yaml"); err != nil {
		color.Red("mapping.yaml 文件未找到，请确定mapping.yaml文件在gateway目录下")
		return err
	}
	args := ctx.Args()
	var argsSlice []string
	if _, ok := C.Mapping.APIHandleMapping[args.First()]; ok {
		argsSlice = reverse(args.Slice())
		color.Green("温馨提示：其实我并不建议你先写要实现的接口，你应该先将控制器明确...")
	} else {
		argsSlice = args.Slice()
	}
	if len(argsSlice) <= 1 {
		color.Red("创建一个控制器时，需要指定其要实现的API接口")
		return nil
	}
	// 生成控制器的接口文件
	genHandles(argsSlice)
	return nil
}

func genHandles(argsSlice []string) {

}
