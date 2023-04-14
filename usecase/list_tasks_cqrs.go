package usecase

import (
	"github.com/pepnova-9/cqrs-sample/usecase/queryservice"
	"github.com/pepnova-9/cqrs-sample/usecase/queryservice/readmodel"
)

type ListTaskCQRSPatternUsecase struct {
	queryService queryservice.ListTasks
}

type ListTaskCQRSPatternUsecaseInput struct {
	Page int
}

func (l *ListTaskCQRSPatternUsecase) Run(input ListTaskCQRSPatternUsecaseInput) ([]readmodel.TaskListRow, error) {
	tasksPerPage := 10
	offset := 10 * input.Page

	list, err := l.queryService.Fetch(tasksPerPage, offset)
	if err != nil {
		return []readmodel.TaskListRow{}, err
	}
	return list, nil
}
