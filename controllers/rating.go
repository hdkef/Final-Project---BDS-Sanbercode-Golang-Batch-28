package controllers

import (
	"bloggo/models"
	"bloggo/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RatingCtl struct {
}

// RatingGetAll godoc
// @Tags ratings
// @Summary get all ratings
// @Description get all ratings navigated by last-id and limit.
// @Produce  json
// @Param article-id path int true "ID of article"
// @Param last-id query int true "ID of the last rating in recent array of rating"
// @Param limit query int true "how many rating you want to take"
// @Success 200 {array} swagmodel.GetRating
// @Router /ratings/{article-id} [get]
func (c *RatingCtl) GetAll() gin.HandlerFunc {
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

		//get all ratings navigated by last-id and limit

		var ratingMdl models.Rating

		ratings, err := ratingMdl.GetAll(db, articleID, lastID, limit)

		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		//send
		c.JSON(http.StatusOK, ratings)
	}
}

// RatingPost godoc
// @Security AuthToken
// @Tags ratings
// @Summary create an rating
// @Description create an rating
// @Param article-id path int true "ID of article"
// @Param  rating body swagmodel.InputRating true "create an rating"
// @Accept json
// @Produce  json
// @Success 200 {object} swagmodel.Response
// @Router /ratings/{article-id} [post]
func (c *RatingCtl) Post() gin.HandlerFunc {
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
		var rating models.Rating
		err = json.NewDecoder(c.Request.Body).Decode(&rating)
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
		err = rating.Post(db, uint(articleID), &usr)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		utils.ResponseOK(c, "rating posted")
	}
}

// RatingPut godoc
// @Security AuthToken
// @Tags ratings
// @Summary update an rating
// @Description update an rating
// @Param id path int true "ID of rating"
// @Param  rating body swagmodel.InputRating true "update an rating"
// @Accept json
// @Produce  json
// @Success 200 {object} swagmodel.Response
// @Router /ratings/{id} [put]
func (c *RatingCtl) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		//get id from path
		id, err := strconv.Atoi(c.Params.ByName("id"))
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		//decode payload
		var rating models.Rating
		err = json.NewDecoder(c.Request.Body).Decode(&rating)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		//store rating to db
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

		err = rating.Put(db, id, &usr)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		utils.ResponseOK(c, "rating edited")
	}
}

// RatingDelete godoc
// @Security AuthToken
// @Tags ratings
// @Summary delete an rating
// @Description delete an rating
// @Param id path int true "ID of rating"
// @Accept json
// @Produce  json
// @Success 200 {object} swagmodel.Response
// @Router /ratings/{id} [delete]
func (c *RatingCtl) Delete() gin.HandlerFunc {
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

		var ratingMdl models.Rating

		//delete article from db
		db, err := utils.ExtractDB(c)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		err = ratingMdl.Delete(db, id, &usr)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		//send
		utils.ResponseOK(c, "rating deleted")
	}
}

// RatingGetSum godoc
// @Tags ratings
// @Summary get summed ratings
// @Description get summed ratings navigated by last-id and limit.
// @Produce  json
// @Param article-id path int true "ID of article"
// @Success 200 {object} swagmodel.GetRatingSum
// @Router /ratings/{article-id}/sum [get]
func (c *RatingCtl) GetSum() gin.HandlerFunc {
	return func(c *gin.Context) {
		//get article-id from path
		articleID, err := strconv.Atoi(c.Params.ByName("article-id"))
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		db, err := utils.ExtractDB(c)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		var ratingMdl models.Rating
		ratingSum, err := ratingMdl.GetSum(db, articleID)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		//send
		c.JSON(http.StatusOK, ratingSum)
	}
}
