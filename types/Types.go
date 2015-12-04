package types

type ElementChannel struct {
	Id      int
	Element ModelType
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
	File       string         `json:"file"`
	Properties PropertiesType `json:"properties"`
}

type ElementsType []ElementType
