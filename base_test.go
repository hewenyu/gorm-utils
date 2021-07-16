package main

import (
	"fmt"
	"testing"

	// "github.com/hewenyu/gorm-utils/config"
	_ "github.com/hewenyu/gorm-utils/config"
	"github.com/spf13/viper"
)

func Test_ddd(t *testing.T) {

	// name := config.GetConfigName()
	fmt.Println(viper.GetString("mysql.name"))
	fmt.Println(viper.GetString("mysql.password"))
	fmt.Println(viper.GetString("mysql.user"))
	fmt.Println(viper.GetString("mysql.host"))
	fmt.Println(viper.GetString("mysql.table"))
	fmt.Println(viper.GetString("mysql.parsetime"))
	fmt.Println(viper.GetString("mysql.parsetime"))
}
