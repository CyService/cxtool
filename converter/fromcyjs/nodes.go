package fromcyjs

import (
	"github.com/cytoscape-ci/cxtool/cyjs"
	"github.com/cytoscape-ci/cxtool/cx"
)


func getNodes(cyjsNodes []cyjs.CyJSNode) (nodes cx.Nodes, attrs []cx.Attribute) {

	var cxNodes []cx.Node

	for _, n := range cyjsNodes {

		idStr, exists := n.Data["id"]
		if !exists {
			continue
		}

		id, err := getIdNumber(idStr.(string))
		if err != nil {
			continue
		}

		node := cx.Node{ID: id, N: idStr.(string)}
		cxNodes = append(cxNodes, node)

		nodeAttrs, err := getAttribute(n.Data)

		if err == nil {
			attrs = append(attrs, nodeAttrs...)
		}
	}

	nodes = cx.Nodes{NodesList:cxNodes}
	return nodes, attrs
}
