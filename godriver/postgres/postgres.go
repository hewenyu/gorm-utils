package postgres

import (
	"log"

	"gorm.io/gorm"
)

func ResetDatabase(db *gorm.DB) {
	CleanDatabase(db)
	SetupDatabase(db)
}

/*
cleanDatabase 删除表
*/
func CleanDatabase(db *gorm.DB) {
	err := db.Migrator().DropTable(modelWithHistory...)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// /*
// setupDatabase 为了使用一些常用组建
// */
// func SetupDatabase(db *gorm.DB) {
// 	db.Exec("create extension IF NOT EXISTS hstore;")
// 	// 为了使用uuid
// 	db.Exec("create extension IF NOT EXISTS \"uuid-ossp\"")
// 	err := db.AutoMigrate(modelWithHistory...)

// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// }
