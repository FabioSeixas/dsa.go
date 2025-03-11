package graph

import "fmt"

type GraphNode struct {
	Id    string
	Nodes []string
}

type Graph struct {
	Nodes []GraphNode
}

func (g *Graph) ValidateCircular() {

	allNodes := make(map[string]*GraphNode)
	for _, node := range g.Nodes {
		allNodes[node.Id] = &node
	}

	stack := make(map[string]bool)
	visited := make(map[string]bool)

	var dfs func(node *GraphNode) error
	dfs = func(node *GraphNode) error {
		fmt.Println("node: ", node)
		fmt.Println("stack: ", stack)
		if stack[node.Id] {
			panic("Circular dependency")
		}

		if visited[node.Id] {
			return nil
		}

		stack[node.Id] = true

		for _, innerNodeId := range node.Nodes {
			if exist := allNodes[innerNodeId]; exist != nil {
				dfs(exist)
			}

		}

		delete(stack, node.Id)

		return nil
	}

	for _, node := range allNodes {
		dfs(node)
	}
}
