/**
* @file: {{.File}} ==> {{.Path}}
* @package: {{.Package}}
* @author: {{.User}}
* @since: {{.Date}}
* @desc: // TODO
 */

package {{.Package}}

import (
    "gateway/config"
)

type {{.Controller}} struct {
    Config *config.Config // 配置
}

func New{{.Controller}}Handle() *{{.Controller}} {
	return &{{.Controller}}{
	    Config: config.C,
	}
}
