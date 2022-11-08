/**
* @file: cmd.go ==> core
* @package: core
* @author: jingxiu
* @since: 2022/11/7
* @desc: //TODO
 */

package core

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var (
	workspace, _ = os.Getwd()
	templatePath = os.Getenv("GOPATH") + "\\pkg\\mod\\github.com\\jingxiu1016\\cli@" + C.Version + "\\tpl"
	handlerPath  = workspace + "\\handler"
	routerPath   = workspace + "\\router"
)

// JingXiu jingXiu 主命令
func JingXiu() {
	app := &cli.App{
		Name:     "jingxiu",
		Usage:    "你好，这是一个代码快速生成命令",
		Commands: someCommands(),
		Flags:    []cli.Flag{},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal("命令未启动：" + err.Error())
	}
}

func someCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:    "start",
			Aliases: []string{"s"},
			Usage:   "开始一个jingxiu cli web脚手架",
			Action:  createProject,
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:     "rpc",
					Aliases:  []string{"r"},
					Usage:    "是否生产 gRPC 工作目录",
					Required: false,
				},
			},
		},
		{
			Name:    "create",
			Aliases: []string{"c"},
			Usage:   "创建应用控制器",
			Action:  createController,
		},
		{
			Name:    "route",
			Aliases: []string{"r"},
			Usage:   "生成路由组",
			Action:  generateRouters,
		},
		{
			Name:    "model",
			Aliases: []string{"m"},
			Usage:   "在当前目录下，从配置的链接数据库中生成dao层",
			Action:  generateDatabase,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "src",
					Aliases:  []string{"s"},
					Usage:    "指定要读取的配置文件",
					Required: true,
				},
			},
		},
	}
}
