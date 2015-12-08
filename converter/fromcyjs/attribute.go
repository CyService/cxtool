package fromcyjs
import (
	"github.com/cytoscape-ci/cxtool/cx"
	"errors"
)

func getNetworkAttribute(cyjsData map[string]interface{}) (attrs []cx.NetworkAttribute) {

	for attrName, cyjsAttr := range cyjsData {
		cxAttr := cx.NetworkAttribute{N:attrName, V:cyjsAttr}
		attrs = append(attrs, cxAttr)
	}

	return attrs
}

// Convert Cytoscpae.js's data object into CX Attributes
func getAttribute(cyjsData map[string]interface{}) (attrs []cx.Attribute, parseError error) {

	// Check ID exists or not
	id, exists := cyjsData["id"]
	if exists == false {
		return attrs, errors.New("No ID!")
	}

	po, err := getIdNumber(id.(string))

	if err != nil {
		return attrs, err
	}

	for attrName, cyjsAttr := range cyjsData {
		// ID is not necessary
		if attrName == "id" {
			continue
		}

		cxAttr := cx.Attribute{PO: po, N:attrName, V:cyjsAttr}
		attrs = append(attrs, cxAttr)
	}

	return attrs, nil
}
