/**
* @file: database.go ==> core
* @package: core
* @author: jingxiu
* @since: 2022/11/7
* @desc: //TODO
 */

package core

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"os"
)

func generateDatabase(c *cli.Context) error {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "./data/query",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
		ModelPkgPath: "./data/model",
	})
	conn, err := ConnectDB(workspace + "\\gateway\\gateway.yaml")
	if err != nil {
		fmt.Printf("%#v", err.Error())
		return err
	}
	g.UseDB(conn)
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
	return nil
}

func ConnectDB(path string) (conn *gorm.DB, err error) {
	type DB struct {
		Type   string `yaml:"Type"`   // 链接类型
		Source string `yaml:"Source"` // 链接dns地址
	}
	config := &struct {
		DB DB `yaml:"DB"`
	}{}
	if f, err := os.Open(path); err != nil {
		return nil, errors.New("配置文件读取失败：" + err.Error())
	} else {
		err := yaml.NewDecoder(f).Decode(config)
		if err != nil {
			return nil, errors.New("配置文件读取失败：" + err.Error())
		}
	}
	conn, err = gorm.Open(mysql.Open(config.DB.Source), &gorm.Config{})
	if err != nil {
		//panic(fmt.Errorf("cannot establish db connection: %w", err))
		return nil, errors.New("数据库链接失败：" + err.Error())
	}
	return conn, nil
}
