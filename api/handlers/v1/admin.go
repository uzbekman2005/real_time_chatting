package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uzbekman2005/real_time_chatting/api/models"
	"github.com/uzbekman2005/real_time_chatting/pkg/etc"
	"github.com/uzbekman2005/real_time_chatting/pkg/logger"
)

// @Summary 		Login admin
// @Description 	Through this api admin can login
// @Tags 			Admin
// @Accept 			json
// @Produce         json
// @Param        	admin_name      	  path  		  string    true 	"admin_name"
// @Param        	password        	  path  		  string    true 	"password"
// @Success         200					  {object} 	models.AdminRes
// @Failure         500                   {object}  models.ResponseError
// @Failure         400                   {object}  models.ResponseError
// @Failure         409                   {object}  models.ResponseError
// @Router          /admin/login/{admin_name}/{password} [get]
func (h *handlerV1) AdminLogin(c *gin.Context) {
	var (
		adminName = c.Param("admin_name")
		password  = c.Param("password")
	)
	res, err := h.storage.Chat().GetAdminByName(adminName)
	if HandleBadRequestErrWithMessage(c, h.log, err, "storage.Chat().GetAdminByName(adminName)") {
		return
	}
	if !etc.CheckPasswordHash(password, res.Password) {
		c.JSON(http.StatusBadRequest, models.FailureInfo{
			Message: "Your password is incorrect",
		})
		return
	}

	h.jwthandler.Sub = res.Name
	h.jwthandler.Role = res.Role
	h.jwthandler.Aud = []string{"universities-frontend"}
	h.jwthandler.SigninKey = h.cfg.SignInKey
	h.jwthandler.Log = h.log
	tokens, err := h.jwthandler.GenerateAuthJWT()
	if HandleInternalWithMessage(c, h.log, err, "Generate JWT Token") {
		return
	}
	accessToken := tokens[0]

	res.AccessToken = accessToken

	c.JSON(http.StatusOK, res)
}

