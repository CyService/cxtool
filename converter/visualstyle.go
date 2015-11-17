package converter


const (
	VisualProperties string = "visualProperties"
)

type VisualProperty struct {

	PropertiesOf string `json:"properties_of"`
	AppliesTo    int `json:"applies_to"`
	View         int `json:"view"`

	Properties   map[string]string `json:"properties"`
}

