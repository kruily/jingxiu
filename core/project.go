/**
* @file: project.go ==> core
* @package: core
* @author: jingxiu
* @since: 2022/11/7
* @desc: 生产一个项目
 */

package core

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"io"
	"os"
	"os/exec"
	"sync"
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
*			- jingxiu.go			# 接口启动
*			- go.mod
*		- services				# 微服务服务层，可选
*		- go.work				# go 工作空间目录
 */

// 运行命令：jingxiu start --rpc ${project-name}
func createProject(c *cli.Context) error {
	// git clone
	// 模板项目下载
	if err := clone(); err != nil {
		color.Red("初始化项目【克隆失败】...")
		return err
	}
	//2. 修改文件名
	rpc := c.Bool("rpc")
	args := c.Args()
	if err := os.Rename("./jingxiu_initial_workspace", args.First()); err != nil {
		color.Red("初始化项目【更改文件名失败】...")
		return err
	}
	// 如果用户主动写入有 --rpc 直接返回，不删除
	if rpc {
		return nil
	} else {
		if err := os.RemoveAll(args.First() + "\\services"); err != nil {
			color.Red("初始化项目【取消rpc服务端失败】...")
			return err
		}
		if err := os.RemoveAll(args.First() + "\\gateway\\services"); err != nil {
			color.Red("初始化项目【取消rpc服务端失败】...")
			return err
		}
	}
	// 删除.git 文件
	if err := os.RemoveAll(".git"); err != nil {
		color.Red("初始化项目失败...")
		return err
	}
	color.Blue("初始化项目成功")
	return nil
}

func clone() error {
	color.Blue("正在初始化项目,请稍后...")
	cmd := exec.Command("git", "clone", templateOrigin)
	return PrintCmdOutput(cmd)
}

func PrintCmdOutput(cmd *exec.Cmd) error {
	cmd.Stdin = os.Stdin
	var wg sync.WaitGroup
	wg.Add(2)
	//捕获标准输出
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	readout := bufio.NewReader(stdout)
	go func() {
		defer wg.Done()
		GetOutput(readout)
	}()

	//捕获标准错误
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	readerr := bufio.NewReader(stderr)
	go func() {
		defer wg.Done()
		GetOutput(readerr)
	}()

	//执行命令
	err = cmd.Run()
	if err != nil {
		return err
	}
	wg.Wait()
	return nil
}

func GetOutput(reader *bufio.Reader) {
	var sumOutput string //统计屏幕的全部输出内容
	outputBytes := make([]byte, 200)
	for {
		n, err := reader.Read(outputBytes) //获取屏幕的实时输出(并不是按照回车分割，所以要结合sumOutput)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			sumOutput += err.Error()
		}
		output := string(outputBytes[:n])
		fmt.Print(output) //输出屏幕内容
		sumOutput += output
	}
}
