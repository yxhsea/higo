package rbac

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mikespook/gorbac"
)

// @Summary 鉴定权限
// @tags rbac
// @Produce  json
// @Param test query string true "Test"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /admin/check/access [get]
func CheckAccess(ctx *gin.Context) {
	urlPath := ctx.Request.URL.Path
	fmt.Printf("urlPath %v \n", urlPath)

	// map[RoleId]PermissionIds
	var jsonRoles map[string][]string

	jsonRoles = map[string][]string{
		"editor": []string{"add-text", "edit-text", "insert-photo", "/v1/admin/check/access"},
	}

	rbac := gorbac.New()
	permissions := make(gorbac.Permissions)

	// Build roles and add them to goRBAC instance
	for rid, pids := range jsonRoles {
		role := gorbac.NewStdRole(rid)
		for _, pid := range pids {
			_, ok := permissions[pid]
			if !ok {
				permissions[pid] = gorbac.NewStdPermission(pid)
			}
			role.Assign(permissions[pid])
		}
		rbac.Add(role)
	}

	_, ok := permissions[urlPath]
	if ok {
		if rbac.IsGranted("editor", permissions[urlPath], nil) {
			fmt.Println("editor allow visits ~")
		} else {
			fmt.Println("editor not allow visits ===")
		}
	} else {
		fmt.Println("editor not allow visits ~")
	}

	ctx.JSON(200, gin.H{
		"message": "test",
	})
}
