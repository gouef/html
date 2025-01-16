package html

import "fmt"

type Attribute struct {
	name  string
	value interface{}
}

// NewAttribute create new attribute for html element
func NewAttribute(name string, value interface{}) *Attribute {
	return &Attribute{
		name:  name,
		value: value,
	}
}

// ToHtml create html string
func (a *Attribute) ToHtml() string {
	return a.name + "=\"" + a.GetValueString() + "\""
}

// GetValueString convert value to string
func (a *Attribute) GetValueString() string {
	switch v := a.value.(type) {
	case string:
		return v
	case int:
		return fmt.Sprintf("%d", v)
	case bool:
		return fmt.Sprintf("%t", v)
	default:
		return fmt.Sprintf("%v", v)
	}
}
