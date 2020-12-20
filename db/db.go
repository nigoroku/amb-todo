package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
	"github.com/volatiletech/sqlboiler/boil"
)

type DbEnv struct {
	DbAddress  string `default:"localhost"`
	DbPort     string `default:"3306"`
	DbName     string `default:"ambitious"`
	DbUser     string `default:"root"`
	DbPassWord string `default:"base0210"`
}

func Init() {
	var dbEnv DbEnv
	envconfig.Process("", &dbEnv)

	// DB接続
	db, err := sql.Open("mysql", dbEnv.DbUser+":"+dbEnv.DbPassWord+"@tcp("+dbEnv.DbAddress+":"+dbEnv.DbPort+")/"+dbEnv.DbName+"?parseTime=true")
	if err != nil {
		log.Fatalf("Cannot connect database: %v", err)
	}
	boil.SetDB(db)
}
