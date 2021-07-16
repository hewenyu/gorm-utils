package mysql

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

/**
 * MYSQL
 * 默认PG的配置文件
 */
type MYSQL struct {
	Name        string `json:"name"`         // 数据库名称
	Password    string `json:"password"`     // 数据库密码
	User        string `json:"user"`         // 数据库用户
	Host        string `json:"host"`         // host
	Port        string `json:"port"`         // 端口
	TablePrefix string `json:"table_prefix"` // 表前缀
	ParseTime   string `json:"parseTime"`    // 时间修改
	Loc         string `json:"loc"`          // defult Local
}

/**
 * NewMYSQL
 * 初始化MYSQL dsn
 */
func NewMYSQL() *MYSQL {
	return &MYSQL{
		Name:        os.Getenv("MYSQL_DB"),
		Password:    os.Getenv("MYSQL_PWD"),
		User:        os.Getenv("MYSQL_USER"),
		Host:        os.Getenv("MYSQL_HOST"),
		Port:        os.Getenv("MYSQL_PORT"),
		TablePrefix: os.Getenv("MYSQL_TablePrefix"),
		ParseTime:   os.Getenv("MYSQL_ParseTime"),
		Loc:         os.Getenv("MYSQL_Loc"),
	}
}

/**
 * NewConnection
 * 创建链接
 *  dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
 */
func (p *MYSQL) NewConnection() *gorm.DB {
	db_url := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=%v&loc=%v",
		p.User,
		p.Password,
		p.Host,
		p.Port,
		p.Name,
		p.ParseTime,
		p.Loc,
	)

	conn, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       db_url, // data source name
		DefaultStringSize:         256,    // default size for string fields
		DisableDatetimePrecision:  true,   // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,   // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,   // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,  // auto configure based on currently MySQL version
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: strings.ToLower(p.TablePrefix), // 表名前缀，`User` 的表名应该是 `t_users`
			// SingularTable: true,                           // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	return conn
}

/**
 * NewMSQLConnection
 * 初始化 链接
 */
func NewMSQLConnection() *gorm.DB {
	_pg_config := NewMYSQL() // 初始化 pg
	_db := _pg_config.NewConnection()

	return _db
}
