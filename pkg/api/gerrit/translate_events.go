package gerrit

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cdevents/sdk-go/pkg/api"
	cdevents "github.com/cdevents/translator/pkg/api/cdevents"
)

func (pEvent *EventGerrit) HandleRefUpdatedEvent() {
	var refUpdated RefUpdated
	err := json.Unmarshal([]byte(pEvent.Event), &refUpdated)
	if err != nil {
		fmt.Println("Error occurred while Unmarshal gerritEvent data into RefUpdated struct", err)
		return
	}
	refUpdated.RepoURL = pEvent.RepoURL
	fmt.Println("RefUpdated EventGerrit received : ", refUpdated.RefUpdate.RefName, refUpdated.Submitter.Name, refUpdated.CommonFields.Type)
	cdevent, _ := refUpdated.TranslateRefUpdatedEvent()
	fmt.Println("Translated RefUpdated EventGerrit received to CDEvent ==>: ", cdevent)

}

func (pEvent *EventGerrit) HandleProjectCreatedEvent() {
	var projectCreated ProjectCreated
	err := json.Unmarshal([]byte(pEvent.Event), &projectCreated)
	if err != nil {
		fmt.Println("Error occurred while Unmarshal gerritEvent data into ProjectCreated struct", err)
		return
	}
	projectCreated.RepoURL = pEvent.RepoURL
	fmt.Println("ProjectCreated EventGerrit received : ", projectCreated.ProjectName, projectCreated.HeadName, projectCreated.CommonFields.Type)
	cdevent, _ := projectCreated.TranslateProjectCreatedEvent()
	fmt.Println("Translated ProjectCreated EventGerrit received to CDEvent: ", cdevent)
	cdevents.SendCDEvent(cdevent)
}

/*
*
Create RepositoryCreated CDEvent using a ProjectCreated struct
*
*/
func (projectCreated *ProjectCreated) TranslateProjectCreatedEvent() (*sdk.RepositoryCreatedEvent, error) {
	cdevent, err := sdk.NewRepositoryCreatedEvent()
	cdevent.SetSource(projectCreated.RepoURL)
	cdevent.SetSubjectName(projectCreated.ProjectName)
	cdevent.SetSubjectId(projectCreated.HeadName)
	cdevent.SetSubjectUrl(projectCreated.RepoURL)
	if err != nil {
		fmt.Println("Error occurred while creating CDEvent - NewRepositoryCreatedEvent ", err)
		return nil, err
	}
	return cdevent, nil
}

/*
*
Create RepositoryCreated CDEvent using a common method by passing sdk.CDEvent
*
*/
func (projectCreated *ProjectCreated) TranslateProjectCreatedEvent1() (*sdk.RepositoryCreatedEvent, error) {
	var event sdk.CDEvent

	event.SetSource(projectCreated.RepoURL)
	event.SetSubjectId(projectCreated.HeadName)
	cdevent, err := cdevents.CreateRepositoryCreatedCDEvent(event)
	if err != nil {
		fmt.Println("Error occurred while CreateRepositoryCreatedCDEvent ", err)
		return nil, err
	}
	return cdevent, nil

}

func (refUpdated *RefUpdated) TranslateRefUpdatedEvent() (*sdk.BranchCreatedEvent, error) {
	cdevent, err := sdk.NewBranchCreatedEvent()
	cdevent.SetSource(refUpdated.RepoURL)
	cdevent.SetSubjectId(refUpdated.RefUpdate.RefName)
	cdevent.SetSubjectRepository(&sdk.Reference{Id: refUpdated.RefUpdate.RefName})
	if err != nil {
		fmt.Println("Error occurred while creating CDEvent - NewBranchCreatedEvent", err)
		return nil, err
	}
	return cdevent, nil
}
