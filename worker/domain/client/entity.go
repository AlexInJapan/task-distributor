package client

import (
	"fmt"
	"main/worker/domain/job"
	"main/worker/domain/queue"
	"main/worker/domain/task"
	"main/worker/domain/worker"
	"sync"
)

type Client struct {
	Queue   *queue.QueueEntity
	TaskCh  chan *task.TaskEntity
	Workers map[int]*worker.WorkerEntity
	Wg      *sync.WaitGroup
	Done    bool
}

func New() *Client {
	var wg sync.WaitGroup

	cl := &Client{
		Queue:   queue.NewQueue(),
		TaskCh:  make(chan *task.TaskEntity, 5),
		Workers: make(map[int]*worker.WorkerEntity, 2),
		Wg:      &wg,
		Done:    false,
	}
	for i := range 2 {
		w := worker.NewWorker(i)
		cl.Workers[i] = w
	}
	if !cl.Done {

		go cl.Listen()
	}

	return cl
}

func (c *Client) Run(t *task.TaskEntity) {
	taskInterface := task.NewTask(t)
	jobInterface := job.NewJob(taskInterface)

	jobInterface.Run(t.GetTaskName())
	delete(c.Queue.List, t.GetID())
	c.Wg.Done()
}

func (c *Client) Enqueue(t *task.TaskEntity) {

	if c.Queue.IsFull() {

	}

	c.Wg.Add(1)

	c.Queue.MU.Lock()
	c.Queue.List[t.GetID()] = t
	fmt.Println("Enqueued task")
	c.Queue.MU.Unlock()

	c.TaskCh <- t

}

func (c *Client) Close() {
	close(c.TaskCh)
	c.Wg.Done()
}

func (c *Client) List() {
	fmt.Println(c.Queue)
}

func (c *Client) Listen() {
	for t := range c.TaskCh {
		fmt.Println("got task >> ", t.GetTaskName())
	}

}
