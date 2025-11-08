package webapp

import (
	"github.com/gocraft/web"

	"github.com/ObscuraProject/obscura-free-market/modules/util"
)

func (c *Context) ViewAdminListReferralPayments(w web.ResponseWriter, r *web.Request) {
	c.ReferralPayments = FindReferralPayments()
	util.RenderTemplate(w, "referral/admin/payments", c)
}
