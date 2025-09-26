package job

import "main/worker/domain/task"

func NewJob(task task.ITask) IJob {
	return &Job{
		Task: task,
	}
}
