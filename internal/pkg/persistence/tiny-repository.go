package persistence

import (
	"errors"
	"fmt"
	"log"

	models "github.com/mazeyqian/go-gin-gee/internal/pkg/models/tiny"
	"github.com/mazeyqian/go-gin-gee/pkg/helpers"
	"github.com/takuoki/clmconv"
)

type TinyRepository struct{}

var tinyRepository *TinyRepository

func GetTinyRepository() *TinyRepository {
	if tinyRepository == nil {
		tinyRepository = &TinyRepository{}
	}
	return tinyRepository
}

func (t *TinyRepository) SaveOriLink(OriLink string) (*models.Tiny, error) {
	OriMd5 := helpers.ConvertStringToMD5Hash(OriLink)
	data, err := t.QueryOriLinkByOriMd5(OriMd5)
	if err != nil {
		log.Println("SaveOriLink error:", err)
		return nil, err
	}
	if data != nil {
		log.Println("Tiny Exist", data)
		return data, nil // errors.New("data exist")
	}
	var tiny models.Tiny
	tiny.OriLink = OriLink
	tiny.OriMd5 = OriMd5
	err = Create(&tiny)
	if err != nil {
		return nil, err
	}
	err = Save(&tiny)
	if err != nil {
		return nil, err
	}
	TinyId := tiny.ID
	// https://github.com/takuoki/clmconv
	converter := clmconv.New(clmconv.WithStartFromOne(), clmconv.WithLowercase())
	TinyKey := converter.Itoa(int(TinyId))
	TinyLink := fmt.Sprintf("https://feperf.com/t/%s", TinyKey) // `${domain}/t/${tiny_key}`;
	_, err = t.SaveTinyLink(TinyId, TinyLink, TinyKey)
	if err != nil {
		return nil, err
	}
	tiny.TinyKey = TinyKey
	tiny.TinyLink = TinyLink
	return &tiny, err
}

func (t *TinyRepository) QueryOriLinkByTinyKey(TinyKey string) (string, error) {
	var tiny models.Tiny
	var err error
	where := models.Tiny{}
	where.TinyKey = TinyKey
	log.Println("Tiny where:", where)
	notFound, err := First(&where, &tiny, []string{})
	log.Println("Tiny notFound:", notFound)
	if err != nil {
		log.Println("Tiny error:", err)
		return "", err
	}
	log.Println("Tiny QueryOriLinkByTinyKey:", tiny)
	return tiny.OriLink, err
}

func (t *TinyRepository) QueryOriLinkByOriMd5(OriMd5 string) (*models.Tiny, error) {
	log.Println("Tiny OriMd5:", OriMd5)
	if OriMd5 == "" {
		return nil, errors.New("OriMd5 is required")
	}
	var tiny models.Tiny
	where := models.Tiny{}
	where.OriMd5 = OriMd5
	log.Println("Tiny where:", where)
	notFound, err := First(&where, &tiny, []string{})
	log.Println("Tiny notFound:", notFound)
	if err != nil {
		log.Println("Tiny error:", err)
		return nil, err
	}
	return &tiny, err
}

func (t *TinyRepository) SaveTinyLink(TinyId uint64, TinyLink string, TinyKey string) (bool, error) {
	var tiny models.Tiny
	var err error
	where := models.Tiny{}
	where.ID = TinyId
	tiny.TinyLink = TinyLink
	tiny.TinyKey = TinyKey
	err = Updates(&where, &tiny)
	if err != nil {
		return false, err
	}
	return true, err
}
