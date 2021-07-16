package sample

import (
	"log"
	"time"

	"github.com/hewenyu/gorm-utils/godriver"
	"gorm.io/gorm"
)

var (
	mysqldb          *gorm.DB
	modelWithHistory = []interface{}{&UserInfo{}}
)

func init() {

	mysqldb = NewMYSQLConn()
	sqlDirver, err := mysqldb.DB()

	if err != nil {
		log.Println(err.Error())
	}

	sqlDirver.SetMaxIdleConns(10)                   //最大空闲连接数
	sqlDirver.SetMaxOpenConns(30)                   //最大连接数
	sqlDirver.SetConnMaxLifetime(time.Second * 300) //设置连接空闲超时

	// defer sqlDirver.Close()

	SetupMysqlDatabase(mysqldb)
}

func NewMYSQLConn() *gorm.DB {

	_pg_config := godriver.MYSQL{
		Name:        "postgres",
		User:        "postgres",
		Host:        "localhost",
		Password:    "example",
		Port:        "5432",
		TablePrefix: "test_",
		// SSLMODE:     "disable",
	}

	return _pg_config.NewConnection()

}

/**
 * GetDB
 */
func GetMYSQLDB() *gorm.DB {

	sqlDirver, err := mysqldb.DB()

	if err != nil {
		log.Println(err.Error())
	}

	if err := sqlDirver.Ping(); err != nil {
		sqlDirver.Close()
		mysqldb = NewMYSQLConn()
	}
	return mysqldb
}

/*
setupDatabase 为了使用一些常用组建
*/
func SetupMysqlDatabase(db *gorm.DB) {
	db.Exec("create extension IF NOT EXISTS hstore;")
	// 为了使用uuid
	db.Exec("create extension IF NOT EXISTS \"uuid-ossp\"")
	err := db.AutoMigrate(modelWithHistory...)

	if err != nil {
		log.Fatal(err.Error())
	}
}
