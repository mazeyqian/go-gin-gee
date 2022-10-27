package persistence

import (
	"errors"
	"log"

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

func (a *Alias2dataRepository) Get(alias string) (*models.Alias2data, error) {
	log.Println("Get alias", alias)
	if alias == "" {
		return nil, errors.New("alias is required")
	}
	var alias2data models.Alias2data
	where := models.Alias2data{}
	where.Alias = alias // strconv.ParseUint(id, 10, 64)
	log.Println("Get where", where)
	notFound, err := First(&where, &alias2data, []string{})
	log.Println("Get notFound", notFound)
	if err != nil {
		log.Println("Get err", err)
		return nil, err
	}
	return &alias2data, err
}

func (a *Alias2dataRepository) Add(alias2data *models.Alias2data) error {
	data, _ := a.Get(alias2data.Alias)
	if data != nil {
		log.Println("Exist", data)
		return errors.New("data exist")
	}
	err := Create(&alias2data)
	if err != nil {
		return err
	}
	err = Save(&alias2data)
	return err
}
