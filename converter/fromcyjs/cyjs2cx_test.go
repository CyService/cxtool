package fromcyjs_test

import (
	"testing"
	"bytes"
	"io"
	"encoding/json"
	"github.com/cytoscape-ci/cxtool/converter/fromcyjs"
	"github.com/cytoscape-ci/cxtool/cx"
)

func TestCyjs2Cx(t *testing.T) {

	output := new(bytes.Buffer)
	resultWriter := io.Writer(output)

	c2c := fromcyjs.Cyjs2Cx{W: &resultWriter}
	c2c.Convert("../../test_data/ecoli_cyjs_format.json")

	result := output.String()

	var cxData []map[string]interface{}

	err := json.Unmarshal([]byte(result), &cxData)
	if err != nil {
		t.Error("error:", err)
		return
	}

	for _, aspect := range cxData {

		t.Log(aspect)
	}
//	cxNodes := cxNetwork.Nodes
//	cxEdges := cxNetwork.Edges
//
//	const numNodesExpected = 138
//	const numEdgesExpected = 128
//
//	numNodes := len(cxNodes.NodesList)
//	numEdges := len(cxEdges.EdgeList)
//
//	t.Log("Num nodes = ", numNodes)
//
//	t.Log(cxNetwork.EdgeAttributes)
//
//	edgeAttrsJSON, err := json.Marshal(cxNetwork.EdgeAttributes)
//	t.Log(edgeAttrsJSON)
//
//	if numNodes != numNodesExpected {
//		t.Error("Nodes Expected: ", numNodesExpected, ", Actual: ", numNodes)
//		t.Fail()
//		return
//	}
//
//	if numEdges != numEdgesExpected {
//		t.Error("Edges Expected: ", numEdgesExpected, ", Actual: ", numEdges)
//		t.Fail()
//		return
//	}
//
//	nodeAttrs := cxNetwork.NodeAttributes
//	edgeAttrs := cxNetwork.EdgeAttributes
//
//	numNodeAttrs := len(nodeAttrs)
//	numEdgeAttrs := len(edgeAttrs)
//
//	// TODO: add more test cases for contents of attributes...
//
//	if numNodeAttrs != 1141 {
//		t.Error("Node Attr Expected: ", numNodesExpected, ", Actual: ", numNodeAttrs)
//		t.Fail()
//	}
//
//	if numEdgeAttrs != 2048 {
//		t.Error("Edge Attr Expected: ", numEdgesExpected, ", Actual: ", numEdgeAttrs)
//		t.Fail()
//	}
//
//	t.Log(edgeAttrs)
	t.Log(result)
}