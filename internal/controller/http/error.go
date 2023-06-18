package http

import (
	"github.com/vozalel/interview-crud-files/pkg/custom_error"
	"github.com/vozalel/interview-crud-files/pkg/logger"

	"github.com/gin-gonic/gin"
)

type properErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
} // @name properErrorResponse

func respondWithCustomError(ctx *gin.Context, err *custom_error.CustomError) {
	if logger.GetLogLevel() == "debug" {
		ctx.AbortWithStatusJSON(err.Code, properErrorResponse{Error: err.Error(), Message: err.Message})
	} else {
		ctx.AbortWithStatusJSON(err.Code, properErrorResponse{Message: err.Message})
	}
	ctx.Error(err)
}
