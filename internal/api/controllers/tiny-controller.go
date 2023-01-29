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
	s := persistence.GetTinyRepository()
	var tiny *models.Tiny
	var tinyData *models.Tiny
	var err error
	// OriLink := c.Query("ori_tiny")
	_ = c.BindJSON(&tiny)
	if tinyData, err = s.SaveOriLink(tiny.OriLink); err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		log.Println("GetTiny error:", err)
	} else {
		c.JSON(http.StatusCreated, gin.H{"data": tinyData})
	}
}
