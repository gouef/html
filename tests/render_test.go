package tests

import (
	"github.com/gouef/html"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRender(t *testing.T) {

	t.Run("Render HTML with children and attributes", func(t *testing.T) {
		el := html.El("div").AddAttributes([]map[string]interface{}{
			{"class": "container"},
			{"id": "main"},
		})

		assert.Equal(t, "div", el.GetName())
		assert.False(t, el.IsEmpty())
		assert.Equal(t, []*html.Attribute{
			html.NewAttribute("class", "container"),
			html.NewAttribute("id", "main"),
		}, el.GetAttributes())

		div := html.El("div").AddAttribute("class", "container")
		p := html.El("p").AddAttribute("class", "ppp").AddAttributes([]map[string]interface{}{
			{"data-float": 7.14},
			{"data-int": 11},
			{"data-bool": true},
		})

		children := []*html.Html{div, p}
		el.AddHtmlChildren(children)
		children2 := []interface{}{div, p}

		assert.Equal(t, children2, el.GetChildren())
		assert.Equal(t, html.NewAttribute("class", "container"), el.GetAttribute("class"))

		expected := "<div class=\"container\" id=\"main\">\n\t<div class=\"container\"></div>\n\n\t<p class=\"ppp\" data-float=\"7.14\" data-int=\"11\" data-bool=\"true\"></p>\n</div>"

		assert.Equal(t, expected, el.Render())
	})

	t.Run("Render self-closing HTML tag", func(t *testing.T) {
		el := html.El("img").AddAttributes([]map[string]interface{}{
			{"class": "container"},
			{"id": "main"},
			{"src": "https://gouef.github.com/images/test.png"},
		})

		assert.Equal(t, "img", el.GetName())
		assert.True(t, el.IsEmpty())
		assert.Equal(t, []*html.Attribute{
			html.NewAttribute("class", "container"),
			html.NewAttribute("id", "main"),
			html.NewAttribute("src", "https://gouef.github.com/images/test.png"),
		}, el.GetAttributes())

		assert.Equal(t, html.NewAttribute("class", "container"), el.GetAttribute("class"))

		expected := `<img class="container" id="main" src="https://gouef.github.com/images/test.png" />`

		assert.Equal(t, expected, el.Render())
	})
}
