package job

import "main/worker/domain/task"

type IJob interface {
	Run(s string)
}
type Job struct {
	Task task.ITask
}

func (j *Job) Run(s string) {
	j.Task.Run(s)
}
