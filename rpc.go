/**
* @file: rpc.go ==>
* @package: main
* @author: jingxiu
* @since: 2023/1/6
* @desc: //TODO
 */

package main

import "github.com/urfave/cli/v2"

func init() {
	registerCommand(&cli.Command{
		Name:    "rpc",
		Aliases: []string{"rp"},
		Usage:   "开启一个gRpc服务",
		Action:  rpc,
	})
}
func rpc(context *cli.Context) error {
	return nil
}
