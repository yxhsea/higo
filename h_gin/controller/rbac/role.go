package rbac

import "github.com/gin-gonic/gin"

// @Summary 获取角色
// @tags rbac
// @Produce  json
// @Param Token header string true "Token"
// @Param test query string true "Test"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /admin/role [get]
func GetRole(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"message": "test",
	})
}

// @Summary 添加角色
// @tags rbac
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param test formData string true "Test"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /admin/role [post]
func PostRole(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"message": "test",
	})
}
