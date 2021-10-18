package controllers

import (
	"bloggo/models"
	"bloggo/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type InboxCtl struct {
}

// InboxGetAll godoc
// @Security AuthToken
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
		//get last-id and limit from query

		lastID, err := strconv.Atoi(c.Query("last-id"))
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		limit, err := strconv.Atoi(c.Query("limit"))
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		usr, err := utils.ExtractUserContext(c)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		db, err := utils.ExtractDB(c)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		//get all inboxes navigated by last-id and limit

		var inboxMdl models.Inbox

		inboxes, err := inboxMdl.GetAll(db, lastID, limit, &usr)

		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		//send
		c.JSON(http.StatusOK, inboxes)
	}
}

// InboxPost godoc
// @Security AuthToken
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
		//get receiver-id from path
		receiverID, err := strconv.Atoi(c.Params.ByName("receiver-id"))
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		//get user detail from token
		usr, err := utils.ExtractUserContext(c)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		//decode payload
		var inbox models.Inbox
		err = json.NewDecoder(c.Request.Body).Decode(&inbox)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		db, err := utils.ExtractDB(c)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		//store to db
		err = inbox.Post(db, uint(receiverID), &usr)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		utils.ResponseOK(c, "inbox posted")
	}
}

// InboxPut godoc
// @Security AuthToken
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
		//get id from path
		id, err := strconv.Atoi(c.Params.ByName("id"))
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		//decode payload
		var inbox models.Inbox
		err = json.NewDecoder(c.Request.Body).Decode(&inbox)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		//store inbox to db
		db, err := utils.ExtractDB(c)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		//get user detail from token
		usr, err := utils.ExtractUserContext(c)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		err = inbox.Put(db, id, &usr)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		utils.ResponseOK(c, "inbox edited")
	}
}

// InboxDelete godoc
// @Security AuthToken
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
		//get id from path
		id, err := strconv.Atoi(c.Params.ByName("id"))
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		//get user detail from token
		usr, err := utils.ExtractUserContext(c)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		var inboxMdl models.Inbox

		//delete article from db
		db, err := utils.ExtractDB(c)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		err = inboxMdl.Delete(db, id, &usr)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		//send
		utils.ResponseOK(c, "inbox deleted")
	}
}
