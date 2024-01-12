package github

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EventGitHub struct {
	event string
}

func NewGitHubEvent(event string) (pEvent *EventGitHub) {
	pEvent = &EventGitHub{event}
	return
}

func HandleTranslateGitHubEvent(c *gin.Context) {
	gitHubEvent := NewGitHubEvent("")
	gitHubEvent.TranslateEvent()
	c.IndentedJSON(http.StatusOK, "GitHub Event translated.")
}

func (pEvent *EventGitHub) TranslateEvent() {
	fmt.Println(pEvent.event)
	//handle translate GitHub event to CDEvent
}
