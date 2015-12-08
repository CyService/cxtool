package fromcyjs

import (
	"github.com/cytoscape-ci/cxtool/cyjs"
	"github.com/cytoscape-ci/cxtool/cx"
)

func getCx(cyjsData cyjs.CyJS) (cxData []interface{}) {
	// Convert network attributes
	netAttr := cyjsData.Data
	cxNetAttr := getNetworkAttribute(netAttr)

	cxData = append(cxData, wrapEntry(cx.NetworkAttributesTag, cxNetAttr))


	// Cytoscape.js Nodes
	cyjsNodes := cyjsData.Elements.Nodes

	// Cytoscape.js Edges
	cyjsEdges := cyjsData.Elements.Edges

	cxNodes, nodeAttrs := getNodes(cyjsNodes)
	cxEdges, edgeAttrs := getEdges(cyjsEdges)

	cxData = append(cxData, cxNodes)
	cxData = append(cxData, cxEdges)
	cxData = append(cxData, wrapEntry(cx.NodeAttributesTag, nodeAttrs))
	cxData = append(cxData, wrapEntry(cx.EdgeAttributesTag, edgeAttrs))

	return cxData
}

func wrapEntry(tag string, data interface{}) (map[string]interface{}) {
	wrapped := make(map[string]interface{})
	wrapped[tag] = data
	return wrapped
}
