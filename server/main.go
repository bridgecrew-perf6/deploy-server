package main

import (
	_ "github.com/go-sql-driver/mysql"
	"oneclick/server/mysql"
	"oneclick/server/route"
)

func main() {
	db := mysql.GetDB()
	defer db.Close()
	route.CollectRoute()
}
