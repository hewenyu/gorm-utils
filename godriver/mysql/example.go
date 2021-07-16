package mysql

import (
	"log"

	"github.com/hewenyu/gorm-utils/sample"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	db               *gorm.DB
	modelWithHistory = []interface{}{&sample.UserInfo{}}
)

func init() {
	db = NewConn()
}

func NewConn() *gorm.DB {
	_pg_config := MYSQL{
		Name:        viper.GetString("mysql.name"),
		Password:    viper.GetString("mysql.password"),
		User:        viper.GetString("mysql.user"),
		Port:        viper.GetString("mysql.port"),
		Host:        viper.GetString("mysql.host"),
		TablePrefix: viper.GetString("mysql.table"),     // defult dev_
		ParseTime:   viper.GetString("mysql.parsetime"), // defult True
		Loc:         viper.GetString("mysql.loc"),       // defult Local
	}
	return _pg_config.NewConnection()
}

/**
 * GetDB
 */
func GetDB() *gorm.DB {

	sqlDirver, err := db.DB()

	if err != nil {
		log.Println(err.Error())
	}

	if err := sqlDirver.Ping(); err != nil {
		sqlDirver.Close()
		db = NewConn()
	}
	return db
}

/*
setupDatabase 为了使用一些常用组建
*/
func SetupDatabase(db *gorm.DB) {
	db.Exec("create extension IF NOT EXISTS hstore;")
	// 为了使用uuid
	db.Exec("create extension IF NOT EXISTS \"uuid-ossp\"")
	err := db.AutoMigrate(modelWithHistory...)

	if err != nil {
		log.Fatal(err.Error())
	}
}
