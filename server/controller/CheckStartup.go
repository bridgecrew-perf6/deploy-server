package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"oneclick/server/Signature"
	"oneclick/server/controller/methods"
	"oneclick/server/pkg"
	"strconv"
)

func CheckStartup(c *gin.Context) {
	var data pkg.StatusByte
	err := c.ShouldBindJSON(&data)
	if err != nil {
		methods.ErrorDesc(c, err.Error())
		return
	}
	solution := make(map[string]string)
	solution["time"] = strconv.Itoa(data.Time)
	solution["sign"] = data.Sign
	solution["frpc"] = strconv.Itoa(data.Frpc.Status)
	solution["udp2raw"] = strconv.Itoa(data.Udp2Raw.Status)
	solution["oneclick"] = strconv.Itoa(data.Oneclick.Status)
	solution["wireguard"] = strconv.Itoa(data.Wireguard.Status)
	signature := Signature.SignatureMethod{}.KuratowskiConstraint(solution, solution["sign"])
	if !signature {
		//返回签名错误
		methods.SignError(c)
		return
	}
	res, _ := json.Marshal(&data)
	fmt.Println(string(res))
	c.JSON(200, "ok")

}
