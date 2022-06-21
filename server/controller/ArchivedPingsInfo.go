package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"oneclick/server/controller/methods"
	"oneclick/server/pkg"
	"runtime"
)

func ArchivedPingsInfo(c *gin.Context) {
	var data pkg.PingReturned
	err := c.ShouldBindJSON(&data)
	if err != nil {
		//返回错误
		methods.ErrorDesc(c, err.Error())
		return
	} else {
		m, _ := json.Marshal(&data.Data)
		//验签
		boolean := methods.VerificationSignatureData(c, data.Time, string(m), data.Sign)
		if !boolean {
			return
		}
		res, _ := json.Marshal(&data)
		fmt.Println(string(res))
		c.JSON(200, "ok")
	}
	fmt.Println(runtime.NumGoroutine())
}
