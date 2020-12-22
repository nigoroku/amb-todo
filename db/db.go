package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
	"github.com/volatiletech/sqlboiler/boil"
)

type DbEnv struct {
	DbURL      string `envconfig:"DB_URL" default:"ambitious-db.cjgaxykqszj0.us-east-1.rds.amazonaws.com"`
	DbPort     string `envconfig:"DB_PORT" default:"3306"`
	DbName     string `envconfig:"DB_NAME" default:"ambitious"`
	DbUser     string `envconfig:"DB_USER" default:"moizumi"`
	DbPassword string `envconfig:"DB_PASSWORD" default:"base0210"`
}

func Init() {
	var dbEnv DbEnv
	envconfig.Process("", &dbEnv)

	// DB接続
	db, err := sql.Open("mysql", dbEnv.DbUser+":"+dbEnv.DbPassword+"@tcp("+dbEnv.DbURL+":"+dbEnv.DbPort+")/"+dbEnv.DbName+"?parseTime=true")
	if err != nil {
		log.Fatalf("Cannot connect database: %v", err)
	}
	boil.SetDB(db)
}
