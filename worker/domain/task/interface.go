package task

type ITask interface {
	Run(s string)
}

type Task struct {
	TE *TaskEntity
}

func NewTask(TE *TaskEntity) ITask {
	return &Task{
		TE: TE,
	}
}

func (t *Task) Run(s string) {
	t.TE.f(s)
}
