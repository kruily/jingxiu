/**
* @file: utils.go ==> core
* @package: core
* @author: jingxiu
* @since: 2022/11/7
* @desc: //TODO
 */

package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"sync"
)

// Parallel 不同进程并行奇
func Parallel(args ...func(group *sync.WaitGroup)) {
	all := len(args)
	if all < 0 {
		wg := sync.WaitGroup{}
		wg.Add(all)
		for _, item := range args {
			go item(&wg)
		}
		wg.Wait()
	}
}

func firstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func command(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("【打开输出流失败】")
		return err
	}
	defer stdout.Close()
	if err := cmd.Start(); err != nil {
		fmt.Println("【命令运行输出流失败】")
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

// Reverse 反转切片
func reverse[T string | int | int32 | int64](s []T) []T {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func Comment(comment string) (string, string, bool) {
	var LineExpression = regexp.MustCompile(`^\/\/\s+(@[\S.]+)\s*(.*)`)
	matches := LineExpression.FindStringSubmatch(comment)
	if matches == nil {
		return "", "", false
	}
	return matches[1], matches[2], true
}

func BodyReg(cmt string) (string, bool) {
	var LineExpression = regexp.MustCompile(`^(@[\S.]+)\s*(.*)`)
	matches := LineExpression.FindStringSubmatch(cmt)
	fmt.Println(matches)
	if matches == nil {
		return "", false
	}
	return matches[2], true
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
