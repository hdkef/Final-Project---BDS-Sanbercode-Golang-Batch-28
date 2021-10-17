package controllers

import (
	"bloggo/models"
	"bloggo/utils"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

var MEDIAPTH string = os.Getenv("MEDIAPTH")

type MediaCtl struct {
}

// MediaGetAll godoc
// @Security AuthToken
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
		//get id from path, last-id and limit from query
		id, err := strconv.Atoi(c.Params.ByName("id"))
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		lastID, err := strconv.Atoi(c.Params.ByName("last-id"))
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		limit, err := strconv.Atoi(c.Params.ByName("limit"))
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
		db, err := utils.ExtractDB(c)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		var mediaMdl models.Media
		//get all media
		media, err := mediaMdl.GetAll(db, &usr, uint(id), lastID, limit)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		//send
		c.JSON(http.StatusOK, media)
	}
}

// mediaPost godoc
// @Security AuthToken
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
		//get uploaded file
		file, err := c.FormFile("file")
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		caption := c.PostForm("caption")
		alt := c.PostForm("alt")
		//store upload file and get url
		err = c.SaveUploadedFile(file, MEDIAPTH+file.Filename)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		//new media object
		var newMedia models.Media = models.Media{
			Url:     MEDIAPTH + file.Filename,
			Caption: caption,
			Alt:     alt,
		}
		//get user detail from token
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
		//save to db
		err = newMedia.Post(db, &usr)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		//send
		utils.ResponseOK(c, "media uploaded")
	}
}

// MediaDelete godoc
// @Security AuthToken
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

		var mediaMdl models.Media

		//delete media from db and from disk
		db, err := utils.ExtractDB(c)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		err = mediaMdl.Delete(db, &usr, uint(id))
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		//send
		utils.ResponseOK(c, "media deleted")
	}
}
