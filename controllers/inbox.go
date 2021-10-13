package controllers

import "github.com/gin-gonic/gin"

type InboxCtl struct {
}

// InboxGetAll godoc
// @Tags inboxes
// @Summary get all inboxes
// @Description get all inboxes navigated by last-id and limit.
// @Produce  json
// @Param receiver-id path int true "ID of user / receiver"
// @Param last-id query int true "ID of the last inbox in recent array of inboxes"
// @Param limit query int true "how many inboxes you want to take"
// @Success 200 {array} swagmodel.GetInbox
// @Router /inboxes/{receiver-id} [get]
func (c *InboxCtl) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// InboxPost godoc
// @Tags inboxes
// @Summary create an inbox
// @Description create an inbox
// @Param receiver-id path int true "ID of user / receiver"
// @Param  inbox body swagmodel.InputInbox true "create an inbox"
// @Accept json
// @Produce  json
// @Success 200 {object} swagmodel.Response
// @Router /inboxes/{receiver-id} [post]
func (c *InboxCtl) Post() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// InboxPut godoc
// @Tags inboxes
// @Summary update an inbox
// @Description update an inbox
// @Param id path int true "ID of inbox"
// @Param  inbox body swagmodel.InputInbox true "update an inbox"
// @Accept json
// @Produce  json
// @Success 200 {object} swagmodel.Response
// @Router /inboxes/{id} [put]
func (c *InboxCtl) Put() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// InboxDelete godoc
// @Tags inboxes
// @Summary delete an inbox
// @Description delete an inbox
// @Param id path int true "ID of inbox"
// @Accept json
// @Produce  json
// @Success 200 {object} swagmodel.Response
// @Router /inboxes/{id} [delete]
func (c *InboxCtl) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
