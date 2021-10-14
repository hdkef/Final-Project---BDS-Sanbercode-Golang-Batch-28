package controllers

import (
	"bloggo/models"
	"bloggo/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DraftCtl struct {
}

// DraftGetAll godoc
// @Tags drafts
// @Summary get all drafts
// @Description get all drafts navigated by last-id and limit.
// @Produce  json
// @Param last-id query int true "ID of the last draft in recent array of draft"
// @Param limit query int true "how many draft you want to take"
// @Success 200 {array} swagmodel.GetDraft
// @Router /drafts [get]
func (c *DraftCtl) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := utils.ExtractDB(c)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		var draftMdl models.Draft

		drafts, err := draftMdl.GetAll(db)

		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, drafts)
	}
}

// DraftGetOne godoc
// @Tags drafts
// @Summary get one draft
// @Description get detail of draft specified by id
// @Produce  json
// @Param id path int true "ID of draft"
// @Success 200 {object} swagmodel.GetDraft
// @Router /drafts/{id} [get]
func (c *DraftCtl) GetOne() gin.HandlerFunc {
	return func(c *gin.Context) {

		db, err := utils.ExtractDB(c)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		//get id from path
		id, err := strconv.Atoi(c.Params.ByName("id"))
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		//get draft from db
		var draftMdl models.Draft
		draft, err := draftMdl.GetOne(db, uint(id))
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		//send
		c.JSON(http.StatusOK, draft)
	}
}

// DraftPost godoc
// @Tags drafts
// @Summary create an draft
// @Description create an draft
// @Param  draft body swagmodel.InputDraft true "create an draft"
// @Accept json
// @Produce  json
// @Success 200 {object} swagmodel.Response
// @Router /drafts [post]
func (c *DraftCtl) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var draft models.Draft

		err := json.NewDecoder(c.Request.Body).Decode(&draft)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		db, err := utils.ExtractDB(c)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		usr, err := utils.ExtractUserContext(c)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		err = draft.Post(db, &usr)

		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		utils.ResponseOK(c, "draft created")
	}
}

// DraftPut godoc
// @Tags drafts
// @Summary update an draft
// @Description update an draft
// @Param id path int true "ID of draft"
// @Param  draft body swagmodel.InputDraft true "update an draft"
// @Accept json
// @Produce  json
// @Success 200 {object} swagmodel.Response
// @Router /drafts/{id} [put]
func (c *DraftCtl) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		var draft models.Draft

		err := json.NewDecoder(c.Request.Body).Decode(&draft)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		db, err := utils.ExtractDB(c)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		id, err := strconv.Atoi(c.Params.ByName("id"))
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		usr, err := utils.ExtractUserContext(c)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		err = draft.Put(db, uint(id), &usr)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		utils.ResponseOK(c, "draft updated")
	}
}

// DraftDelete godoc
// @Tags drafts
// @Summary delete an draft
// @Description delete an draft
// @Param id path int true "ID of draft"
// @Accept json
// @Produce  json
// @Success 200 {object} swagmodel.Response
// @Router /drafts/{id} [delete]
func (c *DraftCtl) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := utils.ExtractDB(c)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		var draftMdl models.Draft

		id, err := strconv.Atoi(c.Params.ByName("id"))
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		usr, err := utils.ExtractUserContext(c)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		err = draftMdl.Delete(db, uint(id), &usr)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		utils.ResponseOK(c, "draft deleted")
	}
}
