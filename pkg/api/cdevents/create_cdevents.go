package cdevents

import (
	"fmt"

	sdk "github.com/cdevents/sdk-go/pkg/api"
)

// type CDEvent struct {
// 	Context sdk.Context
// 	Subject sdk.Subject
// }

func CreateRepositoryCreatedCDEvent(event sdk.CDEventReader) (*sdk.RepositoryCreatedEvent, error) {
	cdevent, err := sdk.NewRepositoryCreatedEvent()
	cdevent.SetSource(event.GetSource())
	cdevent.SetSubjectName(event.GetSubjectSource())
	cdevent.SetSubjectId(event.GetSubjectId())
	cdevent.SetSubjectUrl(event.GetSource())
	if err != nil {
		fmt.Println("Error occurred while Unmarshal gerritEvent data into ProjectCreated struct", err)
		return nil, err
	}
	return cdevent, nil
}

func CreateBranchCreatedCDEvent(event sdk.CDEventReader) (*sdk.BranchCreatedEvent, error) {
	cdevent, err := sdk.NewBranchCreatedEvent()
	cdevent.SetSource(event.GetSource())
	cdevent.SetSubjectId(event.GetSubjectId())
	cdevent.SetSubjectRepository(&sdk.Reference{Id: event.GetSubjectSource()})
	if err != nil {
		fmt.Println("Error occurred while Unmarshal gerritEvent data into ProjectCreated struct", err)
		return nil, err
	}
	return cdevent, nil
}
