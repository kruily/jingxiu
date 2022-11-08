/**
* @file: project.go ==> core
* @package: core
* @author: jingxiu
* @since: 2022/11/7
* @desc: 生产一个项目
 */

package core

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"io"
	"os"
	"os/exec"
)

var templateOrigin = `https://github.com/jingxiu1016/jingxiu_initial_workspace.git`

// 创建一个目录如下的项目（项目名称）
/**
*	project
*		- common				# 公共组件包，包含一些常见的内容
*			- go.mod
*			- README.md
*		- data					# 数据转化层，包含dao层和查询层
*			- go.mod
*			- README.md
*		- gateway				# api 网关层，作为所有api访问的入口
*			- config			# api 网关配置管理中心
*				- config.go
*			- handle			# api 入口层, 生成的控制器文件存储在以下
*			- middleware		# api 中间件
*			- router			# api 路由层，生成的路由文件，存储在以下
*			- services			# gRpc 客户端接入控制层，可选
*			- gateway.yaml		# 网关配置文件
*			- main.go			# 接口启动
*			- go.mod
*		- services				# 微服务服务层，可选
*		- go.work				# go 工作空间目录
 */

// 运行命令：jingxiu start --rpc ${project-name}
func createProject(c *cli.Context) error {
	// git clone
	// 模板项目下载
	if err := clone(); err != nil {
		fmt.Println("初始化项目【克隆失败】...")
		return err
	}
	//2. 修改文件名
	rpc := c.Bool("rpc")
	args := c.Args()
	if err := os.Rename("./jingxiu_initial_workspace", args.First()); err != nil {
		fmt.Println("初始化项目【更改文件名失败】...")
		return err
	}
	// 如果用户主动写入有 --rpc 直接返回，不删除
	if rpc {
		return nil
	} else {
		if err := os.RemoveAll(args.First() + "\\services"); err != nil {
			fmt.Println("初始化项目【取消rpc服务端失败】...")
			return err
		}
		if err := os.RemoveAll(args.First() + "\\gateway\\services"); err != nil {
			fmt.Println("初始化项目【取消rpc客户端失败】...")
			return err
		}
	}
	// 删除.git 文件
	if err := os.RemoveAll(".git"); err != nil {
		fmt.Println("初始化项目...")
		return err
	}
	fmt.Println("初始化项目成功")
	return nil
}

func mkdir(path string) error {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return errors.New("项目初始化：" + path + "创建失败")
	}
	return nil
}

func clone() error {
	cmd := exec.Command("git", "clone", templateOrigin)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("克隆【打开输出流失败】")
		return err
	}
	defer stdout.Close()
	if err := cmd.Start(); err != nil {
		fmt.Println("克隆【命令运行输出流失败】")
		return err
	}
	if opBytes, err := io.ReadAll(stdout); err != nil { // 读取输出结果
		fmt.Println(err.Error())
		return err
	} else {
		fmt.Println(string(opBytes))
	}
	return nil
}
