package rbac

import "github.com/gin-gonic/gin"

// @Summary 获取权限
// @tags rbac
// @Produce  json
// @Param Token header string true "Token"
// @Param test query string true "Test"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /admin/permission [get]
func GetPermission(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"message": "test",
	})
}

// @Summary 添加权限
// @tags rbac
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param test formData string true "Test"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /admin/permission [post]
func PostPermission(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"message": "test",
	})
}
