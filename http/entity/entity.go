package entity

import "time"

type BasicEntity struct {
	CreatedAt time.Time `gorm: "column:create_at"` // 创建时间
	UpdatedAt time.Time `gorm: "column:update_at"` // 更新时间
}

func (e *FCoinAccount) BeforeInsert() {
	e.CreatedAt = time.Now().UTC()
}

func (e *FCoinAccount) BeforeUpdate() {
	e.UpdatedAt = time.Now().UTC()
}
