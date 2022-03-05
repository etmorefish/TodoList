package serializer

import "todo-list/model"

type Task struct {
	ID        uint   `json:"id" example:"1"`
	Title     string `json:"title" example:"title"`
	Content   string `json:"content" example:"content"`
	View      uint64 `json:"view" example:"12"`
	Status    int    `json:"status" example:"0"`
	CreateAt  int64  `json:"create_at"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
}

func BuildTask(item model.Task) Task {
	return Task{
		ID:      item.ID,
		Title:   item.Title,
		Content: item.Content,
		// View:    item.View(),
		Status: item.Status,
		// CreateAt:  item.CreateAt.Unix(),
		StartTime: item.StartTime,
		EndTime:   item.EndTime,
	}
}

func BuildTasks(items []model.Task) (tasks []Task) {
	for _, item := range items {
		task := BuildTask(item)
		tasks = append(tasks, task)
	}
	return tasks
}
