package tiny

import (
	"time"

	"github.com/mazeyqian/go-gin-gee/internal/pkg/models"
)

type Tiny struct {
	models.Model
	OriLink  string `gorm:"column:ori_link;not null;" json:"ori_link" form:"ori_link"`
	OriMd5   string `gorm:"column:ori_md5;not null;" json:"ori_md5" form:"ori_md5"`
	TinyLink string `gorm:"column:tiny_link;not null;" json:"tiny_link" form:"tiny_link"`
	TinyKey  string `gorm:"column:tiny_key;not null;" json:"tiny_key" form:"tiny_key"`
}

func (m *Tiny) BeforeCreate() error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

func (m *Tiny) BeforeUpdate() error {
	m.UpdatedAt = time.Now()
	return nil
}
