/**
* @file: router.go ==> core
* @package: core
* @author: jingxiu
* @since: 2022/11/7
* @desc: //TODO
 */

package core

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli/v2"
	"io"
	"os"
	"strings"
	"text/template"
	"time"
)

type GenRoute struct {
	Handle     string // 解析匹配到的方法名称
	Group      string // 解析匹配到的路由组名称
	Route      string // 解析匹配到的路由名
	Method     string // 解析匹配到的HTTP方法名
	Middleware string // 解析匹配到的中间件方法
	Doc        string // 解析文档
}

var register = make(map[string][]*GenRoute)

func generateRouters(c *cli.Context) error {
	//1. read config from yaml
	if err := Read("./gateway/mapping.yaml"); err != nil {
		fmt.Println("mapping.yaml 文件未找到，请确定mapping.yaml文件在gateway目录下")
		return err
	}
	// 确定模板地址
	templatePath = os.Getenv("GOPATH") + "\\pkg\\mod\\github.com\\jingxiu1016\\jingxiu@" + C.Version + "\\tpl"
	// 先查找是否不存在cli,不存在就下载
	if ok, _ := PathExists(templatePath); !ok {
		if err := command("go", "get", "github.com/jingxiu1016/jingxiu@"+C.Version); err != nil {
			fmt.Println("创建失败【cli 模板集下载失败】")
			return err
		}
	}
	args := c.Args()
	if args.Len() <= 0 {
		rangeDir(handlerPath)
		for key, value := range register {
			writeRouterFile(routerPath, key, value)
		}
	} else {
		if args.Slice()[0] == "append" {
			for _, item := range args.Slice()[1:] {
				rangeDir(handlerPath + "\\" + item)
				for key, value := range register {
					writeRouterFile(routerPath, key, value)
				}
			}
		} else {
			fmt.Println("如果生成路由文件需要指定某个文件的时候，那么应该使用 append 命令")
		}
	}
	return nil
}

// 遍历文件夹内容下的内容
func rangeDir(path string) {
	dir, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("接口目录扫描失败！", err.Error())
		return
	}
	for _, item := range dir {
		if item.IsDir() {
			fmt.Println("进入目录：" + handlerPath + "\\" + item.Name())
			rangeDir(handlerPath + "\\" + item.Name())
		} else {
			sr := strings.Split(path, "\\")
			if !strings.Contains(item.Name(), ".go") {
				continue
			}
			fmt.Println("扫描接口文件：" + path + "\\" + item.Name())
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
		if strings.Contains(str, "//") && strings.IndexAny(str, "/") == 0 {
			results = append(results, str)
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
				case "@Group":
					group := trimPrefix(fo)
					left, right := indexBrackets(group)
					temp.Group = group[left+1 : right]
				case "@Route":
					route := trimPrefix(fo)
					left, right := indexBrackets(route)
					temp.Route = route[left+1 : right]
				case "@Method":
					method := trimPrefix(fo)
					left, right := indexBrackets(method)
					temp.Method = strings.ToUpper(method[left+1 : right])
				case "@Middleware":
					mw := trimPrefix(fo)
					left, right := indexBrackets(mw)
					temp.Middleware = transitMiddle(strings.Split(mw[left+1:right], "|"))
				case "@Doc":
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
	tmp := template.Must(template.ParseFiles(templatePath + "\\route.tpl"))
	create, err := os.OpenFile(path+"\\"+filename, os.O_CREATE, 0666)
	if err != nil {
		panic(filename + "文件创建失败")
	}
	defer create.Close()
	err = tmp.Execute(create, data)
	if err != nil {
		panic(filename + " 模板文件生成失败：" + err.Error())
	}
}

func middleImport(value []*GenRoute) bool {
	for _, route := range value {
		if len(route.Middleware) > 0 {
			return true
		}
	}
	return false
}

func trimPrefix(s string) string {
	s = strings.TrimPrefix(s, "//")
	s = strings.TrimSpace(s)
	return s
}
func indexBrackets(s string) (int, int) {
	return strings.Index(s, "["), strings.Index(s, "]")
}

func transitMiddle(mi []string) string {
	str := ""
	for _, item := range mi {
		str += C.Mapping.APIMiddlewareMapping[item] + ", "
	}
	return str[:len(str)-2]
}
