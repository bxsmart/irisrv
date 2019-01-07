package entity

import (
	"time"
)

type BasicEntity struct {
	CreatedAt time.Time `xorm: "create_at"` // 创建时间
	UpdatedAt time.Time `xorm: "update_at"` // 更新时间
}

func (e *BasicEntity) BeforeInsert() {
	e.CreatedAt = time.Now()
}

func (e *BasicEntity) BeforeUpdate() {
	e.UpdatedAt = time.Now()
}
