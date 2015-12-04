package types

type ElementChannel struct {
	id      int
	element Model
}

type Model struct {
	definition Definition `json:"definition:`
}

type Definition struct {
	product  Product       `json:"product"`
	config   Configuration `json:"config"`
	elements Elements      `json:"elements"`
}

type Product struct {
	version Version `json:"version"`
}

type Version struct {
	major int `json:"major"`
	minor int `json:"minor"`
	patch int `json:"patch"`
}

type Configuration struct {
	environment string `json:"environment"`
	network     string `json:"network"`
}

type Property struct {
	name  string `json:"name"`
	value string `json:"value"`
}

type Properties []Property

type Element struct {
	file       string     `json:"file"`
	properties Properties `json:"properties"`
}

type Elements []Element
