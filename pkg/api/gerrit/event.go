package gerrit

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EventGerrit struct {
	Event   string
	RepoURL string
}

func NewGerritEvent(event string, repoURL string) (pEvent *EventGerrit) {
	pEvent = &EventGerrit{event, repoURL}
	return
}

func HandleTranslateGerritEvent(c *gin.Context) {
	repoURL := c.Request.Header.Get("X-Origin-Url")
	data, err := c.GetRawData()
	if err != nil {
		fmt.Println("Error occurred while getting request data from gerrit webhook", err)
		return
	}
	gerritEvent := NewGerritEvent(string(data), repoURL)
	gerritEvent.TranslateEvent()
	c.IndentedJSON(http.StatusOK, "Gerrit Event translated.")
}

func (pEvent *EventGerrit) TranslateEvent() {
	eventMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(pEvent.Event), &eventMap)
	if err != nil {
		fmt.Println("Error occurred while Unmarshal gerritEvent data into gerritEvent map", err)
		return
	}
	fmt.Println("Received Gerrit EventGerrit: ", pEvent.Event)
	fmt.Println("handling translating to CDEvent from Gerrit EventGerrit type: ", eventMap["type"])

	switch eventMap["type"] {
	case "project-created":
		pEvent.HandleProjectCreatedEvent()
	case "ref-updated":
		pEvent.HandleRefUpdatedEvent()
	default:
		fmt.Println("Not handling CDEvent translation for Gerrit event type: ", eventMap["type"])
	}
}
