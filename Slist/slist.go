package Slink

import (
	"fmt"
)

type Node struct {
	data any
	next *Node
}
type Slink struct {
	head *Node
	tail *Node
}

func (l *Slink) Append(n any) {
	newNode := &Node{data: n}
	if l.head == nil {
		l.head = newNode
	} else {
		l.tail.next = newNode
	}
	l.tail = newNode
}

func (l *Slink) Push(n any, pos int) {
	if pos == 0 {
		newNode := &Node{data: n}
		newNode.next = l.head
		l.head = newNode
	} else if pos > 0 && pos < l.Len() {
		newNode := &Node{data: n}
		p := l.head
		i := 0
		for i < pos-1 {
			p = p.next
			i += 1
		}
		newNode.next = p.next
		p.next = newNode
	} else {
		newNode := &Node{data: n}
		l.tail.next = newNode
		l.tail = newNode
	}

}
func (l *Slink) Len() int {
	p := l.head
	count := 0
	for p != nil {
		p = p.next
		count++
	}
	return count

}

func (l *Slink) Pop(pos int) any {
	//	b := 0
	var b any
	if pos == 0 {
		temp := l.head
		l.head = l.head.next
		b = temp.data
		temp.next = nil
	} else if pos > 0 && pos < l.Len()-1 {
		p := l.head
		i := 0
		for i < pos-1 {
			p = p.next
			i += 1
		}
		temp := p.next
		p.next = p.next.next
		b = temp.data
		temp.next = nil
	} else {
		p := l.head
		for p.next != l.tail {
			p = p.next
		}
		b = p.next.data
		l.tail = p
		p.next = nil
	}
	return b
}

func (l *Slink) Reverse() {

	// var prev *ListNode
	// currNode := head
	// nextNode := head
	// for nextNode != nil {
	// 	nextNode = nextNode.Next
	// 	currNode.Next = prev
	// 	prev = currNode
	// 	currNode = nextNode
	// }
	// return prev
	var prev *Node
	currNode := l.head
	nextNode := l.head
	for nextNode != nil {
		nextNode = nextNode.next
		currNode.next = prev
		prev = currNode
		currNode = nextNode
	}

	p := prev
	for p != nil {
		fmt.Print(p.data, " ")
		p = p.next
	}

	//return prev

}

// func (l *Slink) sort() {
// 	var b []interface{}
// 	var a []interface{}
// 	p := l.tail
// 	for p != nil {
// 		a = append(a, p.data)
// 		p = p.next
// 	}
// 	// sort.Slice(a, func(i, j int) bool {
// 	// 	return a[i] < a[j]
// 	// })
// 	sort.Sort(b(a))

// 	newNode := &Node{}
// 	s := newNode
// 	for _, v := range a {
// 		s.next = &Node{data: v}
// 		s = s.next
// 	}
// 	q := newNode.next
// 	for q != nil {
// 		fmt.Print(q.data, " ")
// 		q = q.next
// 	}

// }
func (l *Slink) Display() {
	p := l.head
	for p != nil {
		fmt.Print(p.data, " ")
		p = p.next
	}

}

func Disp(n *Node) {
	for n != nil {
		fmt.Print(n.data, " ")
		n = n.next
	}

}

// func removeval(head *Node, val int) {
// 	fake := &Node{data: 0}
// 	fake.next = head
// 	curr := fake
// 	for curr.next != nil {
// 		if curr.next.data == val {
// 			curr.next = curr.next.next
// 		} else {
// 			curr = curr.next
// 		}
// 	}
// 	fmt.Println(fake.next.data) // Return the linked list...
// }

func main() {
	l := Slink{}
	l.Append(20)
	l.Append(10)
	l.Append(2)
	l.Append(2)
	l.Append(3)
	l.Append("Hello")
	l.Append(60)
	l.Push(70, 3)
	l.Pop(7)
	l.Display()

	// fmt.Println("")
	// // l.Reverse()
	// fmt.Println("")
	// b := l.Reverse()
	// disp(b)

	// disp(b)
	// fmt.Println("")
	// l.sort()
	// l.Reverse()

}
