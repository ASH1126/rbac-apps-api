package db

import (
	"fmt"
	"log"
	"rbac-api/config"
	"time"

	_ "github.com/lib/pq"

	"xorm.io/xorm"
)

func ConnectDB() *xorm.Engine {
	config := config.LoadConfig(".")
	engine, err := xorm.NewEngine("postgres", config.DbDsn)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	loc, _ := time.LoadLocation("Asia/Jakarta")
	engine.TZLocation = loc
	engine.DatabaseTZ = loc
	engine.ShowSQL()
	err = engine.Ping()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	fmt.Println("Connected DB Success")
	return engine
}
