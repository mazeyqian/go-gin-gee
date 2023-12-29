package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/mazeyqian/go-gin-gee/internal/pkg/models/tiny"
	"github.com/mazeyqian/go-gin-gee/internal/pkg/persistence"
	http_err "github.com/mazeyqian/go-gin-gee/pkg/http-err"
)

func RedirectTiny(c *gin.Context) {
	s := persistence.GetTinyRepository()
	TinyKey := c.Param("key")
	log.Println("GetTiny TinyKey:", TinyKey)
	if data, err := s.QueryOriLinkByTinyKey(TinyKey); err != nil {
		// http_err.NewError(c, http.StatusNotFound, errors.New("data not found"))
		c.HTML(http.StatusNotFound, "index.tmpl", gin.H{
			"title": "404 Link Not Found",
		})
		log.Println("GetTiny error:", err)
	} else {
		c.Redirect(http.StatusFound, data)
	}
}

func GetTiny(c *gin.Context) {
	s := persistence.GetTinyRepository()
	TinyKey := c.Query("tiny_key")
	log.Println("GetTiny TinyKey:", TinyKey)
	if data, err := s.QueryOriLinkByTinyKey(TinyKey); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("data not found"))
		log.Println("GetTiny error:", err)
	} else {
		c.JSON(http.StatusOK, gin.H{"ori_link": data})
	}
}

func CreateTiny(c *gin.Context) {
	type addParams struct {
		models.Tiny
		BaseUrl string `json:"base_url" form:"base_url"`
	}
	var tiny addParams
	var TinyLink string
	var baseUrl string
	var err error
	s := persistence.GetTinyRepository()
	_ = c.BindJSON(&tiny)
	baseUrl = tiny.BaseUrl
	if TinyLink, err = s.SaveOriLink(tiny.OriLink, baseUrl); err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		log.Println("GetTiny Error:", err)
	} else {
		c.JSON(http.StatusCreated, gin.H{"tiny_link": TinyLink})
	}
}
