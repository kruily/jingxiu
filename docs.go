/**
* @file: docs.go ==>
* @package: main
* @author: jingxiu
* @since: 2023/1/6
* @desc: //TODO
 */

package main

import "github.com/urfave/cli/v2"

func init() {
	registerCommand(&cli.Command{
		Name:    "docs",
		Aliases: []string{"d"},
		Usage:   "生成api文档",
		Action:  docs,
	})
}
func docs(context *cli.Context) error {
	// 生成swag 文档
	if err := command("swag", "init"); err != nil {
		return err
	}
	return nil
}
