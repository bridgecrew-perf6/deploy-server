package processor

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"oneclick/server/model"
	"oneclick/server/mysql"
)

type Processor struct {
	db *gorm.DB
}

type IPostController interface {
	model.RestController
}

func NewIPostController() Processor {
	return Processor{
		db: mysql.GetDB(),
	}
}

func (p Processor) Create() error {
	//SN := utils.RandomString(16)
	//user := model.User{
	//	SN:  SN,
	//	Key: SN,
	//}
	//fmt.Println(user)
	//fmt.Println(SN)
	//if err := p.db.Create(&user); err != nil {
	//	fmt.Println(err)
	//	//panic(err)
	//	return nil
	//}
	//p.db.Create(&model.User{
	//	SN:  "sthrert",
	//	Keys: "whrwh",
	//})
	return nil
}

func (p Processor) Update(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (p Processor) Show(key string) bool {
	user := model.User{}
	p.db.Where("keys = ?", key).First(&user)
	if user.ID != 0 {
		return true
	}
	fmt.Println(user)
	return false
}

func (p Processor) Delete(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

//func init() {
//	db = mysql.GetDB()
//}

//func SelectMysql(keys string) (sn, key string, err error) {
//	selectUser, err := DB.Query("select user_key, sn  from user where user_key = ?", keys)
//	if err != nil {
//		fmt.Println(err.Error())
//		log.Info(err.Error())
//		return "", "", err
//	}
//	fmt.Println(DB.Stats())
//	selectUser.Next()
//	selectUser.Scan(&key, &sn)
//	selectUser.Close()
//	log.SetFormatter(&log.TextFormatter{
//		DisableColors: false,
//		FullTimestamp: true,
//	})
//	log.SetReportCaller(true)
//	log.Info("fgghvjbbnkfjnbdskhndgf")
//	log.Error("fgghvjbbnkfjnbdskhndgf")
//	return sn, key, nil
//}

//func (p Processor) insertData(key string) {
//	DB := mysql.GetDB()
//	SN := utils.RandomString(16)
//	fmt.Println(SN)
//	result, err := DB.Exec("UPDATE user SET sn = ? WHERE user_key = ?;", SN, key)
//	if err != nil {
//		fmt.Printf("Insert failed,err:%v", err)
//		return
//	}
//	defer DB.Close()
//	lastInsertID, err := result.LastInsertId() //插入数据的主键id
//	if err != nil {
//		fmt.Printf("Get lastInsertID failed,err:%v", err)
//		return
//	}
//	fmt.Println("LastInsertID:", lastInsertID)
//	rowsaffected, err := result.RowsAffected() //影响行数
//	if err != nil {
//		fmt.Printf("Get RowsAffected failed,err:%v", err)
//		return
//	}
//	fmt.Println("RowsAffected:", rowsaffected)
//}

//
//func (p Processor) DeleteData(key string) bool {
//	DB := mysql.GetDB()
//	result, err := DB.Exec("UPDATE user SET sn = NULL WHERE user_key = ?;", key)
//	if err != nil {
//		fmt.Printf("Insert failed,err:%v", err)
//		return false
//	}
//	defer DB.Close()
//	//插入数据的主键id
//	lastInsertID, err := result.LastInsertId()
//	if err != nil {
//		fmt.Printf("Get lastInsertID failed,err:%v", err)
//		return false
//	}
//	fmt.Println("LastInsertID:", lastInsertID)
//	//影响行数
//	rowsaffected, err := result.RowsAffected()
//	if err != nil {
//		fmt.Printf("Get RowsAffected failed,err:%v", err)
//		return false
//	}
//	fmt.Println("RowsAffected:", rowsaffected)
//	return true
//}
