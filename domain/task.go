package domain

type Task struct {
	TaskID     TaskID
	Name       string
	AssigneeID UserID
	LabelID    LabelID
}

type TaskID string
