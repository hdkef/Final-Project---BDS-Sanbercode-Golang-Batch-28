package controllers

import "github.com/gin-gonic/gin"

type MediaCtl struct {
}

// MediaGetAll godoc
// @Tags media
// @Summary get all media
// @Description get all media navigated by last-id and limit.
// @Produce  json
// @Param last-id query int true "ID of the last media in recent array of media"
// @Param limit query int true "how many media you want to take"
// @Success 200 {array} swagmodel.GetMedia
// @Router /media [get]
func (c *MediaCtl) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// mediaPost godoc
// @Tags media
// @Summary create an media
// @Description create an media
// @Param  media body swagmodel.InputMedia true "create an media"
// @Accept json
// @Produce  json
// @Success 200 {object} swagmodel.Response
// @Router /media [post]
func (c *MediaCtl) Post() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// MediaPut godoc
// @Tags media
// @Summary update an media
// @Description update an media
// @Param id path int true "ID of media"
// @Param  media body swagmodel.InputMedia true "update an media"
// @Accept json
// @Produce  json
// @Success 200 {object} swagmodel.Response
// @Router /media/{id} [put]
func (c *MediaCtl) Put() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// MediaDelete godoc
// @Tags media
// @Summary delete an media
// @Description delete an media
// @Param id path int true "ID of media"
// @Accept json
// @Produce  json
// @Success 200 {object} swagmodel.Response
// @Router /media/{id} [delete]
func (c *MediaCtl) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
