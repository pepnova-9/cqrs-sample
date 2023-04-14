package usecase

import "github.com/pepnova-9/cqrs-sample/adapter/dataaccessor"

type ListTasksUsecase struct {
	taskDataAccessor  dataaccessor.Task
	userDataAccessor  dataaccessor.User
	labelDataAccessor dataaccessor.Label
}

type TaskListRow struct {
	TaskID       string
	TaskName     string
	AssigneeName string
	LabelName    string
}

type ListTasksUsecaseOutput struct {
	TaskList []*TaskListRow
}

// ListTasks は複数テーブルの値が必要になるタスク一覧を取得するユースケース
func (u *ListTasksUsecase) ListTasks() (ListTasksUsecaseOutput, error) {
	tasks, err := u.taskDataAccessor.List()
	if err != nil {
		return ListTasksUsecaseOutput{}, err
	}

	var taskList []*TaskListRow
	// 本当はIN句で取得してN+1問題を回避するが今回簡単のためループを回してます
	for _, task := range tasks {
		var taskListRow TaskListRow

		taskListRow.TaskID = string(task.TaskID)
		taskListRow.TaskName = task.Name

		if task.AssigneeID != "" {
			assignee, err := u.userDataAccessor.Find(task.AssigneeID)
			if err != nil {
				return ListTasksUsecaseOutput{}, err
			}
			taskListRow.AssigneeName = assignee.Name
		}

		if task.LabelID != "" {
			label, err := u.labelDataAccessor.Find(task.LabelID)
			if err != nil {
				return ListTasksUsecaseOutput{}, err
			}
			taskListRow.LabelName = label.Name
		}
		taskList = append(taskList, &taskListRow)
	}

	return ListTasksUsecaseOutput{TaskList: taskList}, nil
}

// ListTasksWithConditionAndPaging は複数テーブルをJoinして条件で絞り込んでページングするユースケース
func (u *ListTasksUsecase) ListTasksWithConditionAndPaging() (ListTasksUsecaseOutput, error) {
	// ???
	// 複数テーブルで条件あってページングする場合JOINするしかない
	// 各DataAccessorにFindしてユースケースでループを回すやり方では実現できない
	// ドメインモデルを返り値としたメソッドをDataAccessorに作ることはできない
	return ListTasksUsecaseOutput{}, nil
}
