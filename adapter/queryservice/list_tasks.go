package queryservice

import (
	"fmt"
	"github.com/pepnova-9/cqrs-sample/usecase/queryservice/readmodel"
)

type ListTasks struct {
}

type TaskListRowDataModel struct {
	TaskID       string
	TaskName     string
	AssigneeName string
	LabelName    string
}

func (l *ListTasks) Fetch(limit int, offset int) ([]readmodel.TaskListRow, error) {
	_ = fmt.Sprintf(`
	SELECT
		tasks.id,
		tasks.name,
		users.name,
		labels.name
	FROM
		tasks
	LEFT JOIN
		users
	ON
		tasks.assignee_id = users.id
		users.name IS NOT '' -- 絞り込み条件(1) usersテーブルへの絞り込み
	LEFT JOIN
		labels
	ON
		tasks.label_id = labels.id
	WHERE
		tasks.name IS NOT '' -- 絞り込み条件(2) tasksテーブルへの絞り込み
	ORDER BY
		users.name DESC -- ソート条件
	LIMIT %d OFFSET %d  -- ページング条件
`, limit, offset)

	// SQLの結果がマッピングされたstructの配列
	var queryResult []TaskListRowDataModel

	var taskList []readmodel.TaskListRow
	for _, dataModel := range queryResult {
		taskList = append(taskList, readmodel.TaskListRow{
			TaskID:       dataModel.TaskID,
			TaskName:     dataModel.TaskName,
			AssigneeName: dataModel.AssigneeName,
			LabelName:    dataModel.LabelName,
		})
	}

	return taskList, nil
}
