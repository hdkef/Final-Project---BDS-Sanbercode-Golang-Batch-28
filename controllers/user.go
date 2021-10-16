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
		//auth for role super-admin
		usr, err := utils.ExtractUserContext(c)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		if usr.Role != "super-admin" {
			utils.ResponseError(c, http.StatusInternalServerError, "not super admin")
			return
		}
		//get user from db

		//send
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
		//auth for id == usr.ID or role == super-admin
		usr, err := utils.ExtractUserContext(c)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		if usr.ID != uint(id) || usr.Role != "super-admin" {
			utils.ResponseError(c, http.StatusInternalServerError, "unauthorized access")
			return
		}
		//get user from db

		//send
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
		// id, err := strconv.Atoi(c.Params.ByName("id"))
		// if err != nil {
		// 	utils.ResponseError(c, http.StatusInternalServerError, err.Error())
		// 	return
		// }
		//get user from db

		//set sensitive information = null

		//send
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
		//auth role == super-admin
		usr, err := utils.ExtractUserContext(c)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		if usr.Role != "super-admin" {
			utils.ResponseError(c, http.StatusInternalServerError, "not super-admin")
			return
		}
		//decode
		var data models.User
		err = json.NewDecoder(c.Request.Body).Decode(&data)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		//store to db

		//send
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
		//get id from path
		id, err := strconv.Atoi(c.Params.ByName("id"))
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		//auth for id == usr.ID or role == super-admin
		usr, err := utils.ExtractUserContext(c)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		if usr.ID != uint(id) || usr.Role != "super-admin" {
			utils.ResponseError(c, http.StatusInternalServerError, "unauthorized access")
			return
		}
		//decode
		var data models.User
		err = json.NewDecoder(c.Request.Body).Decode(&data)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		//store to db
		//send
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
		//auth for id == usr.ID or role == super-admin
		usr, err := utils.ExtractUserContext(c)
		if err != nil {
			utils.ResponseError(c, http.StatusInternalServerError, err.Error())
			return
		}
		if usr.ID != uint(id) || usr.Role != "super-admin" {
			utils.ResponseError(c, http.StatusInternalServerError, "unauthorized access")
			return
		}
		//delete to db

		//send
	}
}
