/**
* @file: runProject.go ==> core
* @package: core
* @author: jingxiu
* @since: 2022/11/18
* @desc: 监控项目路径下所有目录的更新时间，如果更新时间就发生在刚才的五秒（暂定）偏差之内，重新运行项目
 */

package core

import (
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func runProject(c *cli.Context) error {
	runGateWay()
	return nil
}

func runGateWay() {
	if err := command("go", "run", "./gateway/main.go"); err != nil {
		color.Red("项目运行【失败】：" + err.Error())
	}
}
