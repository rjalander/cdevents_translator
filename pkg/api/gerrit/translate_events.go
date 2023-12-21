package gerrit

import (
	"encoding/json"
	"fmt"
	cdevents "github.com/cdevents/sdk-go/pkg/api"
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
	fmt.Println("Translated ProjectCreated EventGerrit received to CDEvent ==>: ", cdevent)
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
}

func (projectCreated *ProjectCreated) TranslateProjectCreatedEvent() (*cdevents.RepositoryCreatedEvent, error) {
	cdevent, err := cdevents.NewRepositoryCreatedEvent()
	cdevent.SetSource(projectCreated.RepoURL)
	cdevent.SetSubjectName(projectCreated.ProjectName)
	cdevent.SetSubjectId(projectCreated.HeadName)
	cdevent.SetSubjectUrl(projectCreated.RepoURL)
	if err != nil {
		fmt.Println("Error occurred while Unmarshal gerritEvent data into ProjectCreated struct", err)
		return nil, err
	}
	return cdevent, nil
}

func (refUpdated *RefUpdated) TranslateRefUpdatedEvent() (*cdevents.BranchCreatedEvent, error) {
	cdevent, err := cdevents.NewBranchCreatedEvent()
	cdevent.SetSource(refUpdated.RepoURL)
	cdevent.SetSubjectId(refUpdated.RefUpdate.RefName)
	cdevent.SetSubjectRepository(&cdevents.Reference{Id: refUpdated.RefUpdate.RefName})
	if err != nil {
		fmt.Println("Error occurred while Unmarshal gerritEvent data into ProjectCreated struct", err)
		return nil, err
	}
	return cdevent, nil
}
