package api

import (
	"github.com/cdevents/translator/pkg/api/gerrit"
	"github.com/cdevents/translator/pkg/api/github"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EventTranslator interface {
	TranslateEvent()
}

func initialiseRoutes(router *gin.Engine) {
	router.GET("/", handleGetTranslator)
	router.POST("/gerrit-webhooks", gerrit.HandleTranslateGerritEvent)
	router.POST("/github-webhooks", github.HandleTranslateGitHubEvent)
}

func handleGetTranslator(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello cdevents-translator!!")
}
