package Controllers

import (
	"context"
	"net/http"
	"time"
	"vtrial/pkg/Models"
	"vtrial/pkg/services"

	"github.com/gin-gonic/gin"
)

//This is where the actual magic happens
//Our controllers should use the I_Controller implementation


type Controller struct{
	SearchService services.ISearchService
}


func NewController(searchService services.ISearchService) *Controller {
	return &Controller{SearchService: searchService}
}

func (ct *Controller) Routes(router *gin.RouterGroup) {
	router.GET("/", ct.healthCheck)
	router.POST("UseSavePage", ct.UseSavePage)
	router.GET("UseComputeResult", ct.UseComputeResult)
}

func (ct *Controller) healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Alive")
}
func (ct *Controller) UseSavePage(c *gin.Context) {
	var newPage models.Page
	if err := c.BindJSON(&newPage); err != nil {
	c.IndentedJSON(http.StatusBadRequest,err)
	return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	insertErr := ct.SearchService.SavePage(ctx, newPage)
	if insertErr != nil {
	c.IndentedJSON(http.StatusInternalServerError, insertErr)
	return
}
	c.IndentedJSON(http.StatusCreated, newPage)
}
func (ct *Controller) UseComputeResult(c *gin.Context) {
	var words models.Keywords
	if err := c.BindJSON(&words); err != nil {
		c.IndentedJSON(http.StatusBadRequest,err)
	return
	}
	ans, err := ct.SearchService.ComputeResult(c, words)
	if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}
	c.IndentedJSON(http.StatusOK, ans)
}

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}
