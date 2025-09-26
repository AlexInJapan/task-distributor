package task

import "github.com/google/uuid"

type TaskEntity struct {
	id   uuid.UUID
	name string
	f    func(s string)
}

func New(name string, f func(string)) *TaskEntity {
	id := uuid.New()
	return &TaskEntity{
		id:   id,
		name: name,
		f:    f,
	}
}

func (t *TaskEntity) GetTaskName() string {
	return t.name
}

func (t *TaskEntity) GetFunc() func(string) {
	return t.f
}

func (t *TaskEntity) GetID() uuid.UUID {
	return t.id
}
