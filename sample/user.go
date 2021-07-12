package sample

import (
	"time"

	"github.com/gofrs/uuid"
)

/**
 * UserInfo 用户信息表 测试使用
 */
type UserInfo struct {
	ID        uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4();comment:主键id"`
	Name      string    `json:"name" gorm:"type:varchar(50);not null;default:'';comment:用户名"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp without time zone;not null;default:now();comment:用户创建时间"`
}
