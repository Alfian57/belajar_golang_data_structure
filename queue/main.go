package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		fmt.Println("Cannot dequeue because queue is empty")
		return -1
	}

	toDequeue := q.items[0]

	q.items = q.items[1:]

	return toDequeue
}

func main() {
	myQueue := Queue{}

	myQueue.Enqueue(100)
	myQueue.Enqueue(200)
	myQueue.Enqueue(300)

	fmt.Println(myQueue)

	myQueue.Dequeue()

	fmt.Println(myQueue)
}
