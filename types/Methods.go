package types

func (f *Model) SetDefinition(definition Definition) {
	f.definition = definition
}

func (f Model) Definition() Definition {
	return f.definition
}

func (f *Definition) SetProduct(product Product) {
	f.product = product
}

func (f Definition) Product() Product {
	return f.product
}

func (f *Definition) SetConfiguration(config Configuration) {
	f.config = config
}

func (f Definition) Configuration() Configuration {
	return f.config
}

func (f *Definition) SetElements(elements Elements) {
	f.elements = elements
}

func (f Definition) Elements() Elements {
	return f.elements
}

func (f *Product) SetVersion(version Version) {
	f.version = version
}

func (f Product) Version() Version {
	return f.version
}

func (f *Version) SetMajor(major int) {
	f.major = major
}

func (f Version) Major() int {
	return f.major
}

func (f *Version) SetMinor(minor int) {
	f.minor = minor
}

func (f Version) Minor() int {
	return f.minor
}

func (f *Version) SetPatch(patch int) {
	f.patch = patch
}

func (f Version) Patch() int {
	return f.patch
}

func (f *Configuration) SetEnvironment(environment string) {
	f.environment = environment
}

func (f Configuration) Environment() string {
	return f.environment
}

func (f *Configuration) SetNetwork(network string) {
	f.network = network
}

func (f Configuration) Network() string {
	return f.network
}

func (f *Property) SetName(name string) {
	f.name = name
}

func (f Property) Name() string {
	return f.name
}

func (f *Property) SetValue(value string) {
	f.value = value
}

func (f Property) Value() string {
	return f.value
}

func (f *Element) SetFile(file string) {
	f.file = file
}

func (f Element) File() string {
	return f.file
}

func (f *Element) SetProperties(properties Properties) {
	f.properties = properties
}

func (f Element) Properties() Properties {
	return f.properties
}
