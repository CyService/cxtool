package cx

const (
	Mappings = "mappings"
	Passthrough = "PASSTHROUGH"
	Discrete = "DISCRETE"
	Continuous = "CONTINUOUS"
)


type VisualProperty struct {
	PropertiesOf string `json:"properties_of"`
	AppliesTo    int    `json:"applies_to"`
	View         int    `json:"view"`

	Properties map[string]string `json:"properties"`
	Mappings map[string]Mapping `json:"mappings"`
}

type Mapping struct {

	Type string `json:"type"`
	Definition string `json:"definition"`
}

