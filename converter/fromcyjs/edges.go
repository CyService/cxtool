package fromcyjs
import (
	"github.com/cytoscape-ci/cxtool/cyjs"
	"github.com/cytoscape-ci/cxtool/cx"
	"errors"
)


// Creates both edges and their attributes.
func getEdges(cyjsEdges []cyjs.CyJSEdge) (cxEdges cx.Edges, edgeAttrs []cx.Attribute) {

	var edges []cx.Edge

	for _, cyjsEdge := range cyjsEdges {

		// Create an edge object
		edge, err := buildEdge(cyjsEdge)

		if err == nil {
			edges = append(edges, edge)
		}

		// Create attribute for the edge
		attrs, err := getAttribute(cyjsEdge.Data)

		if err == nil {
			edgeAttrs = append(edgeAttrs, attrs...)
		}
	}

	cxEdges = cx.Edges{EdgeList:edges}

	return cxEdges, edgeAttrs
}

func buildEdge(cyjsEdge cyjs.CyJSEdge) (edge cx.Edge, edgeErr error) {
	data := cyjsEdge.Data

	idStr, exists := data[cyjs.Id]

	if !exists {
		return edge, errors.New("Missing Edge ID.")
	}

	id, err := getIdNumber(idStr.(string))

	if err != nil {
		return edge, err
	}

	source := data[cyjs.Source]
	sourceId, err := getIdNumber(source.(string))

	target := data[cyjs.Target]
	targetId, err := getIdNumber(target.(string))

	interaction, exists := data[cyjs.Interaction]

	if exists {
		edge = cx.Edge{ID:id, S: sourceId, T:targetId, I:interaction.(string)}
	} else {
		edge = cx.Edge{ID:id, S: sourceId, T:targetId}
	}

	return edge, nil
}


