package controllers

import "github.com/gin-gonic/gin"

type ArticleCtl struct {
}

// ArticleGetAll godoc
// @Tags articles
// @Summary get all articles
// @Description get all articles navigated by last-id and limit.
// @Produce  json
// @Param last-id query int true "ID of the last article in recent array of article"
// @Param limit query int true "how many article you want to take"
// @Success 200 {array} swagmodel.GetArticle
// @Router /articles [get]
func (c *ArticleCtl) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// ArticleGetOne godoc
// @Tags articles
// @Summary get one article
// @Description get detail of article specified by id
// @Produce  json
// @Param id path int true "ID of article"
// @Success 200 {object} swagmodel.GetArticle
// @Router /articles/{id} [get]
func (c *ArticleCtl) GetOne() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// ArticlePost godoc
// @Tags articles
// @Summary create an article
// @Description create an article
// @Param  article body swagmodel.InputArticle true "create an article"
// @Accept json
// @Produce  json
// @Success 200 {object} swagmodel.Response
// @Router /articles [post]
func (c *ArticleCtl) Post() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// ArticlePut godoc
// @Tags articles
// @Summary update an article
// @Description update an article
// @Param id path int true "ID of article"
// @Param  article body swagmodel.InputArticle true "update an article"
// @Accept json
// @Produce  json
// @Success 200 {object} swagmodel.Response
// @Router /articles/{id} [put]
func (c *ArticleCtl) Put() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// ArticleDelete godoc
// @Tags articles
// @Summary delete an article
// @Description delete an article
// @Param id path int true "ID of article"
// @Accept json
// @Produce  json
// @Success 200 {object} swagmodel.Response
// @Router /articles/{id} [delete]
func (c *ArticleCtl) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
