package html

import "strings"

// Render render html to string
// Example:
// el := El("div")
// children := []string{"some text", "<p>another text</p>"}
// el.AddHtmlChildren(children)
// html := el.Render()
func (h *Html) Render(level ...int) string {
	indentLevel := 0
	if len(level) > 0 {
		indentLevel = level[0]
	}

	indent := strings.Repeat("\t", indentLevel)

	html := indent + h.startElement()

	if len(h.children) > 0 {
		for _, child := range h.children {
			switch v := (child).(type) {
			case string:
				html += v
			case *Html:
				html += "\n" + v.Render(indentLevel+1)
			}

		}
		html += indent + h.endElement()
	} else if len(h.children) == 0 && !h.IsEmpty() {
		html += h.endElement() + "\n"
	}

	if h.IsEmpty() {
		return indent + h.startElement()
	}

	return html
}

func (h *Html) startElement() string {
	html := "<" + h.name + h.attributesToString()

	if h.IsEmpty() {
		return html + " />"
	}

	return "<" + h.name + h.attributesToString() + ">"
}

func (h *Html) endElement() string {
	return "</" + h.name + ">"
}

func (h *Html) attributesToString() string {
	var attributes string

	for _, attr := range h.attributes {
		attributes += " " + attr.ToHtml()
	}

	return attributes
}
