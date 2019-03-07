package main

import (
	ll "data-structures/linked-list"
	"data-structures/queue"
	"data-structures/set"
	"data-structures/stack"
	"log"
)

func testLinkedList() {
	linkedList := ll.New()

	linkedList.Append("node1")
	linkedList.Append("node2")
	linkedList.Append("node3")
	linkedList.Append("node4")

	log.Printf("List length is: %+v", linkedList.Length())
	log.Printf("List is: %+v", linkedList.List())

	index2 := linkedList.ElementByIndex(2)
	log.Printf("Index 2: %+v", index2)

	linkedList.Remove("node4")
	log.Printf("List length is: %+v", linkedList.Length())
	log.Printf("List is: %+v", linkedList.List())

	linkedList.RemoveByIndex(2)
	log.Printf("List length is: %+v", linkedList.Length())
	log.Printf("List is: %+v", linkedList.List())
}

func testStack() {
	s := stack.New()

	s.Push("first")
	s.Push("second")
	s.Push("third")

	log.Printf("Stack size is: %d", s.Size())
	log.Printf("Popped element is: %s", s.Pop())
	log.Printf("Stack size is: %d", s.Size())
}

func testSet() {
	set1 := set.New()
	set1.Add("1")
	set1.Add("2")
	set1.Add("3")
	set1.Add("4")
	log.Printf("Set 1 is: %+v", set1.Values())

	set2 := set.New()
	set2.Add("1")
	set2.Add("5")
	set2.Add("7")
	set2.Add("8")
	log.Printf("Set 2 is: %+v", set2.Values())

	set3 := set.New()
	set3.Add("1")
	set3.Add("5")
	set3.Add("7")
	set3.Add("8")
	log.Printf("Set 3 is: %+v", set3.Values())

	log.Printf("Union set1 & set2 %+v", set1.Union(set2))
	log.Printf("Intersection set1 & set2 %+v", set1.Intersect(set2))
	log.Printf("Subset set1 & set2: %+v", set1.Subset(set2))
	log.Printf("Subset set2 & sete: %+v", set2.Subset(set3))
}

func testQueue() {
	q := queue.New()

	q.Enqueue("1")
	q.Enqueue("2")
	q.Enqueue("3")

	log.Printf("Enqueued 1: %s", q.Dequeue())
	log.Printf("Enqueued 1: %s", q.Dequeue())
	log.Printf("Enqueued 1: %s", q.Dequeue())

}

func main() {
	// testLinkedList()
	// time.Sleep(3 * time.Second)

	// testStack()
	// time.Sleep(3 * time.Second)

	// testQueue()
	// time.Sleep(3 * time.Second)

	testSet()
}
