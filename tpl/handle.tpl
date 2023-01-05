/**
* @file: {{.File}}==> {{.Path}}
* @package: router
* @author: jingxiu
* @since: {{.Date}}
* @desc: //Write Doc
 */

package {{.Package}}

import (
    "github.com/gin-gonic/gin"
)
// {{.CurrentHandle}}
// @Handle {{.CurrentHandle}}
// @Summary $(一个简短的操作概述)
// @Description $(操作行为的详细说明)
// @Accept json
// @Produce  json
// @Param $(用空格分隔的请求参数 param nameparam typedata typeis mandatory?comment attribute(optional))
// @Success 200 {object} $(返回数据模型) "请求成功"
// @Failure 400 {object} $(返回数据模型) "请求错误"
// @Router /{{.Package}}/$(请求子路由) [$(请求方法)]
// @Middleware [$(中间件缩写)]
func (x *{{.Controller}}) {{.CurrentHandle}}(c *gin.Context) {
	// TODO This is where you should write the logic code
}