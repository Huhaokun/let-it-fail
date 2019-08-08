package admin

type NodeRegistry struct {
	nodes map[string]Node
}

func (n *NodeRegistry) Get(id string) Node {
	if node, ok := n.nodes[id]; ok {
		return node
	} else {
		return nil
	}
}

func (n *NodeRegistry) List() []Node {
	var nodes []Node
	for _, node := range n.nodes {
		nodes = append(nodes, node)
	}

	return nodes
}
