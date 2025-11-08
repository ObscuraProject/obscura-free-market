package webapp

import (
	"net/http"

	"github.com/gocraft/web"
	"github.com/ObscuraProject/obscura-free-market/modules/util"
)

func (c *Context) ViewAppDownload(w web.ResponseWriter, r *web.Request) {
	if c.ViewUser != nil && c.ViewUser.HasDownloadedApp == false {
		c.ViewUser.User.HasDownloadedApp = true
		c.ViewUser.User.Save()
	}
	http.Redirect(w, r.Request, "/tochka.apk", http.StatusFound)
}

func (c *Context) ViewAppDescription(w web.ResponseWriter, r *web.Request) {
	if c.ViewUser != nil && c.ViewUser.HasVisitedDownloadAppPage == false {
		c.ViewUser.User.HasVisitedDownloadAppPage = true
		c.ViewUser.User.Save()
	}
	util.RenderTemplate(w, "app/description", c)
}
