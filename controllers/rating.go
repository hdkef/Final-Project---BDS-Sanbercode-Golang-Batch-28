package controllers

import "github.com/gin-gonic/gin"

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

	}
}
