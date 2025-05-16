package main

import "fmt"

type Node struct {
	Value string
	Next  *Node
}

func (n *Node) Add(item string) {
	last := n
	for last.Next != nil {
		last = last.Next
	}

	last.Next = &Node{item, last.Next}
}

func main() {
	head := &Node{
		Value: "url/playlist/1",
		Next:  nil,
	}

	head.Add("url/playlist/2")
	head.Add("url/playlist/3")

	fmt.Print(head)
	fmt.Print(head.Next)
	fmt.Print(head.Next.Next)
}
