package persistence

import (
	models "github.com/mazeyqian/go-gin-gee/internal/pkg/models/alias2data"
)

type Alias2dataRepository struct{}

var alias2dataRepository *Alias2dataRepository

func GetAlias2dataRepository() *Alias2dataRepository {
	if alias2dataRepository == nil {
		alias2dataRepository = &Alias2dataRepository{}
	}
	return alias2dataRepository
}

func (r *Alias2dataRepository) Get(alias string) (*models.Alias2data, error) {
	var alias2data models.Alias2data
	where := models.Alias2data{}
	where.Alias = alias // strconv.ParseUint(id, 10, 64)
	_, err := First(&where, &alias2data, []string{})
	if err != nil {
		return nil, err
	}
	return &alias2data, err
}

func (r *Alias2dataRepository) Add(alias2data *models.Alias2data) error {
	err := Create(&alias2data)
	if err != nil {
		return err
	}
	err = Save(&alias2data)
	return err
}
