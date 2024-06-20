package presenter

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/hutoboshi/go_rest_api_gin/api/internal/controller/system"
	"github.com/hutoboshi/go_rest_api_gin/api/internal/controller/user"
)

const latest = "/v1"

type Server struct{}

func (s *Server) Run(ctx context.Context) error {
	r := gin.Default()
	v1 := r.Group(latest)

	// 死活監視用
	{
		systemHandler := system.NewSystemHandler()
		v1.GET("/health", systemHandler.Health)
	}

	// ユーザー管理機能
	{
		userHandler := user.NewUserHandler()
		v1.GET("", userHandler.GetUsers)
		v1.GET("/:id", userHandler.GetUserById)
		v1.POST("", userHandler.EditUser)
		v1.DELETE("/:id", userHandler.DeleteUser)
	}

	err := r.Run()
	if err != nil {
		return err
	}

	return nil
}

func NewServer() *Server {
	return &Server{}
}
