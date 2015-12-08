package fromcyjs

import (
	"github.com/cytoscape-ci/cxtool/cyjs"
	"github.com/cytoscape-ci/cxtool/cx"
)

func getCx(cyjsData cyjs.CyJS) (cx.CX) {
	// Convert network attributes
	netAttr := cyjsData.Data
	cxNetAttr := getNetworkAttribute(netAttr)

	// Cytoscape.js Nodes
	cyjsNodes := cyjsData.Elements.Nodes

	// Cytoscape.js Edges
	cyjsEdges := cyjsData.Elements.Edges

	cxNodes, nodeAttrs := getNodes(cyjsNodes)
	cxEdges, edgeAttrs := getEdges(cyjsEdges)

	cxData := cx.CX{
		NetworkAttributes:cxNetAttr, Nodes:cxNodes, Edges:cxEdges,
		NodeAttributes:nodeAttrs, EdgeAttributes:edgeAttrs}

	return cxData
}
