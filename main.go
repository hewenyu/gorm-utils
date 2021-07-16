package main

import (
	"fmt"

	_ "github.com/hewenyu/gorm-util/config"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println(viper.GetString("mysql.name"))
	fmt.Println(viper.GetString("mysql.password"))
	fmt.Println(viper.GetString("mysql.user"))
	fmt.Println(viper.GetString("mysql.host"))
	fmt.Println(viper.GetString("mysql.table"))
	fmt.Println(viper.GetString("mysql.parsetime"))
	fmt.Println(viper.GetString("mysql.parsetime"))
}
