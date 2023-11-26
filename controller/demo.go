package controller

import (
	"gorm-hex/controller/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewDemoController(router *gin.Engine) {
	ping := router.Group("/ping")
	{
		ping.GET("/", demoPing)
		ping.GET("/:name", demoPingName)
		ping.POST("/postjson", postJson)
	}
}

func demoPing(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func demoPingName(ctx *gin.Context) {
	// name := ctx.Param("name")
	ctx.JSON(http.StatusOK, gin.H{
		"Hello": "hello",
	})
}

func postJson(ctx *gin.Context) {
	person := model.Person{}
	ctx.ShouldBindJSON(&person)

	ctx.JSON(http.StatusOK, gin.H{
		"message" : "pong " + person.Name + " " + strconv.Itoa(person.Age),
	})
}