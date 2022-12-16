/**
* @file: depend.go ==> core
* @package: core
* @author: jingxiu
* @since: 2022/12/4
* @desc: //TODO
 */

package core

import (
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"sync"
)

func downloadDepend(ctx *cli.Context) error {
	tidy := func() error {
		if err := command("go", "mod", "tidy"); err != nil {
			return err
		}
		return nil
	}
	cd := func(path string) error {
		if err := command("cd", path); err != nil {
			return err
		}
		return nil
	}
	Parallel(func(group *sync.WaitGroup) {
		defer group.Done()
		if err := cd(".\\common"); err != nil {
			color.Red(err.Error())
			return
		}
		if err := tidy(); err != nil {
			color.Red(err.Error())
			return
		}
	}, func(group *sync.WaitGroup) {
		defer group.Done()
		if err := cd(".\\data"); err != nil {
			color.Red(err.Error())
			return
		}
		if err := tidy(); err != nil {
			color.Red(err.Error())
			return
		}
	}, func(group *sync.WaitGroup) {
		defer group.Done()
		if err := cd(".\\gateway"); err != nil {
			color.Red(err.Error())
			return
		}
		if err := tidy(); err != nil {
			color.Red(err.Error())
			return
		}
	})
	return nil
}