// @Summary 		Create Admin
// @Description 	Through this api admin can be created only super admin can create new admin.
// @Tags 			Admin
// @Security        BearerAuth
// @Accept 			json
// @Produce         json
// @Param           Admin        body  	  models.CreateAdminReq true "Admin"
// @Success         200					  {object} 	models.AdminRes
// @Failure         500                   {object}  models.FailureInfo
// @Failure         400                   {object}  models.FailureInfo
// @Failure         409                   {object}  models.FailureInfo
// @Router          /admin [post]
func (h *handlerV1) CreateAdmin(c *gin.Context) {
	body := &models.CreateAdminReq{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, &models.FailureInfo{Message: "Please check your data correctness"})
		h.log.Error("Error while getting admin info from request", logger.Error(err))
		return
	}

	isunique, err := h.storage.Chat().IsUniqueAdminName(body.Name)
	if HandleBadRequestErrWithMessage(c, h.log, err, "storage.Chat().IsUniqueAdminName(body.Name)") {
		return
	}

	if !isunique {
		c.JSON(http.StatusConflict, &models.FailureInfo{
			Message: "This admin name was already used",
		})
	}

	body.Password, err = etc.HashPassword(body.Password)
	if HandleBadRequestErrWithMessage(c, h.log, err, "etc.HashPassword(body.Password)") {
		return
	}

	res, err := h.storage.Chat().CreateAdmin(body)
	if HandleInternalWithMessage(c, h.log, err, "storage.Chat().CreateAdmin(body)") {
		return
	}

	h.jwthandler.Sub = body.Name
	h.jwthandler.Role = body.Role
	h.jwthandler.Aud = []string{"universities-frontend"}
	h.jwthandler.SigninKey = h.cfg.SignInKey
	h.jwthandler.Log = h.log
	tokens, err := h.jwthandler.GenerateAuthJWT()
	accessToken := tokens[0]
	res.AccessToken = accessToken
	if HandleInternalWithMessage(c, h.log, err, "h.jwthandler.GenerateAuthJWT()") {
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary 		Get all admins profile
// @Description 	Through this api all admins can be taken
// @Tags 			Admin
// @Security        BearerAuth
// @Produce         json
// @Success         200					  {object} 	[]models.AdminRes
// @Failure         500                   {object}  models.ResponseError
// @Failure         400                   {object}  models.ResponseError
// @Failure         409                   {object}  models.FailureInfo
// @Router          /admin/all [get]
func (h *handlerV1) GetAllAdmins(c *gin.Context) {
	res, err := h.storage.Chat().GetAllAdmins()
	if HandleInternalWithMessage(c, h.log, err, "h.storage.Chat().GetAllAdmins()") {
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary 		Admin Admin
// @Description 	Through this api admin can be updated.
// @Tags 			Admin
// @Security        BearerAuth
// @Accept 			json
// @Produce         json
// @Param           Admin        body  	  models.UpdateAdminInfo true "Admin"
// @Success         200					  {object} 	models.SuccessInfo
// @Failure         500                   {object}  models.ResponseError
// @Failure         400                   {object}  models.ResponseError
// @Failure         409                   {object}  models.FailureInfo
// @Router          /admin [put]
func (h *handlerV1) UpdateAdmin(c *gin.Context) {
	body := &models.UpdateAdminInfo{}
	err := c.ShouldBindJSON(&body)
	if HandleBadRequestErrWithMessage(c, h.log, err, "c.ShouldBindJSON(&body)") {
		return
	}

	isunique, err := h.storage.Chat().IsUniqueAdminName(body.NewName)
	if HandleBadRequestErrWithMessage(c, h.log, err, "c.ShouldBindJSON(&body)") {
		return
	}
	if !isunique {
		c.JSON(http.StatusConflict, &models.FailureInfo{
			Message: "This admin name was already used",
		})
	}

	body.NewPassword, err = etc.HashPassword(body.NewPassword)
	if HandleInternalWithMessage(c, h.log, err, "HashPassword(body.NewPassword)") {
		return
	}

	admin, err := h.storage.Chat().GetAdmin(body.Id)
	if HandleBadRequestErrWithMessage(c, h.log, err, "storage.Chat().GetAdmin(body.Id)") {
		return
	}
	if !etc.CheckPasswordHash(body.OldPassword, admin.Password) {
		c.JSON(http.StatusBadRequest, &models.FailureInfo{Message: "Your old password is incorrect"})
		return
	}

	err = h.storage.Chat().UpdateAdmin(body)
	if HandleInternalWithMessage(c, h.log, err, "h.storage.Chat().UpdateAdmin(body)") {
		return
	}

	c.JSON(http.StatusOK, models.SuccessInfo{
		Message: "Admin credientials are succesfully updated",
	})
}

// @Summary 		Delete Admin
// @Description 	Through this api admin can be deleted.
// @Tags 			Admin
// @Security        BearerAuth
// @Accept 			json
// @Produce         json
// @Param           Admin        body  	  models.DeleteAdminReq true "Admin"
// @Success         200					  {object} 	models.SuccessInfo
// @Failure         500                   {object}  models.ResponseError
// @Failure         400                   {object}  models.ResponseError
// @Failure         409                   {object}  models.FailureInfo
// @Router          /admin [delete]
func (h *handlerV1) DeleteAdmin(c *gin.Context) {
	body := &models.DeleteAdminReq{}
	err := c.ShouldBindJSON(&body)
	if HandleBadRequestErrWithMessage(c, h.log, err, "c.ShouldBindJSON(&body)") {
		return
	}

	err = h.storage.Chat().DeleteAdmin(body.Id)
	if HandleInternalWithMessage(c, h.log, err, "h.storage.Chat().DeleteAdmin(body.Id)") {
		return
	}

	c.JSON(http.StatusOK, models.SuccessInfo{
		Message: "Admin succesfully deleted",
	})
}
