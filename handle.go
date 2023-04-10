/**
* @file: handle.go ==>
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
	"strings"
	"text/template"
	"time"
)

func init() {
	registerCommand(&cli.Command{
		Name:    "handle",
		Aliases: []string{"h"},
		Usage:   "为一个控制器新增一个接口文件",
		Action:  handle,
	})
}

func handle(ctx *cli.Context) error {
	args := ctx.Args()
	if args.Len() <= 1 {
		//color.Red("生成接口需要提供一个控制器的名字、一个接口名字")
		return errors.New("生成接口需要提供一个控制器的名字、一个接口名字")
	}
	//1. read config from yaml
	if err := Read("./etc/mapping.yaml"); err != nil {
		color.Red("mapping.yaml 文件未找到，请确定mapping.yaml文件在etc目录下")
		return err
	}
	// 判断是否存在命名控制器
	if ok, _ := PathExists(JingXiu.HandlePath + "\\" + args.First()); !ok {
		fmt.Println("不存在命名控制器【" + args.First() + "】")
		return errors.New("不存在命名控制器【" + args.First() + "】")
	}
	s := args.Slice()
	handletmp := template.Must(template.ParseFiles(JingXiu.TemplatePath + "\\handle.tpl"))
	mapper := map[string]interface{}{
		"File":          s[1] + ".go",
		"Path":          JingXiu.HandlePath + "\\" + s[0],
		"Date":          time.Now().Format("01/02/2006"),
		"Package":       s[0],
		"CurrentHandle": firstUpper(s[1]),
		"Controller":    firstUpper(s[0]),
		"CurrentRoute":  s[1],
	}
	filename := mapper["File"].(string)
	file, err := os.Create(mapper["Path"].(string) + "\\" + filename)
	if err != nil {
		panic(mapper["Path"].(string) + "\\" + filename + "文件生成错误: " + err.Error())
	}
	defer file.Close()
	if err = handletmp.Execute(file, mapper); err != nil {
		panic(file.Name() + "模板文件生成失败: " + err.Error())
	}
	return nil
}

func currentHandle(filename string, gen *GenController, item string, tmp *template.Template) {
	file, err := os.Create(gen.Path + "\\" + filename)
	if err != nil {
		panic(gen.Path + filename + "文件生成错误: " + err.Error())
	}
	defer file.Close()
	gen.CurrentHandle = item
	gen.CurrentRoute = strings.ToLower(item)
	if err = tmp.Execute(file, gen); err != nil {
		panic(file.Name() + "模板文件生成失败: " + err.Error())
	}
}
