package sample

import (
	"math/rand"
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

// func Test_demo(t *testing.T) {
// 	_db := GetDB()

// 	for {
// 		_uid, _ := uuid.NewV4()

// 		_name := NonceStr()
// 		_new_user := UserInfo{
// 			ID:        _uid,
// 			Name:      _name,
// 			CreatedAt: time.Now().Local(),
// 		}

// 		// 更新用户字段
// 		_db.Clauses(clause.OnConflict{
// 			Columns:   []clause.Column{{Name: "id"}},              // key colume
// 			DoUpdates: clause.AssignmentColumns([]string{"name"}), // column needed to be updated
// 		}).Create(&_new_user)

// 		fmt.Println("db commit")

// 		time.Sleep(time.Second)

// 	}

// }
