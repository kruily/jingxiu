/**
* @file: readYaml.go ==> core
* @package: core
* @author: jingxiu
* @since: 2022/11/8
* @desc: //TODO
 */

package core

import (
	"gopkg.in/yaml.v3"
	"os"
)

type (
	Config struct {
		Version string  `yaml:"Version"`
		Mapping Mapping `yaml:"Mapping"`
	}
	Mapping struct {
		APIMatchMapping      []string            `yaml:"APIMatchMapping"`
		APIMiddlewareMapping map[string]string   `yaml:"APIMiddlewareMapping"`
		APIHandleMapping     map[string][]string `yaml:"APIHandleMapping"`
	}
)

var C *Config

func Read(path string) error {
	var err error
	C, err = ReadYamlConfig(path)
	return err
}

func ReadYamlConfig(path string) (*Config, error) {
	conf := &Config{}
	if f, err := os.Open(path); err != nil {
		return nil, err
	} else {
		err := yaml.NewDecoder(f).Decode(conf)
		if err != nil {
			return nil, err
		}
	}
	return conf, nil
}
