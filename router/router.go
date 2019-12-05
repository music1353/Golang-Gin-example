package router

import (
    "github.com/gin-gonic/gin";
    "my-gin/router/apis/v1"
)

func InitRouter() *gin.Engine {
    r := gin.Default()

    apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/", v1.IndexApi)
        apiv1.POST("/person", v1.AddPersonApi)
        apiv1.GET("/person", v1.GetPersonApi)
	}

    return r
}