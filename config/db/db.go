package db

import (
	"fmt"
	"log"
	"rbac-api/config"

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
	engine.ShowSQL()
	err = engine.Ping()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	fmt.Println("Connected DB Success")
	return engine
}
