package sqldriver

import (
	"database/sql/driver"
	"encoding/json"
)

/**
 * JSONB
 * 用于处理对应PG的jsonb格式的数据
 */
type JSONB map[string]interface{}

func (j JSONB) Value() (driver.Value, error) {
	valueString, err := json.Marshal(j)
	return string(valueString), err
}

func (j *JSONB) Scan(value interface{}) error {
	if err := json.Unmarshal(value.([]byte), &j); err != nil {
		return err
	}
	return nil
}
