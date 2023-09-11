package tiny

import (
	"time"

	"github.com/mazeyqian/go-gin-gee/internal/pkg/models"
)

type Tiny struct {
	models.Model
	OriLink  string `gorm:"column:ori_link;not null;size:350" json:"ori_link" form:"ori_link"`
	OriMd5   string `gorm:"column:ori_md5;not null;size:40" json:"ori_md5" form:"ori_md5"`
	TinyLink string `gorm:"column:tiny_link;not null;size:30" json:"tiny_link" form:"tiny_link"`
	TinyKey  string `gorm:"column:tiny_key;not null;size:20" json:"tiny_key" form:"tiny_key"`
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
