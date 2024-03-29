package jin

import "strings"

type node struct {
	pattern  string
	part     string
	children []*node
	isWild   bool
}

func (n *node) MatchChild(part string) *node {
	// Match node of certain parts
	for _, child := range n.children {
		if child.part == part || child.isWild == true {
			return child
		}
	}
	return nil
}

func (n *node) MatchAllChildren(part string) []*node {
	// Match all node with certain part
	ret := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild == true {
			ret = append(ret, child)
		}
	}
	return ret
}

func (n *node) Insert(pattern string, parts []string, height int) {
	// Insert a new route to route tree
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.MatchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.Insert(pattern, parts, height+1)
}

func (n *node) Search(parts []string, height int) *node {
	// Search routes
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.MatchAllChildren(part)
	for _, child := range children {
		result := child.Search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
