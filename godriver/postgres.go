package godriver

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db               *gorm.DB
	ModelWithHistory []interface{}
)

func SetUp() {
	db = NewPostgresConnection()

	sqlDirver, err := db.DB()

	if err != nil {
		log.Println(err.Error())
	}

	sqlDirver.SetMaxIdleConns(10)                   //最大空闲连接数
	sqlDirver.SetMaxOpenConns(30)                   //最大连接数
	sqlDirver.SetConnMaxLifetime(time.Second * 300) //设置连接空闲超时

	// defer sqlDirver.Close()
}

/**
 * GetDB 获取db
 */
func GetDB() *gorm.DB {

	sqlDirver, err := db.DB()

	if err != nil {
		log.Println(err.Error())
	}

	if err := sqlDirver.Ping(); err != nil {
		sqlDirver.Close()
		db = NewPostgresConnection()
	}
	return db
}

/**
 *open 链接PG数据库
 *dsn "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
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
