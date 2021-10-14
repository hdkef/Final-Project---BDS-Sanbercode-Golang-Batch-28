package controllers

import (
	"bloggo/models"
	"bloggo/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentCtl struct {
}

// CommentGetAll godoc
// @Tags comments
// @Summary get all comments
// @Description get all comments navigated by last-id and limit.
// @Produce  json
// @Param article-id path int true "ID of article"
// @Param last-id query int true "ID of the last comment in recent array of comment"
// @Param limit query int true "how many comment you want to take"
// @Success 200 {array} swagmodel.GetComment
// @Router /comments/{article-id} [get]
func (c *CommentCtl) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		//get last-id and limit from query and article-id from path

		articleID, err := strconv.Atoi(c.Params.ByName("article-id"))
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

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

		db, err := utils.ExtractDB(c)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		//get all comments navigated by last-id and limit

		var commentMdl models.Comment

		comments, err := commentMdl.GetAll(db, articleID, lastID, limit)

		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		//send
		c.JSON(http.StatusOK, comments)
	}
}

// CommentPost godoc
// @Tags comments
// @Summary create an comment
// @Description create an comment
// @Param article-id path int true "ID of article"
// @Param  comment body swagmodel.InputComment true "create an comment"
// @Accept json
// @Produce  json
// @Success 200 {object} swagmodel.Response
// @Router /comments/{article-id} [post]
func (c *CommentCtl) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		//get article-id from path
		articleID, err := strconv.Atoi(c.Params.ByName("article-id"))
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
		var comment models.Comment
		err = json.NewDecoder(c.Request.Body).Decode(&comment)
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
		err = comment.Post(db, uint(articleID), &usr)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		utils.ResponseOK(c, "comment posted")
	}
}

// CommentPut godoc
// @Tags comments
// @Summary update an comment
// @Description update an comment
// @Param id path int true "ID of comment"
// @Param  comment body swagmodel.InputComment true "update an comment"
// @Accept json
// @Produce  json
// @Success 200 {object} swagmodel.Response
// @Router /comments/{id} [put]
func (c *CommentCtl) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		//get id from path
		id, err := strconv.Atoi(c.Params.ByName("id"))
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		//decode payload
		var comment models.Comment
		err = json.NewDecoder(c.Request.Body).Decode(&comment)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		//store comment to db
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

		err = comment.Put(db, id, &usr)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		utils.ResponseOK(c, "comment updated")
	}
}

// CommentDelete godoc
// @Tags comments
// @Summary delete an comment
// @Description delete an comment
// @Param id path int true "ID of comment"
// @Accept json
// @Produce  json
// @Success 200 {object} swagmodel.Response
// @Router /comments/{id} [delete]
func (c *CommentCtl) Delete() gin.HandlerFunc {
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

		var commentMdl models.Comment

		//delete article from db
		db, err := utils.ExtractDB(c)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		err = commentMdl.Delete(db, id, &usr)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		//send
		utils.ResponseOK(c, "comment deleted")
	}
}
