package types

type ElementChannel struct {
	Id      int
	Element ElementDataType
}

type ModelType struct {
	Definition DefinitionType `json:"definition:`
}

type DefinitionType struct {
	Product  ProductType       `json:"product"`
	Config   ConfigurationType `json:"config"`
	Elements ElementsType      `json:"elements"`
}

type ProductType struct {
	Version VersionType `json:"version"`
}

type VersionType struct {
	Major int `json:"major"`
	Minor int `json:"minor"`
	Patch int `json:"patch"`
}

type ConfigurationType struct {
	Environment string `json:"environment"`
	Network     string `json:"network"`
}

type PropertyType struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type PropertiesType []PropertyType

type ElementType struct {
	File       string          `json:"file"`
	Properties PropertiesType  `json:"properties"`
	Data       ElementDataType `json:"data"`
}

type ElementsType []ElementType

type ElementDataType struct {
	Id      int         `json:"id"`
	Name    string      `json:"name"`
	Type    string      `json:"type"`
	Product ProductType `json:"product"`
	Tiers   TiersType   `json:"tier"`
}

type TierType struct {
	Id          int      `json:"id"`
	MemberCount int      `json:"member-count"`
	Node        NodeType `json:"node"`
	Components  []string `json:"components"`
}

type TiersType []TierType

type NodeType struct {
	Name    string `json:"name"`
	Version int    `json:"version"`
}
