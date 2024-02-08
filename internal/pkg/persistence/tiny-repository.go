package persistence

import (
	"errors"
	"fmt"
	"log"
	"os"

	models "github.com/mazeyqian/go-gin-gee/internal/pkg/models/tiny"
	"github.com/mazeyqian/go-gin-gee/pkg/helpers"
	"github.com/mazeyqian/gurl"
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

func (r *TinyRepository) SaveOriLink(OriLink string, addBaseUrl string) (string, error) {
	var err error
	var tiny models.Tiny
	var linkForEncode string
	if addBaseUrl != "" {
		linkForEncode, err = gurl.SetHashParam(OriLink, "base_url", addBaseUrl)
		log.Println("linkForEncode:", linkForEncode)
		if err != nil {
			return "", err
		}
	} else {
		linkForEncode = OriLink
	}
	OriMd5 := helpers.ConvertStringToMD5Hash(linkForEncode)
	data, _ := r.QueryOriLinkByOriMd5(OriMd5)
	if data != nil {
		log.Println("Tiny Exist:", data)
		return data.TinyLink, nil
	}
	baseUrl := os.Getenv("BASE_URL")
	if addBaseUrl != "" {
		baseUrl = addBaseUrl
	}
	if baseUrl == "" {
		return "", errors.New("BASE_URL is required")
	}
	tiny.OriLink = OriLink
	tiny.OriMd5 = OriMd5
	err = Create(&tiny)
	if err != nil {
		return "", err
	}
	err = Save(&tiny)
	if err != nil {
		return "", err
	}
	TinyId := tiny.ID
	// https://github.com/takuoki/clmconv
	converter := clmconv.New(clmconv.WithStartFromOne(), clmconv.WithLowercase())
	TinyKey := converter.Itoa(int(TinyId))
	TinyLink := fmt.Sprintf("%s/t/%s", baseUrl, TinyKey)
	_, err = r.SaveTinyLink(TinyId, TinyLink, TinyKey)
	if err != nil {
		return "", err
	}
	tiny.TinyKey = TinyKey
	tiny.TinyLink = TinyLink
	return tiny.TinyLink, err
}

func (r *TinyRepository) QueryOriLinkByTinyKey(TinyKey string) (string, error) {
	var tiny models.Tiny
	var err error
	where := models.Tiny{}
	where.TinyKey = TinyKey
	// log.Println("Tiny where:", where)
	notFound, err := First(&where, &tiny, []string{})
	// log.Println("Tiny notFound:", notFound)
	log.Printf("Tiny notFound: %t", notFound)
	if err != nil {
		log.Printf("Tiny error: %v", err)
		return "", err
	}
	go r.RecordVisitCountByTinyKey(TinyKey)
	log.Printf("Tiny QueryOriLinkByTinyKey: %s", tiny.OriLink)
	return tiny.OriLink, err
}

func (r *TinyRepository) RecordVisitCountByTinyKey(TinyKey string) (bool, error) {
	var tiny models.Tiny
	var err error
	where := models.Tiny{}
	where.TinyKey = TinyKey
	notFound, err := First(&where, &tiny, []string{})
	if err != nil {
		return false, err
	}
	if notFound {
		return false, errors.New("link not found")
	}
	tiny.VisitCount = tiny.VisitCount + 1
	err = Updates(&where, &tiny)
	if err != nil {
		return false, err
	}
	log.Printf("Tiny Current Count: %d", tiny.VisitCount)
	return true, err
}

func (r *TinyRepository) QueryOriLinkByOriMd5(OriMd5 string) (*models.Tiny, error) {
	var tiny models.Tiny
	log.Println("Tiny OriMd5:", OriMd5)
	if OriMd5 == "" {
		return nil, errors.New("OriMd5 is required")
	}
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

func (r *TinyRepository) SaveTinyLink(TinyId uint64, TinyLink string, TinyKey string) (bool, error) {
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
