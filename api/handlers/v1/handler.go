package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uzbekman2005/real_time_chatting/api/models"
	t "github.com/uzbekman2005/real_time_chatting/api/tokens"
	"github.com/uzbekman2005/real_time_chatting/config"
	"github.com/uzbekman2005/real_time_chatting/pkg/logger"
	"github.com/uzbekman2005/real_time_chatting/storage"
)

type handlerV1 struct {
	log        logger.Logger
	cfg        config.Config
	storage    storage.IStorage
	jwthandler t.JWTHandler
}

// HandlerV1Config ...
type HandlerV1Config struct {
	Logger     logger.Logger
	Cfg        config.Config
	Storage    storage.IStorage
	Jwthandler t.JWTHandler
}

// New ...
func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		log:        c.Logger,
		cfg:        c.Cfg,
		storage:    c.Storage,
		jwthandler: c.Jwthandler,
	}
}

// Default path
// @Summary 		Ping
// @Description 	Ping pong
// @Tags 			Ping
// @Produce         json
// @Success         200					  {object} 	models.SuccessInfo
// @Failure         500                   {object}  models.FailureInfo
// @Router          / [get]
func (h *handlerV1) AppIsRunning(c *gin.Context) {
	c.JSON(200, models.SuccessInfo{
		Message: "Server is running successfully",
	})
}

const (
	//ErrorCodeInvalidURL ...
	ErrorCodeInvalidURL = "INVALID_URL"
	//ErrorCodeInvalidJSON ...
	ErrorCodeInvalidJSON = "INVALID_JSON"
	//ErrorCodeInvalidArgument ...
	ErrorCodeInvalidArgument = "INVALID_ARGUMENT"
	//ErrorCodeInvalidParams ...
	ErrorCodeInvalidParams = "INVALID_PARAMS"
	//ErrorCodeInternal ...
	ErrorCodeInternal = "INTERNAL"
	//ErrorCodeUnauthorized ...
	ErrorCodeUnauthorized = "UNAUTHORIZED"
	//ErrorCodeAlreadyExists ...
	ErrorCodeAlreadyExists = "ALREADY_EXISTS"
	//ErrorCodeNotFound ...
	ErrorCodeNotFound = "NOT_FOUND"
	//ErrorCodeInvalidCode ...
	ErrorCodeInvalidCode = "INVALID_CODE"
	//ErrorBadRequest ...
	ErrorBadRequest = "BAD_REQUEST"
	//ErrorCodeForbidden ...
	ErrorCodeForbidden = "FORBIDDEN"
	//ErrorCodeNotApproved ...
	ErrorCodeNotApproved = "NOT_APPROVED"
	// ErrorUpgradeRequired ...
	ErrorUpgradeRequired = "UPGRADE_REQUIRED"
	// ErrorUserBlocked ...
	ErrorUserBlocked = "USER_IS_BLOCKED_BY_ADMIN"
	// ErrorInvalidCredentials ...
	ErrorInvalidCredentials = "INVALID_CREDENTIALS"
)

func HandleInternalWithMessage(c *gin.Context, l logger.Logger, err error, message string, args ...interface{}) bool {
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.ServerError{
				Status:  ErrorCodeInternal,
				Message: "Internal Server Error",
			},
		})
		l.Error(message, logger.Error(err), logger.Any("req", args))
		return true
	}

	return false
}

func HandleBadRequestErrWithMessage(c *gin.Context, l logger.Logger, err error, message string, args ...interface{}) bool {
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.ServerError{
				Status:  ErrorCodeInvalidJSON,
				Message: "Incorrect data supplied",
			},
		})
		l.Error(message, logger.Error(err), logger.Any("req", args))
		return true
	}
	return false
}
