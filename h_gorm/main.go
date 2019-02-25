package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"higo/h_gorm/pool"
)

func main() {
	// init database connection.
	pool.InitDB()

	// start http service
	r := gin.New()
	r.GET("/test", func(c *gin.Context) {
		xClientKey := "a_40febbc165577f48fc7d75eaabda5026647b4f"
		token := "431131"
		authId := "+8618373288694"

		var action int64
		var accountId sql.NullInt64
		row := pool.DBPool.Table("sms_tokens").
			Where("app_key = ?  AND token = ? AND auth_id = ? AND status = ?", xClientKey, token, authId, 1).
			Order("created_at desc").
			Limit(1).
			Select("action, account_id").
			Row()
		err := row.Scan(&action, &accountId)
		if err != nil {
			err = fmt.Errorf("query information through token failure, error: (%v), extra params: [xClientKey: (%v), token: (%v), authId: (%v)]", err.Error(), xClientKey, token, authId)
			println(err.Error())
			return
		}
		fmt.Printf("%+v", accountId)

	})

	err := r.Run(":8081")
	if err != nil {
		panic(err)
	}
}
