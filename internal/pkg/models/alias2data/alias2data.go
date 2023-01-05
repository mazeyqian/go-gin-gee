package alias2data

import (
	"time"

	"github.com/mazeyqian/go-gin-gee/internal/pkg/models"
)

type Alias2data struct {
	models.Model
	Alias  string `gorm:"column:alias;not null;" json:"alias" form:"alias"`
	Data   string `gorm:"column:data;not null;" json:"data" form:"data"`
	Public bool   `gorm:"column:public;not null;" json:"public" form:"public"`
}

func (m *Alias2data) BeforeCreate() error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

func (m *Alias2data) BeforeUpdate() error {
	m.UpdatedAt = time.Now()
	return nil
}
