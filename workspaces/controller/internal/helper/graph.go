package helper

// DetectGraphCycle checks if there is a cycle involving a given node in directed graph
//
// Assumptions:
//   - all nodes have AT MOST one OUTGOING edge
//
// Parameters:
//   - startNode: the node to start the cycle detection from
//   - checkedNodes: a map of nodes which have already been checked for cycles (updated in-place if no cycle is detected)
//   - edgeMap: a map representing the edges in the graph
//
// Returns:
//   - returns nil, if no cycle is detected
//   - returns a slice of nodes representing the cycle, if a cycle is detected
func DetectGraphCycle(startNode string, checkedNodes map[string]bool, edgeMap map[string]string) []string {
	currentPath := make([]string, 0)
	currentPathNodes := make(map[string]bool)
	var currentNode = startNode
	for {
		// if the current node has already been checked, no cycle is detected
		if checkedNodes[currentNode] {
			break
		}

		// if the current node is already in the current path, a cycle is detected
		if currentPathNodes[currentNode] {
			for i, node := range currentPath {
				// the cycle starts from the first occurrence of the current node
				if node == currentNode {
					return currentPath[i:]
				}
			}
		}

		// add the current node to the current path
		currentPath = append(currentPath, currentNode)
		currentPathNodes[currentNode] = true

		// get the next node
		nextNode, exists := edgeMap[currentNode]
		if !exists {
			// if there is no outgoing edge, no cycle is detected
			break
		}
		currentNode = nextNode
	}

	// mark all nodes in the current path as checked
	for node := range currentPathNodes {
		checkedNodes[node] = true
	}

	return nil
}
