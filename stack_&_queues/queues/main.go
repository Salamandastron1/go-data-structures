package main

import "fmt"

// Queue Struct represents a ordered queue that holds a slice
type Queue struct {
	items []int
}

// Enqueue put a new item at the end of the queue
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Dequeue removes an item from the start of the queue
// and RETURNS the removed value
func (q *Queue) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]

	return toRemove
}

func main() {
	myQueue := Queue{}
	fmt.Println(myQueue)

	myQueue.Enqueue(1)
	myQueue.Enqueue(2)
	myQueue.Enqueue(3)

	fmt.Println(myQueue)

	fmt.Printf("item removed: %d\n", myQueue.Dequeue())
	fmt.Println(myQueue)
}
