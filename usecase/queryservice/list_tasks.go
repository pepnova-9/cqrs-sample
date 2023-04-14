package queryservice

import "github.com/pepnova-9/cqrs-sample/usecase/queryservice/readmodel"

type ListTasks interface {
	Fetch(limit int, offset int) ([]readmodel.TaskListRow, error)
}
