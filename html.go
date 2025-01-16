package html

type Html struct {
	attributes []*Attribute
	children   []interface{}
	name       string
	empty      bool
}

func newHtml() *Html {
	return &Html{
		empty:      false,
		attributes: []*Attribute{},
		children:   []interface{}{},
	}
}

// El create html element
// Example:
// el := El("p")
func El(name string) *Html {
	el := newHtml()
	el.setName(name)

	return el
}

// setName set html element tag name
func (h *Html) setName(name string) *Html {
	h.name = name
	return h.UpdateEmpty()
}

// GetName get html element name
func (h *Html) GetName() string {
	return h.name
}

// UpdateEmpty update element status pair
func (h *Html) UpdateEmpty() *Html {
	h.empty = isEmptyElement(h.name)
	return h
}

// IsEmpty check if element is pair
func (h *Html) IsEmpty() bool {
	return h.empty
}

// AddAttribute add attribute to element
// Example:
// el := El("div")
// el.AddAttribute("class", "container")
func (h *Html) AddAttribute(name string, value interface{}) *Html {
	return h.addAttributeObject(NewAttribute(name, value))
}

func (h *Html) addAttributeObject(attribute *Attribute) *Html {
	if existAttr := h.existAttribute(attribute); existAttr != nil {
		h.attributes[*existAttr] = attribute
		return h
	}

	h.attributes = append(h.attributes, attribute)
	return h
}

func (h *Html) existAttribute(attribute *Attribute) *int {
	for index, existingAttr := range h.attributes {
		if existingAttr.name == attribute.name {
			return &index
		}
	}

	return nil
}

// AddAttributes add attributes to element
// Example:
// el := El("div")
//
//	el.AddAttributes([]map[string]interface{}{
//		{"class": "container"},
//		{"id": "main"},
//	})
func (h *Html) AddAttributes(attributes []map[string]interface{}) *Html {
	for _, attr := range attributes {
		for name, value := range attr {
			h.AddAttribute(name, value)
		}
	}

	return h
}

// GetAttributes get attributes of element
func (h *Html) GetAttributes() []*Attribute {
	return h.attributes
}

// GetAttribute get single attribute of element
func (h *Html) GetAttribute(name string) *Attribute {
	for _, attr := range h.attributes {
		if attr.name == name {
			return attr
		}
	}

	return nil
}

// RemoveAttribute remove attribute from element
func (h *Html) RemoveAttribute(name string) *Html {
	for i, attr := range h.attributes {
		if attr.name == name {
			h.attributes = append(h.attributes[:i], h.attributes[i+1:]...)
			break
		}
	}

	return h
}

// RemoveAttributes remove multiple attributes from element
func (h *Html) RemoveAttributes(names []string) *Html {
	for _, name := range names {
		h.RemoveAttribute(name)
	}

	return h
}

// AddHtml add html child to element
// Example:
// el := El("div")
// child := El("p").AddString("some text")
// el.AddHtml(child)
func (h *Html) AddHtml(child *Html) *Html {
	return h.AddChild(child)
}

// AddHtmlChildren add html children to element
// Example:
// el := El("div")
// p := El("p").AddString("some text")
// div := El("div")
// children := []*html.Html{div, p}
// el.AddHtmlChildren(children)
func (h *Html) AddHtmlChildren(children []*Html) *Html {
	for _, child := range children {
		h.AddHtml(child)
	}
	return h
}

// AddString add html children to element
// Example:
// p := El("p").AddString("some text")
func (h *Html) AddString(str string) *Html {
	var child interface{} = str
	return h.AddChild(child)
}

// AddStringChildren add string children to element
// Example:
// el := El("div")
// children := []string{"some text", "<p>another text</p>"}
// el.AddHtmlChildren(children)
func (h *Html) AddStringChildren(strings []string) *Html {
	for _, child := range strings {
		h.AddString(child)
	}

	return h
}

// AddChild add mixed child to element
func (h *Html) AddChild(child interface{}) *Html {
	switch v := child.(type) {
	case *Html:
		h.children = append(h.children, v)
	case string:
		h.children = append(h.children, v)
	default:
	}
	return h
}

// AddChildren add mixed children to element
func (h *Html) AddChildren(children []interface{}) *Html {
	for _, child := range children {
		h.AddChild(child)
	}

	return h
}

// GetChildren get html children of element
func (h *Html) GetChildren() []interface{} {
	return h.children
}

// RemoveChildren clear children node of element
func (h *Html) RemoveChildren() *Html {
	h.children = []interface{}{}
	return h
}
