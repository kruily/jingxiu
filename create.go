/**
* @file: create.go ==>
* @package: main
* @author: jingxiu
* @since: 2022/12/18
* @desc: //TODO
 */

package main

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"os"
	"os/user"
	"strings"
	"text/template"
	"time"
)

func init() {
	registerCommand(&cli.Command{
		Name:    "create",
		Aliases: []string{"c"},
		Usage:   "创建应用控制器",
		Action:  create,
	})
}

type GenController struct {
	File          string   // 要生成的文件名
	Path          string   // 要生成的文件路径
	User          string   // 当前的用户昵称
	Date          string   // 当前的日期
	Package       string   // 要定义的包
	Controller    string   // 要生成的控制器
	Interface     string   // 要生成实现的接口
	Handle        []string // 要生成的所有方法名称
	CurrentHandle string   // 当前要生成的方法
	CurrentRoute  string   // 当前要生成的子路由
}

// Create 定义生成应用控制器命令：
// jingxiu create {$app} {$APIHandler}
// $app 应用名称 $APIHandler 实现的接口
func create(ctx *cli.Context) error {
	// read config from yaml
	if err := Read(".\\etc\\mapping.yaml"); err != nil {
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
	// 生成控制器文件
	g, err := genController(argsSlice)
	if err != nil {
		return err
	}
	// 生成控制器下的接口文件
	genHandler(g)
	return nil
}

func genController(s []string) (*GenController, error) {
	u, _ := user.Current()
	gen := &GenController{
		File:       s[0] + ".go",
		Path:       JingXiu.HandlePath + "\\" + s[0],
		User:       u.Name,
		Date:       time.Now().Format("01/02/2006"),
		Package:    s[0],
		Controller: firstUpper(s[0]),
		Interface:  s[1],
		Handle:     C.Mapping.APIHandleMapping[s[1]],
	}
	// 根据模板生成文件，先根据控制器生成目录
	err := os.Mkdir(gen.Path, os.ModePerm)
	if err != nil {
		fmt.Println(gen.Path + "文件夹生成错误: " + err.Error())
		return nil, errors.New(gen.Path + "文件夹生成错误: " + err.Error())
	}
	//	在控制器目录下生成控制器文件
	ctl, err := os.Create(gen.Path + "\\" + gen.File)
	if err != nil {
		panic(gen.Path + gen.File + "文件夹生成错误: " + err.Error())
	}
	defer ctl.Close()
	ctltmp := template.Must(template.ParseFiles(JingXiu.TemplatePath + "\\controller.tpl"))
	if err = ctltmp.Execute(ctl, gen); err != nil {
		panic(ctl.Name() + " 模板文件生成失败：" + err.Error())
	}
	return gen, nil
}
func genHandler(gen *GenController) {
	handletmp := template.Must(template.ParseFiles(JingXiu.TemplatePath + "\\handle.tpl"))
	for _, item := range gen.Handle {
		filename := strings.ToLower(item) + ".go"
		currentHandle(filename, gen, item, handletmp)
	}
}
