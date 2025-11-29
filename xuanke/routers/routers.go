package routers

import (
	"xuanke/api"

	"github.com/gin-gonic/gin"

	"xuanke/utils"
)

func RegisterRouters() {
	h := gin.Default()
	userGroup := h.Group("/user")
	adminGroup := h.Group("/admin")
	h.GET("/getallcourse", api.GetAllCourse)
	adminGroup.PUT("/addcourse", utils.JWTTokenAuth(), api.AddCourse)
	userGroup.PUT("/register", api.UserRegister)
	userGroup.POST("/login", api.LoginUser)
	userGroup.POST("/qiangke", utils.JWTTokenAuth(), api.Qiangke)
	userGroup.GET("/getchosecourse", utils.JWTTokenAuth(), api.Getchosecourse)
	userGroup.DELETE("/dropcourse", utils.JWTTokenAuth(), api.Dropcourse)
	err := h.Run(":8080")
	if err != nil {
		return
	}
}
