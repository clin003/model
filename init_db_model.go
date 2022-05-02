package model

import (
	"time"

	"gorm.io/gorm"
)

// type BaseModel struct {
// 	Id        uint64    `gorm:"column:id" json:"id"`
// 	CreatedAt time.Time `gorm:"column:created_at" json:"-"`
// 	UpdatedAt time.Time `gorm:"column:updated_at" json:"-"`
// 	// DeletedAt *time.Time `gorm:"column:deleted_at" sql:"index" json:"-"`
// }
// type BaseModel struct {
// 	gorm.Model
// }

type Model struct {
	ID        uint64         `json:"id" gorm:"primarykey; autoIncrement"`
	CreatedAt time.Time      `json:"-" form:"-"` //created_at
	UpdatedAt time.Time      `json:"-" form:"-"` //updated_at
	DeletedAt gorm.DeletedAt `json:"-" form:"-" gorm:"index"`
}

type GModel struct {
	GID       uint64         `json:"gid" gorm:"primarykey; autoIncrement"`
	CreatedAt time.Time      `json:"-" form:"-"`
	UpdatedAt time.Time      `json:"-" form:"-"`
	DeletedAt gorm.DeletedAt `json:"-" form:"-" gorm:"index"`
}

// type Model struct {
// 	gorm.Model
// }

// type Model struct {
// 	ID        uint `gorm:"primarykey"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt DeletedAt `gorm:"index"`
// }
