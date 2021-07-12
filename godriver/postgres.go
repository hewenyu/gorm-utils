package godriver

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB               *gorm.DB
	ModelWithHistory []interface{}
)

/*
init 初始化
*/
func init() {
	// export POSTGRES_HOST=localhost POSTGRES_USER=gorm POSTGRES_PWD=gorm POSTGRES_DB=gorm POSTGRES_PORT=9920 POSTGRES_SSLMODE=disable
	db_url := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=Asia/Shanghai",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PWD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_SSLMODE"),
	)

	DB = OpenPG(db_url)

	SetupDatabase(DB)

	// return
}

/*
open 链接PG数据库
dsn "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
*/
func OpenPG(dsn string) (db *gorm.DB) {

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // 禁用缓存
	}), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())

	}
	return
}

func ResetDatabase(db *gorm.DB) {
	//cleanDatabase(db)
	CleanDatabase(db)
	SetupDatabase(db)
}

/*
cleanDatabase 删除表
*/
func CleanDatabase(db *gorm.DB) {
	err := db.Migrator().DropTable(ModelWithHistory...)
	if err != nil {
		log.Fatal(err.Error())
	}
}

/*
setupDatabase 为了使用一些常用组建
*/
func SetupDatabase(db *gorm.DB) {
	db.Exec("create extension IF NOT EXISTS hstore;")
	// 为了使用uuid
	db.Exec("create extension IF NOT EXISTS \"uuid-ossp\"")
	err := db.AutoMigrate(ModelWithHistory...)

	if err != nil {
		log.Fatal(err.Error())
	}
}
