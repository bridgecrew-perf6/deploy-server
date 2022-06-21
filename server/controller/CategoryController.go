package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"oneclick/server/Signature"
	"oneclick/server/controller/methods"
	"oneclick/server/model"
	"oneclick/server/mysql"
	"oneclick/server/pkg"
	"oneclick/server/response"
	"oneclick/server/returned"
	"oneclick/server/utils"
	"runtime"
	"strconv"
)

type Processor struct {
	db *gorm.DB
}

func (p Processor) Info(c *gin.Context) {
	var data pkg.PingReturned
	err := c.ShouldBindJSON(&data)
	if err != nil {
		methods.ErrorDesc(c, err.Error())
		return
	}
	m, _ := json.Marshal(&data.Data)
	//验签
	solution := make(map[string]string)
	solution["time"] = strconv.Itoa(data.Time)
	solution["data"] = string(m)
	signature := Signature.SignatureMethod{}.KuratowskiConstraint(solution, data.Sign)
	if !signature {
		//返回签名错误
		response.SignError(c)
		return
	}
	res, _ := json.Marshal(&data)
	fmt.Println(string(res))
	c.JSON(200, "ok")
	fmt.Println(runtime.NumGoroutine())
}

func (p Processor) Create(c *gin.Context) {
	SN := utils.RandomString(16)
	user := model.User{
		SN: SN,
	}
	if err := p.db.Create(&user).Error; err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func (p Processor) Update(c *gin.Context) {
	var data pkg.ClientBinding
	err := c.ShouldBindJSON(&data)
	if err != nil {
		response.Fail(c, "请求失败", gin.H{
			"content": err.Error(),
		})
		return
	}
	solution := make(map[string]string)
	solution["time"] = strconv.Itoa(data.Time)
	solution["sign"] = data.Sign
	solution["domain"] = data.Domain
	solution["networklaundry"] = strconv.Itoa(data.Networklaundry)
	solution["key"] = data.Key
	//验签
	signature := Signature.SignatureMethod{}.KuratowskiConstraint(solution, solution["sign"])
	if !signature {
		//response.SignError(c)
		//return
	}
	catKey, err := p.SelectKey(data.Key)
	if err != nil {
		response.Fail(c, "密钥不存在", gin.H{})
		return
	}
	if catKey.SN != "" {
		methods.ErrorDesc(c, "密钥已绑定")
		return
	}
	fmt.Println(catKey)
	SN := utils.RandomString(16)
	if err = p.db.Model(&catKey).Update("sn", SN).Error; err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(data)
	switch {
	case data.Networklaundry == 1:
		c.JSON(200, returned.SuccessWg(1, "请求成功", utils.WireguardConfig, utils.FrpcConfig))
	case data.Networklaundry == 2:
		c.JSON(200, returned.ReturnUdp2raw(1, "请求成功", 4000, "127.0.0.1:50000", "speedbox", utils.WireguardConfig, utils.FrpcConfig))
	}
}

func (p Processor) Show(c *gin.Context) {
	var data pkg.ClientBinding
	err := c.ShouldBindJSON(&data)
	if err != nil {
		//返回错误
		response.Fail(c, "请求失败", gin.H{
			"content": err.Error(),
		})
		return
	}
	catKey, err := p.SelectKey(data.Key)
	if err != nil {
		fmt.Println(err.Error())
		response.Fail(c, "密钥不存在", gin.H{})
		return
	}
	if catKey.WgKey == "" {
		response.Fail(c, "密钥不存在", gin.H{})
		return
	}
	response.Success(c, "密钥存在", gin.H{})
}

func (p Processor) Delete(c *gin.Context) {
	var data pkg.ClientBinding
	err := c.ShouldBindJSON(&data)
	if err != nil {
		//返回错误
		response.Fail(c, "密钥不存在", gin.H{})
		return
	}
	//验签
	solution := make(map[string]string)
	solution["time"] = strconv.Itoa(data.Time)
	solution["key"] = data.Key
	signature := Signature.SignatureMethod{}.KuratowskiConstraint(solution, data.Sign)
	if !signature {
		//返回签名错误
		response.SignError(c)
		return
	}
	catKey, err := p.SelectKey(data.Key)
	if err != nil {
		fmt.Println(err.Error())
		response.Fail(c, "密钥不存在", gin.H{})
		return
	}
	//if err := p.db.Where(&model.User{WgKey: data.Key}).Delete(model.User{}).Error; err != nil {
	//	response.Fail(c, "解绑失败，请重试", gin.H{})
	//	return
	//}
	if err = p.db.Model(&catKey).Update("sn", nil).Error; err != nil {
		response.Fail(c, "解绑失败，请重试", gin.H{})
		return
	}
	response.Success(c, "解绑成功", gin.H{})
}

func (p Processor) SelectKey(key string) (catKey *model.User, err error) {
	catKey = &model.User{}
	if err = p.db.First(&catKey, "wg_key = ?", key).Error; err != nil {
		return nil, err
	}
	return catKey, nil
}

func (p Processor) DeleteById(key string) error {
	if err := p.db.Delete(model.User{}, key).Error; err != nil {
		return err
	}
	return nil
}

func NewIPostController() Processor {
	return Processor{
		db: mysql.GetDB(),
	}
}
