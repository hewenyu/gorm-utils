package godriver

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

/**
 * POSTGRES
 * 默认PG的配置文件
 */
type POSTGRES struct {
	Name        string `json:"name"`         // 数据库名称
	Password    string `json:"password"`     // 数据库密码
	User        string `json:"user"`         // 数据库用户
	Host        string `json:"host"`         // host
	Port        string `json:"port"`         // 端口
	TablePrefix string `json:"table_prefix"` // 标前缀
	SSLMODE     string `json:"ssl_mode"`     // ssl mode
}

/**
 * NewPOSTGRES
 * 初始化PG dsn
 */
func NewPOSTGRES() *POSTGRES {
	return &POSTGRES{
		Name:        os.Getenv("POSTGRES_DB"),
		Password:    os.Getenv("POSTGRES_PWD"),
		User:        os.Getenv("POSTGRES_USER"),
		Host:        os.Getenv("POSTGRES_HOST"),
		Port:        os.Getenv("POSTGRES_PORT"),
		TablePrefix: os.Getenv("POSTGRES_PER"),
		SSLMODE:     os.Getenv("POSTGRES_SSLMODE"),
	}
}

/**
 * NewConnection
 * 创建链接
 */
func (p *POSTGRES) NewConnection() *gorm.DB {
	db_url := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=Asia/Shanghai",
		p.Host,
		p.User,
		p.Password,
		p.Name,
		p.Port,
		p.SSLMODE,
	)

	conn, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  db_url,
		PreferSimpleProtocol: true,
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
 * NewConnection
 * 初始化 链接
 */
func NewConnection() *gorm.DB {
	_pg_config := NewPOSTGRES() // 初始化 pg
	_db := _pg_config.NewConnection()

	return _db
}
