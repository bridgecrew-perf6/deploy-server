package methods

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oneclick/server/Signature"
	"oneclick/server/returned"
	"strconv"
)

type ReturnController struct {
	RuleReturnScope
}

type RuleReturnScope interface {
	ErrorDesc(c *gin.Context)
	SignError(c *gin.Context)
}

func VerificationSignatureData(c *gin.Context, time int, data string, sign string) bool {
	solution := make(map[string]string)
	solution["time"] = strconv.Itoa(time)
	solution["sign"] = sign
	solution["data"] = data
	signature := Signature.SignatureMethod{}.KuratowskiConstraint(solution, solution["sign"])
	if !signature {
		//返回签名错误
		SignError(c)
		return false
	}
	return true
}

func ErrorDesc(c *gin.Context, msg string) {
	//返回错误
	ers := returned.ReturnErrors(0, msg)
	c.JSON(http.StatusOK, ers)
}

func ErrorDescs(c *gin.Context, msg, content string) {
	//返回错误
	ers := returned.ReturnErrors(2, msg)
	c.JSON(http.StatusOK, ers)
}

func SignError(c *gin.Context) {
	c.JSON(http.StatusBadRequest, returned.ReturnErrors(0, "签名错误"))
}
