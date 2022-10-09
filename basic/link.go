package main

import "fmt"

type LinkNode struct {
	Data     int64
	NextNode *LinkNode
}

func main() {
	node := new(LinkNode)
	node.Data = 1

	node1 := new(LinkNode)
	node1.Data = 2

	node2 := new(LinkNode)
	node2.Data = 3

	node.NextNode = node1
	node1.NextNode = node2

	nownode := node

	for {
		if nownode != nil {
			fmt.Println(nownode.Data)

			nownode = nownode.NextNode

			continue
		}
		break
	}
}
