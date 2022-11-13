/**
* @file: controller.go ==> core
* @package: core
* @author: jingxiu
* @since: 2022/11/7
* @desc: //TODO
 */

package core

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

// createController 定义生成应用控制器命令：
// jingxiu create {$app} {$APIHandler}
// $app 应用名称 $APIHandler 实现的接口
func createController(c *cli.Context) error {
	//1. read config from yaml
	if err := Read("./mapping.yaml"); err != nil {
		fmt.Println("mapping.yaml 文件未找到，请确定mapping.yaml文件在gateway目录下")
		return err
	}
	// 确定模板地址
	templatePath = os.Getenv("GOPATH") + "\\pkg\\mod\\github.com\\jingxiu1016\\cli@" + C.Version + "\\tpl"
	args := c.Args()
	var argsSlice []string
	if _, ok := C.Mapping.APIHandleMapping[args.First()]; ok {
		argsSlice = reverse(args.Slice())
		fmt.Println("温馨提示：其实我并不建议你先写要实现的接口，你应该先将控制器明确...")
	} else {
		argsSlice = args.Slice()
	}
	GenHandles(argsSlice)
	return nil
}

// Reverse 反转切片
func reverse[T string | int | int32 | int64](s []T) []T {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
