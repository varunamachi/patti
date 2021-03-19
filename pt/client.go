package pt

type Client interface {
	CreateList(list *TaskList) error
	UpdateList(list *TaskList) error
	RemoveList(listId string) error
	CreateTask(listId string, task *Task) error
	UpdateTask(listId string, task *Task) error
	RemoveTask(listId string, taskId string) error
	SetTaskStatus(taskId string, status Status) error
	SetListStatus(listId string, status Status) error
}

type LocalClient struct{}

func (lc LocalClient) CreateList(list *TaskList) error {
	return nil
}

func (lc LocalClient) UpdateList(list *TaskList) error {
	return nil
}

func (lc LocalClient) RemoveList(listId string) error {
	return nil
}

func (lc LocalClient) CreateTask(listId string, task *Task) error {
	return nil
}

func (lc LocalClient) UpdateTask(listId string, task *Task) error {
	return nil
}

func (lc LocalClient) RemoveTask(listId string, taskId string) error {
	return nil
}

func (lc LocalClient) SetTaskStatus(taskId string, status Status) error {
	return nil
}

func (lc LocalClient) SetListStatus(listId string, status Status) error {
	return nil
}
