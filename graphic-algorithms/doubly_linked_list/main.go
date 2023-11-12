package doubly_linked_list

import (
	"fmt"
)

type Node struct {
	data string
	prev *Node
	next *Node
}

var head *Node = new(Node)
var tail *Node = new(Node)

func initial() {
	head.data = "San Francisco"
	head.next = nil
	head.prev = nil

	var nodeOakland *Node = &Node{data: "Oakland", prev: head, next: nil}
	head.next = nodeOakland

	var nodeBerkeley *Node = &Node{data: "Berkeley", prev: nodeOakland, next: nil}
	nodeOakland.next = nodeBerkeley

	tail.data = "Fermont"
	tail.prev = nodeBerkeley
	tail.next = nil
	nodeBerkeley.next = tail
}

func output(node *Node) {
	var p = node
	var end *Node = nil

	for {
		if p == nil {
			break
		}

		fmt.Printf("%s", p.data)
		p = p.next
	}
	fmt.Printf("End\n\n")

	p = end
	for {
		if p == nil {
			break
		}

		fmt.Printf("%s->", p.data)
		p = p.prev
	}
	fmt.Printf("Start\n\n")
}

func append(data string) {
	var newNode *Node = new(Node)
	newNode.data = data
	newNode.next = nil
	tail.next = newNode
	newNode.prev = tail
	tail = newNode
}

func insert(insertPosition int, data string) {
	var p = head
	var i = 0
	for {
		if p.next == nil || i >= insertPosition-1 {
			break
		}
		p = p.next
		i++
	}

	var newNode *Node = new(Node)
	newNode.data = data
	newNode.next = p.next // newNode next point to next Node
	p.next = newNode      // current next point to newNode
	newNode.prev = p
	newNode.next.prev = newNode
}

func removeNode(removePosition int) {
	var p = head
	var i = 0
	for {
		if p.next == nil || i >= removePosition-1 {
			break
		}

		p = p.next
		i++
	}
	var temp = p.next
	p.next = p.next.next
	p.next.prev = p
	temp.next = nil
	temp.prev = nil
}

func main() {
	initial()
	output(head)
}
