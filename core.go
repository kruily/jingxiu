/**
* @file: core.go ==>
* @package: main
* @author: jingxiu
* @since: 2022/12/18
* @desc: //TODO
 */

package main

import (
	"errors"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"os"
)

// JingxiuCLI 命令结构体
type JingxiuCLI struct {
	Version      string   // 版本号
	Workspace    string   // 工作目录
	TemplatePath string   // 模板目录
	HandlePath   string   // 接口生成目录
	RouterPath   string   // 路由生成目录
	Error        error    // 出错信息
	App          *cli.App // cli 实例
}

var JingXiu *JingxiuCLI
var Commands []*cli.Command

func GetJingXiuCLI() *JingxiuCLI {
	if JingXiu == nil {
		JingXiu = new(JingxiuCLI)
		JingXiu.Version = "v0.3.7"
		JingXiu.App = &cli.App{
			Name:     "jingxiu",
			Usage:    "你好，这是一个项目快速开发脚手架",
			Commands: Commands,
			Flags:    []cli.Flag{},
		}
		JingXiu.Workspace, _ = os.Getwd()
		JingXiu.TemplatePath = os.Getenv("GOPATH") + "\\pkg\\mod\\github.com\\jingxiu1016\\jingxiu@" + JingXiu.Version + "\\tpl"
		JingXiu.HandlePath = JingXiu.Workspace + "\\gateway\\handle"
		JingXiu.RouterPath = JingXiu.Workspace + "\\gateway\\router"
	}
	return JingXiu
}

func (c *JingxiuCLI) check() {
	if c.Version == "" {
		c.Error = errors.New("jingxiu cli version error")
		return
	}
	if c.Workspace == "" {
		c.Error = errors.New("jingxiu cli not in right workspace")
		return
	}
	if c.TemplatePath == "" {
		c.Error = errors.New("the correct template directory was not found,please check whether the cli version is correct")
		return
	}
}

func (c *JingxiuCLI) Run() {
	c.check()
	if c.Error != nil {
		color.Red(c.Error.Error())
		return
	}
	if err := c.App.Run(os.Args); err != nil {
		color.Red(err.Error())
		return
	}
}

func registerCommand(cmd ...*cli.Command) {
	Commands = append(Commands, cmd...)
}
