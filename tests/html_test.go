package tests

import (
	html "github.com/gouef/html"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHtml(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		el := html.El("p")

		assert.Equal(t, "p", el.GetName())
		assert.Equal(t, false, el.IsEmpty())
	})
	t.Run("Empty", func(t *testing.T) {
		el := html.El("br")

		assert.Equal(t, "br", el.GetName())
		assert.Equal(t, true, el.IsEmpty())
	})

	t.Run("Add attribute", func(t *testing.T) {
		el := html.El("div")

		el.AddAttributes([]map[string]interface{}{
			{"class": "container"},
			{"id": "main"},
		})

		assert.Equal(t, "div", el.GetName())
		assert.Equal(t, false, el.IsEmpty())
		assert.Equal(t, []*html.Attribute{
			html.NewAttribute("class", "container"),
			html.NewAttribute("id", "main"),
		}, el.GetAttributes())
		assert.Equal(t, html.NewAttribute("class", "container"), el.GetAttribute("class"))
	})

	t.Run("Rewrite attribute", func(t *testing.T) {
		el := html.El("div")

		el.AddAttributes([]map[string]interface{}{
			{"class": "container"},
			{"id": "main"},
		})

		assert.Equal(t, "div", el.GetName())
		assert.Equal(t, false, el.IsEmpty())
		assert.Equal(t, []*html.Attribute{
			html.NewAttribute("class", "container"),
			html.NewAttribute("id", "main"),
		}, el.GetAttributes())
		assert.Equal(t, html.NewAttribute("class", "container"), el.GetAttribute("class"))
		el.AddAttribute("class", "container-fluid")
		assert.Equal(t, html.NewAttribute("class", "container-fluid"), el.GetAttribute("class"))
	})

	t.Run("Remove attribute", func(t *testing.T) {
		el := html.El("div")

		el.AddAttributes([]map[string]interface{}{
			{"class": "container"},
			{"id": "main"},
			{"data-id": "main"},
			{"data-name": "main"},
			{"data-img": "img.gouef.cz/images/test.png"},
		})

		assert.Equal(t, "div", el.GetName())
		assert.Equal(t, false, el.IsEmpty())

		assert.Equal(t, []*html.Attribute{
			html.NewAttribute("class", "container"),
			html.NewAttribute("id", "main"),
			html.NewAttribute("data-id", "main"),
			html.NewAttribute("data-name", "main"),
			html.NewAttribute("data-img", "img.gouef.cz/images/test.png"),
		}, el.GetAttributes())

		el.RemoveAttributes([]string{
			"data-id",
			"data-name",
			"data-img",
		})

		assert.Equal(t, []*html.Attribute{
			html.NewAttribute("class", "container"),
			html.NewAttribute("id", "main"),
		}, el.GetAttributes())
		assert.Equal(t, html.NewAttribute("class", "container"), el.GetAttribute("class"))
	})

	t.Run("Add children (Html + string)", func(t *testing.T) {
		el := html.El("div")
		el.AddAttributes([]map[string]interface{}{
			{"class": "container"},
		})

		child1 := html.El("p").AddString("Paragraph text")
		child2 := html.El("span").AddString("Span text")
		textNode := "Simple text node"

		assert.Equal(t, `<p>Paragraph text</p>`, child1.Render())

		el.AddChildren([]interface{}{child1, textNode, child2})

		assert.Equal(t, 3, len(el.GetChildren()))
		assert.Equal(t, child1, el.GetChildren()[0])
		assert.Equal(t, textNode, el.GetChildren()[1])
		assert.Equal(t, child2, el.GetChildren()[2])

		rendered := el.Render()
		expected := "<div class=\"container\">\n\t<p>Paragraph text\t</p>Simple text node\n\t<span>Span text\t</span></div>"
		assert.Equal(t, expected, rendered)
	})

	t.Run("Add string children", func(t *testing.T) {
		el := html.El("div")
		el.AddAttributes([]map[string]interface{}{
			{"class": "container"},
		})

		child1 := html.El("p").AddStringChildren([]string{"Paragraph text", "<span>Test</span>"})
		child2 := html.El("span").AddString("Span text")
		textNode := "Simple text node"

		assert.Equal(t, `<p>Paragraph text<span>Test</span></p>`, child1.Render())

		el.AddChildren([]interface{}{child1, textNode, child2})

		assert.Equal(t, 3, len(el.GetChildren()))
		assert.Equal(t, child1, el.GetChildren()[0])
		assert.Equal(t, textNode, el.GetChildren()[1])
		assert.Equal(t, child2, el.GetChildren()[2])

		rendered := el.Render()
		expected := "<div class=\"container\">\n\t<p>Paragraph text<span>Test</span>\t</p>Simple text node\n\t<span>Span text\t</span></div>"
		assert.Equal(t, expected, rendered)
	})

	t.Run("Remove children", func(t *testing.T) {
		el := html.El("div")

		child1 := html.El("p").AddString("Paragraph text")
		child2 := html.El("span").AddString("Span text")
		textNode := "Simple text node"

		el.AddChildren([]interface{}{child1, textNode, child2})
		assert.Equal(t, 3, len(el.GetChildren()))

		el.RemoveChildren()
		assert.Empty(t, el.GetChildren())
	})

	t.Run("Get non-exist attribute", func(t *testing.T) {
		el := html.El("div")

		child1 := html.El("p").AddString("Paragraph text")
		child2 := html.El("span").AddString("Span text")
		textNode := "Simple text node"

		el.AddChildren([]interface{}{child1, textNode, child2})
		assert.Equal(t, 3, len(el.GetChildren()))

		el.RemoveChildren()
		assert.Empty(t, el.GetChildren())
		assert.Nil(t, el.GetAttribute("data-id"))
	})
}
