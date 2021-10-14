package controllers

import (
	"bloggo/models"
	"bloggo/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

		//get last-id and limit

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

		//get all articles navigated by last-id and limit

		var articleMdl models.Article

		articles, err := articleMdl.GetAll(db, lastID, limit)

		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		//send
		c.JSON(http.StatusOK, articles)
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

		//get article from db
		var articleMdl models.Article
		article, err := articleMdl.GetOne(db, uint(id))
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		//send
		c.JSON(http.StatusOK, article)
	}
}

// ArticlePost godoc
// @Security AuthToken
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

		//decode payload
		var article models.Article

		err := json.NewDecoder(c.Request.Body).Decode(&article)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

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

		//post to db
		err = article.Post(db, &usr)

		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		//send
		utils.ResponseOK(c, "article created")
	}
}

// ArticlePut godoc
// @Security AuthToken
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

		//decode payload
		var article models.Article
		err := json.NewDecoder(c.Request.Body).Decode(&article)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

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

		//get user detail from token
		usr, err := utils.ExtractUserContext(c)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		//update article to db
		err = article.Put(db, uint(id), &usr)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		//send
		utils.ResponseOK(c, "article updated")
	}
}

// ArticleDelete godoc
// @Security AuthToken
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
		db, err := utils.ExtractDB(c)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		var articleMdl models.Article

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

		//delete article from db
		err = articleMdl.Delete(db, uint(id), &usr)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		//send
		utils.ResponseOK(c, "article deleted")
	}
}
