package main

import (
	"fmt"
	"main/worker/domain/client"
	"main/worker/domain/task"
)

func main() {
	// #1 thread will listen for enqueue
	cl := client.New()

	t1 := task.New("task name 1", MyTaskFunction1)
	cl.Enqueue(t1)

	t2 := task.New("task name 2", MyTaskFunction2)
	cl.Enqueue(t2)
	cl.List()

	t3 := task.New("task name 3", MyTaskFunction2)
	cl.Enqueue(t3)
	cl.List()

	cl.Run(t3)

	cl.Run(t1)

	cl.Run(t2)

	cl.Wg.Done()
	cl.Wg.Wait()
}

func MyTaskFunction1(s string) {
	s = s + " is done Processing"
	fmt.Println(s)
}

func MyTaskFunction2(s string) {
	s = s + " is done Processing"
	fmt.Println(s)
}
