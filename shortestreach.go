package main

import (
	"errors"
	"fmt"
)

var nodesById map[int]*Node = map[int]*Node{}

type NodeLevel struct {
	node  *Node
	level int
}

type Queue interface {
	Enqueue(NodeLevel)
	Dequeue() (NodeLevel, error)
	Peek() (NodeLevel, error)
	IsEmpty() bool
}

type nodeLevelQueue struct {
	data []NodeLevel
	size int
}

func NewNodeLevelQueue() Queue {
	return &nodeLevelQueue{data: []NodeLevel{}}
}

func (n *nodeLevelQueue) Enqueue(value NodeLevel) {
	n.data = append(n.data, value)
	n.size += 1
}

func (n *nodeLevelQueue) Dequeue() (NodeLevel, error) {
	if n.size > 0 {
		value := n.data[0]
		n.size -= 1
		n.data = n.data[1:]

		return value, nil
	}

	return NodeLevel{}, errors.New("No Such Element")
}

func (n *nodeLevelQueue) Peek() (NodeLevel, error) {
	if n.size > 0 {
		value := n.data[0]
		return value, nil
	}

	return NodeLevel{}, errors.New("No Such Element")
}

func (n *nodeLevelQueue) IsEmpty() bool {
	return n.size == 0
}

func (n *nodeLevelQueue) String() string {
	return fmt.Sprintf("%v", n.data)
}

type Node struct {
	id        int
	adjacents []*Node
	added     map[int]struct{}
}

func NewNode(id int) *Node {
	return &Node{id: id,
		adjacents: []*Node{},
		added:     map[int]struct{}{},
	}
}

func (n *Node) AddAdjacent(node *Node) {
	if _, present := n.added[node.id]; !present {
		n.added[node.id] = struct{}{}
		n.adjacents = append(n.adjacents, node)
		node.AddAdjacent(n)
	}
}

func (n *Node) String() string {
	return fmt.Sprintf("{Node id: %v}", n.id)
}

func main() {
	var numTestCases int

	fmt.Scanf("%v\n", &numTestCases)

	var numNodes int
	var numEdges int

	for i := 0; i < numTestCases; i++ {
		fmt.Scanf("%v %v\n", &numNodes, &numEdges)

		for n := 1; n <= numNodes; n++ {
			node := NewNode(n)
			nodesById[n] = node
		}

		var fromNode int
		var toNode int

		for e := 0; e < numEdges; e++ {
			fmt.Scanf("%v %v\n", &fromNode, &toNode)

			from := nodesById[fromNode]
			to := nodesById[toNode]

			from.AddAdjacent(to)
		}

		var startingNode int

		fmt.Scanf("%v\n", &startingNode)

		for node := 1; node <= numNodes; node++ {
			shortestDistance(startingNode, node)
		}

		fmt.Println()
	}
}

func shortestDistance(from, to int) {
	if from != to {
		processed := map[int]struct{}{}
		queue := NewNodeLevelQueue()

		start := nodesById[from]

		queue.Enqueue(NodeLevel{node: start})

		var target *Node

	BeginSearch:
		for target == nil && !queue.IsEmpty() {
			current, _ := queue.Dequeue()

			node := current.node
			level := current.level

			if _, visited := processed[node.id]; !visited {
				processed[node.id] = struct{}{}

				for _, adjacent := range node.adjacents {
					if adjacent.id == to {
						target = adjacent
						fmt.Print(6*(level+1), " ")
						break BeginSearch
					}

					if _, seen := processed[adjacent.id]; !seen {
						queue.Enqueue(NodeLevel{node: adjacent, level: level + 1})
					}
				}
			}
		}

		found := target != nil

		if !found {
			fmt.Print(-1, " ")
		}
	}
}
