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
// @Tags {{.Package}}
// @Summary
// @Description
// @Accept json
// @Produce  json
// @Router /{{.Package}}/{{.CurrentRoute}} [get]
// @Success 200 {object} result.Res
// @Failure 400 {object} result.Res
// @Middleware []
func (x *{{.Controller}}) {{.CurrentHandle}}(c *gin.Context) {
	// TODO This is where you should write the logic code
}