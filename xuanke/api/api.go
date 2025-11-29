package api

import (
	"errors"
	"fmt"
	"log"
	"xuanke/model"
	"xuanke/respond"
	"xuanke/service"

	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/gin-gonic/gin"
)

const StatusBadRequest int = 400

// id0
// id0
// id0
// id0
// id0
// id0
func LoginUser(c *gin.Context) {
	var user model.User

	err := c.BindJSON(&user)

	if err != nil {
		c.JSON(StatusBadRequest, respond.WrongParamType)
		return
	} ////////////////
	success, tokens, err := service.UserLogin(user)
	if !success {
		c.JSON(StatusBadRequest, respond.WrongUsernameOrPwd)
		return
	}
	c.JSON(200, gin.H{
		"message":       "登录成功",
		"access_token":  tokens.AccessToken,
		"refresh_token": tokens.RefreshToken,
		"user_id":       user.UserID,
	})
}

// id0
// id0
// id0
// id0
// id0
// id0
func AddCourse(c *gin.Context) {
	userIDValue, e := c.Get("user_id")
	fmt.Println(userIDValue)
	if !e {
		c.JSON(401, gin.H{"error": "用户信息不存在"})
		return
	}

	course := model.Course{}
	handlerID, ok := userIDValue.(int)
	if !ok {
		c.JSON(401, gin.H{"error": "用户ID格式错误"})
		return
	}
	fmt.Printf("AddCourse: 获取到 user_id = %d\n", handlerID)
	err := c.BindJSON(&course)
	if err != nil {
		c.JSON(StatusBadRequest, respond.WrongParamType)
		return
	}

	err = service.AddCourse(course, handlerID)
	if err != nil {
		switch {
		case errors.Is(err, respond.ErrUnauthorized), errors.Is(err, respond.Ok), errors.Is(err, respond.Ok):
			c.JSON(401, err)
			return
		default:
			c.JSON(500, respond.WrongParamType)
			return
		}
	}
	c.JSON(200, respond.Ok)
}

func GetAllCourse(c *gin.Context) {
	//(course []model.GetAllCourse, err error)

	course, err := service.GetAllCourse()
	if err != nil {
		c.JSON(StatusBadRequest, respond.WrongParamType)
	}
	c.JSON(200, course)
}
func UserRegister(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(StatusBadRequest, respond.WrongParamType)
		return
	}
	err = service.AddUser(user)
	if err != nil {
		switch {
		case errors.Is(err, respond.InvalidName), errors.Is(err, respond.Ok),
			errors.Is(err, respond.WrongParamType), errors.Is(err, respond.Ok): //如果是无效ID或者缺少参数的错误
			c.JSON(StatusBadRequest, err)
			return
		default:
			c.JSON(500, respond.Ok)
			return
		}
	}
	c.JSON(200, respond.Ok)
}

// /////////////////////////////////////
func Qiangke(c *gin.Context) {
	fmt.Printf("=== API.Qiangke 开始 ===\n")
	//userid := int(c.GetFloat64("user_id"))
	userIDValue, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "用户未登录"})
		return
	}

	userid, _ := userIDValue.(int)
	fmt.Println(userid)
	//
	//courseIDValue, _ := c.Get(" course_id")
	//
	//CourseID, _ := courseIDValue.(int)
	//fmt.Printf("抢课: user_id=%d, course_id=%d\n", userid, CourseID)
	//err := service.Qiangke(userid, CourseID)
	//if err != nil {
	//	c.JSON(500, respond.WrongParamType)
	//	return
	//}
	//c.JSON(consts.StatusOK, respond.Ok)
	//courseid := int(c.GetFloat64("course_id"))
	//
	//
	var request struct {
		CourseID int `json:"course_id" binding:"required"`
	}
	if err := c.BindJSON(&request); err != nil {
		fmt.Printf("❌ JSON 绑定错误: %v\n", err)
		c.JSON(400, respond.WrongParamType)
		return
	}
	fmt.Printf("抢课: user_id=%d, course_id=%d\n", userid, request.CourseID)
	err := service.Qiangke(userid, request.CourseID)
	if err != nil {
		c.JSON(500, respond.WrongParamType)
		return
	}
	c.JSON(consts.StatusOK, respond.Ok)
}

// ///////////////////////////////////////
func Getchosecourse(c *gin.Context) {
	userIDValue, _ := c.Get("user_id")
	userID := userIDValue.(int)
	course, err := service.Getchosecourse(userID)
	if err != nil {
		c.JSON(StatusBadRequest, respond.WrongParamType)
	}
	c.JSON(200, course)
	//var u model.User
	//err := c.ShouldBindJSON(&u)
	//if err != nil {
	//	c.JSON(StatusBadRequest, respond.WrongParamType)
	//}
	//course, err := service.Getchosecourse(u.UserID)
	//if err != nil {
	//	c.JSON(StatusBadRequest, respond.WrongParamType)
	//}
	//c.JSON(200, course)
}
func Dropcourse(c *gin.Context) {
	userIDValue, _ := c.Get("user_id")
	userID := userIDValue.(int)
	var request struct {
		CourseID int `json:"course_id" binding:"required"`
	}
	err := c.ShouldBindJSON(&request)
	fmt.Printf(" request.courseID:%d\n", request.CourseID)
	if err != nil {
		log.Println("请求参数错误：", err)
		return
	}
	err = service.DropCourse(userID, request.CourseID)
	if err != nil {
		c.JSON(500, respond.WrongParamType)
		return
	}

	c.JSON(200, respond.Ok)
}
