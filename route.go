/**
* @file: route.go ==>
* @package: main
* @author: jingxiu
* @since: 2022/12/18
* @desc: //TODO
 */

package main

import (
	"bufio"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"io"
	"os"
	"strings"
	"text/template"
	"time"
)

func init() {
	registerCommand(&cli.Command{
		Name:    "route",
		Aliases: []string{"r"},
		Usage:   "生成路由组",
		Action:  route,
	})
}

type GenRoute struct {
	Handle     string // 解析匹配到的方法名称
	Group      string // 解析匹配到的路由组名称
	Route      string // 解析匹配到的路由名
	Method     string // 解析匹配到的HTTP方法名
	Middleware string // 解析匹配到的中间件方法
	Doc        string // 解析文档
}

var register = make(map[string][]*GenRoute)

func route(ctx *cli.Context) error {
	//1. read config from yaml
	if err := Read("./etc/mapping.yaml"); err != nil {
		color.Red("mapping.yaml 文件未找到，请确定mapping.yaml文件在gateway目录下")
		return err
	}
	args := ctx.Args()
	if args.Len() <= 0 {
		rangeDir(JingXiu.HandlePath)
		for key, value := range register {
			writeRouterFile(JingXiu.RouterPath, key, value)
		}
	} else {
		if args.Slice()[0] == "append" {
			for _, item := range args.Slice()[1:] {
				rangeDir(JingXiu.HandlePath + "\\" + item)
				for key, value := range register {
					writeRouterFile(JingXiu.RouterPath, key, value)
				}
			}
		} else {
			color.Green("如果生成路由文件需要指定某个文件的时候，那么应该使用 append 命令")
		}
	}
	return nil
}

// 遍历文件夹内容下的内容
func rangeDir(path string) {
	dir, err := os.ReadDir(path)
	if err != nil {
		color.Red("接口目录扫描失败！", err.Error())
		return
	}
	for _, item := range dir {
		if item.IsDir() {
			color.Green("进入目录：" + JingXiu.HandlePath + "\\" + item.Name())
			rangeDir(JingXiu.HandlePath + "\\" + item.Name())
		} else {
			sr := strings.Split(path, "\\")
			if !strings.Contains(item.Name(), ".go") {
				continue
			}
			color.Green("扫描接口文件：" + path + "\\" + item.Name())
			register[sr[len(sr)-1]] = append(register[sr[len(sr)-1]], openFile(path+"\\"+item.Name())...)
		}
	}
}

// 文件扫描
func openFile(file string) []*GenRoute {
	if !strings.Contains(file, ".go") {
		return nil
	}
	open, err := os.Open(file)
	if err != nil {
		panic(file + "文件打开错误" + err.Error())
	}
	defer open.Close()
	reader := bufio.NewReader(open)
	var results []string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		str := string(line)
		if attr, body, ok := Comment(str); ok {
			results = append(results, attr+" "+body)
		}
	}
	var routeList = make([]*GenRoute, 0)
	if len(results) > 0 {
		routeList = append(routeList, matchKeywords(results))
	}
	return routeList
}

// 匹配注释中的关键词
func matchKeywords(info []string) *GenRoute {
	first := trimPrefix(info[0])
	temp := &GenRoute{
		Handle: first,
	}
	for _, item := range C.Mapping.APIMatchMapping {
		for _, fo := range info[1:] {
			if strings.Contains(fo, item) {
				switch item {
				case "@Handle":
					str := trimPrefix(fo)
					temp.Handle = strings.Split(str, " ")[1]
				case "@Router":
					//	1. 匹配 /*/* 得到路由
					if body, ok := RouteReg(fo); ok {
						ar := strings.Split(body, " ")
						//	2. 从 第一步中获取路由组
						temp.Group = strings.SplitN(ar[0], "/", 3)[1]
						// 3. 从第一步中获取子路由
						temp.Route = strings.Replace(ar[0], "/"+temp.Group, "", 1)
						// 4. 从第一不中获取http-method
						left, right := indexBrackets(ar[1])
						temp.Method = strings.ToUpper(ar[1][left+1 : right])
					}
				case "@Middleware":
					mw := trimPrefix(fo)
					left, right := indexBrackets(mw)
					temp.Middleware = transitMiddle(strings.Split(mw[left+1:right], "|"))
				case "@Summary":
					doc := trimPrefix(fo)
					left, right := indexBrackets(doc)
					temp.Doc = doc[left+1 : right]
				}
			}
		}
	}
	return temp
}

// 写入路由文件
func writeRouterFile(path, key string, value []*GenRoute) {
	filename := key + "_router.gen.go"
	funcName := firstUpper(key) + "Router"
	//写入文件时，使用带缓存的 *Writer
	data := map[string]interface{}{
		"filename":     filename,
		"filepath":     path,
		"date":         time.Now().Format("01/02/2006"),
		"doc":          key + " 路由",
		"funcName":     funcName,
		"higherDir":    key,
		"pak":          firstUpper(key),
		"group":        key,
		"routers":      value,
		"middleImport": middleImport(value),
	}
	tmp := template.Must(template.ParseFiles(JingXiu.TemplatePath + "\\route.tpl"))
	file, err := os.OpenFile(path+"\\"+filename, os.O_CREATE, 0666)
	if err != nil {
		panic(filename + "文件创建失败")
	}
	defer file.Close()
	err = tmp.Execute(file, data)
	if err != nil {
		panic(filename + " 模板文件生成失败：" + err.Error())
	}
}
