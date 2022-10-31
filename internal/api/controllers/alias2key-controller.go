package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/mazeyqian/go-gin-gee/internal/pkg/models/alias2data"
	"github.com/mazeyqian/go-gin-gee/internal/pkg/persistence"
	http_err "github.com/mazeyqian/go-gin-gee/pkg/http-err"
)

func GetDataByAlias(c *gin.Context) {
	s := persistence.GetAlias2dataRepository()
	alias := c.Query("alias")
	log.Println("GetDataByAlias alias", alias)
	if data, err := s.Get(alias); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("data not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}

func CreateAlias2data(c *gin.Context) {
	s := persistence.GetAlias2dataRepository()
	var alias2dataInput models.Alias2data
	_ = c.BindJSON(&alias2dataInput)
	if err := s.Add(&alias2dataInput); err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
	} else {
		c.JSON(http.StatusCreated, alias2dataInput)
	}
}
