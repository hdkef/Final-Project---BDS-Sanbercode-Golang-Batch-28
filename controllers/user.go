package controllers

import (
	"bloggo/models"
	"bloggo/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserCtl struct {
}

// userGetAll godoc
// @Security AuthToken
// @Tags users
// @Summary get all users
// @Description get all users navigated by last-id and limit.
// @Produce  json
// @Param last-id query int true "ID of the last user in recent array of user"
// @Param limit query int true "how many user you want to take"
// @Success 200 {array} swagmodel.GetUser
// @Router /users [get]
func (c *UserCtl) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		usr, err := utils.ExtractUserContext(c)
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

		var usrModel models.User
		//get user from db
		users, err := usrModel.GetAll(db, &usr, lastID, limit)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		//send
		c.JSON(http.StatusOK, users)
	}
}

// UserGetOne godoc
// @Security AuthToken
// @Tags users
// @Summary get one user
// @Description get detail of user specified by id
// @Produce  json
// @Param id path int true "ID of user"
// @Success 200 {object} swagmodel.GetUser
// @Router /users/{id} [get]
func (c *UserCtl) GetOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		//get id from path
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
		db, err := utils.ExtractDB(c)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		//get user from db
		var usrModel models.User
		user, err := usrModel.GetOne(db, &usr, uint(id))
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		//send
		c.JSON(http.StatusOK, user)
	}
}

// UserGetOnePublic godoc
// @Tags users
// @Summary get one user public information (without password information)
// @Description get detail of user specified by id
// @Produce  json
// @Param id path int true "ID of user"
// @Success 200 {object} swagmodel.GetUserPublic
// @Router /users/{id}/public [get]
func (c *UserCtl) GetOnePublic() gin.HandlerFunc {
	return func(c *gin.Context) {
		//get id from path
		id, err := strconv.Atoi(c.Params.ByName("id"))
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		db, err := utils.ExtractDB(c)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		//get user from db
		var usrModel models.User
		user, err := usrModel.GetOnePublic(db, uint(id))
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		//send
		c.JSON(http.StatusOK, user)
	}
}

// UserPost godoc
// @Security AuthToken
// @Tags users
// @Summary create a user (BY SUPER-ADMIN)
// @Description create a user (BY SUPER-ADMIN)
// @Param  user body swagmodel.InputUser true "create an user"
// @Accept json
// @Produce  json
// @Success 200 {object} swagmodel.Response
// @Router /users [post]
func (c *UserCtl) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		//decode
		var data models.User
		err := json.NewDecoder(c.Request.Body).Decode(&data)
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
		//store to db
		err = data.Post(db, &usr)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}

		//send
		utils.ResponseOK(c, "user saved")
	}
}

// UserPut godoc
// @Security AuthToken
// @Tags users
// @Summary update an user, only super-admin can update role
// @Description update an user
// @Param id path int true "ID of user"
// @Param  user body swagmodel.InputUser true "update an user"
// @Accept json
// @Produce  json
// @Success 200 {object} swagmodel.Response
// @Router /users/{id} [put]
func (c *UserCtl) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		//decode
		var data models.User
		err := json.NewDecoder(c.Request.Body).Decode(&data)
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

		//store to db
		data.Put(db, &usr, uint(id))
		//send
		utils.ResponseOK(c, "user edited")
	}
}

// UserDelete godoc
// @Security AuthToken
// @Tags users
// @Summary delete an user
// @Description delete an user
// @Param id path int true "ID of user"
// @Accept json
// @Produce  json
// @Success 200 {object} swagmodel.Response
// @Router /users/{id} [delete]
func (c *UserCtl) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		//get id from path
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
		db, err := utils.ExtractDB(c)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		//delete to db
		var userMdl models.User
		userMdl.Delete(db, &usr, uint(id))
		//send
		utils.ResponseOK(c, "user deleted")
	}
}
