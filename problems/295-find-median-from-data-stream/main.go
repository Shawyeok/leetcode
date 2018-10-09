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

	if this.MiddleNode.Val > num {
		// insert before node
		node := this.MiddleNode
		for node.Val > num && node.Previous != nil {
			node = node.Previous
		}
		next := node.Next
		newNode.Next = next
		newNode.Previous = node
		node.Next = newNode
		if next != nil {
			next.Previous = newNode
		}

		this.MiddleNode = this.MiddleNode.Previous
	} else {
		// insert after node
		node := this.MiddleNode
		for node.Val < num && node.Next != nil {
			node = node.Next
		}
		next := node.Next
		node.Next = newNode
		newNode.Previous = node
		newNode.Next = next
		node.Previous = newNode
		if next != nil {
			next.Previous = newNode
		}

		this.MiddleNode = this.MiddleNode.Next
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
