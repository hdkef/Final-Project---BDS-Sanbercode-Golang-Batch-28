package controllers

import "github.com/gin-gonic/gin"

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
		//get article-id, last-id, and limit
		//get all comments navigated by last-id and limit
		//send
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

	}
}
