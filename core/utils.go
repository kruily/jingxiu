/**
* @file: utils.go ==> core
* @package: core
* @author: jingxiu
* @since: 2022/11/7
* @desc: //TODO
 */

package core

import (
	"fmt"
	"io"
	"os"
	"os/exec"
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

func commend(name string, args ...string) error {
	cmd := exec.Command(name, args...)
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
