package main

func main() {
}

type MedianFinder struct {
	Len        int
	MiddleNode *Node
}

type Node struct {
	Val      int
	Next     *Node
	Previous *Node
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	var mf MedianFinder
	return mf
}

func (this *MedianFinder) AddNum(num int) {
	this.Len++

	newNode := &Node{}
	newNode.Val = num
	if this.MiddleNode == nil {
		this.MiddleNode = newNode
		return
	}

	node := this.MiddleNode
	if num <= this.MiddleNode.Val {
		// insert left side of MiddleNode
		for num < node.Val && node.Previous != nil {
			node = node.Previous
		}
		if num > node.Val {
			this.insertRight(node, newNode)
		} else {
			this.insertLeft(node, newNode)
		}

		if this.Len%2 == 0 {
			this.MiddleNode = this.MiddleNode.Previous
		}
	} else {
		// insert right side of MiddleNode
		for num > node.Val && node.Next != nil {
			node = node.Next
		}
		if num > node.Val {
			this.insertRight(node, newNode)
		} else {
			this.insertLeft(node, newNode)
		}

		if this.Len%2 != 0 {
			this.MiddleNode = this.MiddleNode.Next
		}
	}
}

func (this *MedianFinder) insertLeft(node *Node, newNode *Node) {
	previous := node.Previous
	newNode.Next = node
	newNode.Previous = previous
	node.Previous = newNode
	if previous != nil {
		previous.Next = newNode
	}
}

func (this *MedianFinder) insertRight(node *Node, newNode *Node) {
	next := node.Next
	newNode.Next = next
	newNode.Previous = node
	node.Next = newNode
	if next != nil {
		next.Previous = newNode
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.Len%2 == 0 {
		return (float64(this.MiddleNode.Val) + float64(this.MiddleNode.Next.Val)) / 2
	}

	return float64(this.MiddleNode.Val)
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
