package types

import "fmt"

func (f *Model) ToString() string {
	s := fmt.Sprintf("%+v", f)
	return s
}

func (f *Definition) ToString() string {
	s := fmt.Sprintf("%+v", f)
	return s
}

func (f *Product) ToString() string {
	s := fmt.Sprintf("%+v", f)
	return s
}

func (f *Version) ToString() string {
	s := fmt.Sprintf("%+v", f)
	return s
}

func (f *Configuration) ToString() string {
	s := fmt.Sprintf("%+v", f)
	return s
}

func (f *Property) ToString() string {
	s := fmt.Sprintf("%+v", f)
	return s
}

func (f *Element) ToString() string {
	s := fmt.Sprintf("%+v", f)
	return s
}
