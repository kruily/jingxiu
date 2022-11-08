/**
* @file: parallel.go ==> core
* @package: core
* @author: jingxiu
* @since: 2022/11/7
* @desc: //TODO
 */

package core

import (
	"os"
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
