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
    "github.com/gin-gonic/gin"
)

type {{.Controller}} struct {
    Config *config.Config // 全局配置项
    // 可以挂载一些连接型实例对象参数，例如mysql链接实例，kafka 链接，rpc 链接，redis 链接等。
}

func (*{{.Controller}}) Route(c *gin.Engine, do func(c *gin.Engine)) {
	do(c)
}

func New{{.Controller}}Handle() *{{.Controller}} {
	return &{{.Controller}}{
	    Config: config.C,
	}
}
