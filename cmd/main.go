package main

import (
	"github.com/eadydb/spider/pkg/auth"
	"github.com/eadydb/spider/pkg/repo"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	repo.RedisClient = repo.NewDefaultRedisConfig().
		NewRedisClient()

	v1 := r.Group("/v1")
	{
		v1.POST("/token", auth.Register)
		v1.GET("/parse", auth.Parse)
	}

	r.Run(":8080")
}
