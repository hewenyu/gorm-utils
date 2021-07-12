package sample

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"gorm.io/gorm/clause"
)

func RandomString(n int) string {

	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	length := len(letter)

	for i := range b {
		b[i] = letter[rand.Intn(length)]
	}

	return string(b)
}

func NonceStr() string {

	return RandomString(32)
}

func Test_demo(t *testing.T) {
	_db := GetDB()

	for {

		_new_user := UserInfo{
			Name: NonceStr(),
		}

		// 更新用户字段
		_db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},              // key colume
			DoUpdates: clause.AssignmentColumns([]string{"name"}), // column needed to be updated
		}).Create(&_new_user)

		db.Commit()

		time.Sleep(time.Second * 10)

		fmt.Println("db commit")

	}

}
