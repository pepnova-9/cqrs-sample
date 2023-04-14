package dataaccessor

import "github.com/pepnova-9/cqrs-sample/domain"

type Task struct {
}

// Create はタスクを作成する
func (t *Task) Create(entity Task) error {
	_ = `
	INSERT xxx INTO tasks;
`
	return nil
}

// List はタスク一覧を取得する
func (t *Task) List() ([]*domain.Task, error) {
	_ = `
	SELECT * FROM tasks;
`
	// DataModelからEntityに変換する

	return []*domain.Task{}, nil
}
