package webapp

import (
	"github.com/gocraft/web"

	"github.com/ObscuraProject/obscura-free-market/modules/util"
)

func (c *Context) ViewAboutUser(w web.ResponseWriter, r *web.Request) {
	c.SelectedSection = "info"
	if len(r.URL.Query()["section"]) > 0 {
		c.SelectedSection = r.URL.Query()["section"][0]
	}
	util.RenderTemplateOrAPIResponse(w, r, "user/about", c, c.IsAPIRequest)
}
