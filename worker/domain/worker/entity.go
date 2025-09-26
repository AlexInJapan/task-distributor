package worker

import "main/worker/domain/task"

type WorkerEntity struct {
	ID        int
	WorkCh    chan task.TaskEntity
	Available bool
}

func NewWorker(id int) *WorkerEntity {
	return &WorkerEntity{
		ID:        id,
		WorkCh:    make(chan task.TaskEntity, 1),
		Available: true,
	}
}

func (w *WorkerEntity) Work() {

}
