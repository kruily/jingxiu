/**
* @file: mapping.go ==> jx_api
* @package: jx_api
* @author: jingxiu
* @since: 2022/11/7
* @desc: 仅用作接口函数的的映射
 */

package jx_api

var (
	// APIMatchMapping API文件扫描匹配字段
	APIMatchMapping = []string{"@Group", "@Route", "@Method", "@Middleware", "@Doc"}
	// APIMiddlewareMapping api路由中间件匹配
	APIMiddlewareMapping = map[string]string{
		"JWT":  "middleware.JWTAuth()",
		"Auth": "middleware.UserAuth()",
	}
	// APIHandleMapping 应用生成可提供的接口
	APIHandleMapping = map[string][]string{
		"APIHandler": {
			"Create", "List", "Info", "Delete", "Update", "Status",
		},
	}
)

func AppendMapping() {

}
