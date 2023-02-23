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

var {{.higherDir}}Obj *{{.higherDir}}.{{.pak}}

func init(){
    instance = {{.higherDir}}.New{{.pak}}Handle()
    Register(RegisterRoute{
    	Ro: {{.higherDir}}Obj,
        Do: func(e *gin.Engine) { {{.pak}}Router(e) },
    })
}

func {{.funcName}}(c *gin.Engine){
    group := c.Group("/{{.group}}/")
    { {{range $k,$v := .routers}}{{if and $v.Method  $v.Route}}
       // {{$v.Doc}}
       group.{{$v.Method}}("{{$v.Route}}", {{if $v.Middleware }}{{$v.Middleware}},{{end}} {{.higherDir}}Obj.{{$v.Handle}}){{end}}{{end}}
    }
}
