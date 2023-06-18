package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/vozalel/interview-crud-files/internal/controller/http/dto"
	"github.com/vozalel/interview-crud-files/internal/entity"
	"github.com/vozalel/interview-crud-files/pkg/logger"
)

func UserToContextSetter(c *gin.Context) {
	//
	// here you need your implementation of getting the user from the request and setting it to the context
	logger.Instance.Warn("test user embedded in context")
	userID := 1
	c.Set(dto.ContextKeyUser, entity.User{
		Name: "admin-test",
		ID:   &userID,
	})
}
