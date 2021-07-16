package mysql

import (
	"fmt"
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/hewenyu/gorm-utils/sample"
	"github.com/hewenyu/gorm-utils/utils"
	"gorm.io/gorm/clause"
)

func Test_demo(t *testing.T) {
	_db := GetDB()

	for {
		_uid, _ := uuid.NewV4()

		_name := utils.NonceStr()
		// UserInfo :=
		_new_user := sample.UserInfo{
			ID:        _uid,
			Name:      _name,
			CreatedAt: time.Now().Local(),
		}

		// 更新用户字段
		_db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},              // key colume
			DoUpdates: clause.AssignmentColumns([]string{"name"}), // column needed to be updated
		}).Create(&_new_user)

		fmt.Println("db commit")

		time.Sleep(time.Second)

	}

}
