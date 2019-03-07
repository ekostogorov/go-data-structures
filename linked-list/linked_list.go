package ll

// LinkedList represents linked list structure
type LinkedList struct {
	length int
	first  *Node
	last   *Node
}

// Node represents linked list node
type Node struct {
	Data     string
	Previous *Node
	Next     *Node
}

// New constructs linked list
func New() *LinkedList {
	return &LinkedList{}
}

// Append adds element to linked list
func (l *LinkedList) Append(data string) {
	node := Node{
		Data: data,
	}

	if l.length == 0 {
		l.first, l.last = &node, &node
	} else {
		current := l.last
		node.Previous = current
		current.Next = &node

		l.last = &node
	}

	l.length++
}

// Remove deletes element from linked list
func (l *LinkedList) Remove(data string) {
	if l.length == 0 {
		panic("List is empty")
	}

	current := l.first
	for {
		if current.Data != data {
			current = current.Next
		} else {
			previous := current.Previous

			if current.Next != nil {
				next := current.Next
				previous.Next = next
				next.Previous = previous
			}

			previous.Next = nil
			current.Next = nil
			current.Previous = nil

			l.length--

			break
		}
	}
}

// ElementByIndex returns element by index
func (l *LinkedList) ElementByIndex(index int) string {
	var counter int

	current := l.first
	for {
		if counter < index {
			current = current.Next

			counter++
		} else {
			break
		}
	}

	return current.Data
}

// RemoveByIndex deletes element by index
func (l *LinkedList) RemoveByIndex(index int) {
	element := l.ElementByIndex(index)
	l.Remove(element)
}

// List returns all elements in linked list as slice
func (l *LinkedList) List() (list []string) {
	current := l.first

	for {
		list = append(list, current.Data)

		if current.Next != nil {
			current = current.Next
		} else {
			break
		}
	}

	return
}

// Length returns length of linked list
func (l *LinkedList) Length() int {
	return l.length
}
