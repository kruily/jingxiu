/**
* @file: {{.filename}}==> {{.filepath}}
* @package: router
* @author: jingxiu
* @since: {{.date}}
* @desc: {{.doc}}
 */

package router

import (
    "gateway/handle/{{.higherDir}}"
    {{if .middleImport }}"gateway/middleware"{{end}}
    "github.com/gin-gonic/gin"
)

var {{.pak}}Obj *{{.higherDir}}.{{.pak}}

func init(){
    instance = {{.pak}}.New{{.pak}}Handle()
    Register(RegisterRoute{
    	Ro: {{.pak}}Obj,
        Do: func(e *gin.Engine) { {{.pak}}Router(e) },
    })
}

func {{.funcName}}(c *gin.Engine){
    group := c.Group("/{{.group}}/")
    { {{range $k,$v := .routers}}{{if and $v.Method  $v.Route}}
       // {{$v.Doc}}
       group.{{$v.Method}}("{{$v.Route}}", {{if $v.Middleware }}{{$v.Middleware}},{{end}} {{.pak}}Obj.{{$v.Handle}}){{end}}{{end}}
    }
}
