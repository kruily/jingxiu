/**
* @file: handler.go ==> core
* @package: core
* @author: jingxiu
* @since: 2022/11/7
* @desc: 生成控制器
 */

package core

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"text/template"
	"time"
)

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
}

func GenHandles(s []string) {
	// 先查找是否不存在cli,不存在就下载
	if ok, _ := PathExists(templatePath); !ok {
		if err := goGetJingXiuCli(); err != nil {
			fmt.Println("创建失败【cli 模板集下载失败】")
			return
		}
	}
	u, _ := user.Current()
	gen := &GenController{
		File:       s[0] + ".go",
		Path:       workspace + "\\handle\\" + s[0],
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
		//panic(gen.Path + "文件夹生成错误: " + err.Error())
		fmt.Println(gen.Path + "文件夹生成错误: " + err.Error())
		return
	}
	//	在控制器目录下生成控制器文件
	ctl, err := os.Create(gen.Path + "\\" + gen.File)
	if err != nil {
		panic(gen.Path + gen.File + "文件夹生成错误: " + err.Error())
	}
	defer ctl.Close()
	ctltmp := template.Must(template.ParseFiles(templatePath + "\\controller.tpl"))
	if err = ctltmp.Execute(ctl, gen); err != nil {
		panic(ctl.Name() + " 模板文件生成失败：" + err.Error())
	}
	handletmp := template.Must(template.ParseFiles(templatePath + "\\handle.tpl"))
	for _, item := range gen.Handle {
		filename := strings.ToLower(item) + ".go"
		currentHandle(filename, gen, item, handletmp)
	}
}

func currentHandle(filename string, gen *GenController, item string, tmp *template.Template) {
	file, err := os.Create(gen.Path + "\\" + filename)
	if err != nil {
		panic(gen.Path + filename + "文件生成错误: " + err.Error())
	}
	defer file.Close()
	gen.CurrentHandle = item
	if err = tmp.Execute(file, gen); err != nil {
		panic(file.Name() + "模板文件生成失败: " + err.Error())
	}
}

func goGetJingXiuCli() error {
	cmd := exec.Command("go", "get", "github.com/jingxiu1016/cli@"+C.Version)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("下载【打开输出流失败】")
		return err
	}
	defer stdout.Close()
	if err := cmd.Start(); err != nil {
		fmt.Println("下载【命令运行输出流失败】")
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
