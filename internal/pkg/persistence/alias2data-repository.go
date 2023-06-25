package persistence

import (
	"errors"
	"log"
	"strconv"

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

func (r *Alias2dataRepository) Add(alias2data *models.Alias2data) error {
	data, _ := r.Get(alias2data.Alias)
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

func (r *Alias2dataRepository) CountByAlias(alias string) (int, error) {
	if alias == "" {
		return 0, errors.New("alias is required")
	}

	alias2data := models.Alias2data{Alias: alias}
	notFound, err := First(&alias2data, &alias2data, []string{})
	if err != nil {
		log.Println("First err", err)
		// return 0, err
	}
	if notFound {
		// log.Println("First notFound", notFound)
		alias2data.Data = "0"
		err = Create(&alias2data)
		if err != nil {
			return 0, err
		}
	}

	count, err := strconv.Atoi(alias2data.Data)
	if err != nil {
		// log.Println("Atoi err", err)
		return 0, err
	}

	count++
	alias2data.Data = strconv.Itoa(count)

	err = Save(&alias2data)
	if err != nil {
		// log.Println("Save err", err)
		return 0, err
	}

	return count, nil
}

// With this query, you can retrieve the 10 aliases with the lowest counts from the last 2 months.
// SELECT alias, MAX(created_at) AS last_count_date, SUM(data) AS count
// FROM alias2data
// WHERE created_at >= '2023-03-16 00:00:00'
// GROUP BY alias
// ORDER BY count ASC
// LIMIT 10;
