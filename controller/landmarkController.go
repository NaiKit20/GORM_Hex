package controller

import (
	"gorm-hex/model"
	"gorm-hex/service"
	"net/http"
	"github.com/gin-gonic/gin"
)

func NewLandmarkController(router *gin.Engine) {
	ping := router.Group("/landmark")
	{
		ping.GET("", getAllLandmark)
		ping.GET("/:name", getLandmarkByName)
	}
}

func getAllLandmark(ctx *gin.Context) {
	searchService := service.NewSearchService()
	landmarks := []model.Landmark{}
	landmarks, err := searchService.GetAllLandmarks()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error" : "error นะจ้ะ",
		})
	}

	ctx.JSON(http.StatusOK, landmarks)
}

func getLandmarkByName(ctx *gin.Context) {
	name := ctx.Param("name")
	searchService := service.NewSearchService()
	landmarks := []model.Landmark{}
	landmarks, err := searchService.FindByName(name)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error" : "error นะจ้ะ",
		})
	}

	ctx.JSON(http.StatusOK, landmarks)
}