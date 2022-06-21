package model

import (
	"github.com/gin-gonic/gin"
	"time"
)

type User struct {
	ID        uint `json:"id" gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	SN        string `json:"sn" gorm:"type:char(16);unique;force"`
	WgKey     string `json:"wg_key" gorm:"type:varchar(120);not null;unique"`
}

type T struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Sign string `json:"sign"`
	Data struct {
		Content string `json:"content"`
	} `json:"data"`
}

type RestController interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Show(c *gin.Context)
	Delete(c *gin.Context)
}
