package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"oneclick/server/Signature"
	"oneclick/server/controller/methods"
	"oneclick/server/pkg"
	"oneclick/server/response"
	"strconv"
)

func CheckKey(c *gin.Context) {
	var data pkg.BindingRequest
	err := c.ShouldBindJSON(&data)
	if err != nil {
		//返回错误
		//ers := returned.ReturnErrors(0, "请求失败", err.Error())
		response.Fail(c, "请求失败", gin.H{
			"content": err.Error(),
		})
		return
	} else {
		m, _ := json.Marshal(&data)
		solution := make(map[string]string)
		json.Unmarshal(m, &solution)
		solution["time"] = strconv.Itoa(data.Time)
		solution["sign"] = data.Sign
		solution["key"] = data.Key
		signature := Signature.SignatureMethod{}.KuratowskiConstraint(solution, solution["sign"])
		if !signature {
			//返回签名错误
			methods.SignError(c)
			return
		}
		//_, key, _ := processor.SelectMysql(data.Key)
		//if len(key) == 0 {
		//	ers := returned.ReturnErrors(0, "请求失败", "密钥不存在或已过期")
		//	c.JSON(http.StatusBadRequest, ers)
		//	return
		//}
		//ers := returned.ReturnErrors(1, "请求成功", "密钥存在")
		//c.JSON(http.StatusOK, ers)
		//return
	}
}
