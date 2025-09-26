package queue

import (
	"main/worker/domain/task"
	"sync"

	"github.com/google/uuid"
)

type QueueEntity struct {
	List map[uuid.UUID]*task.TaskEntity
	MU   sync.RWMutex
}

func NewQueue() *QueueEntity {
	return &QueueEntity{
		List: make(map[uuid.UUID]*task.TaskEntity, 5),
	}
}

func (Q *QueueEntity) IsFull() bool {
	if len(Q.List) == 5 {
		return true
	}
	return false
}
