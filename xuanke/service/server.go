package service

import (
	"fmt"
	"xuanke/dao"
	"xuanke/model"
	"xuanke/respond"
	"xuanke/utils"
)

func AddUser(user model.User) error {
	if user.Username == "" || user.Password == "" {
		return respond.Weishuru
	}
	hashedPwd, err := utils.HashPassword(user.Password)
	user.Password = hashedPwd //将user的密码字段改为加密后的密码
	err = dao.AddUser(user)   //调用dao层的方法
	if err != nil {
		return err
	}
	return nil

}

func UserLogin(user model.User) (bool, model.Tokens, error) {
	var tokens model.Tokens
	fmt.Printf("=== Service.UserLogin 开始 ===\n")
	fmt.Printf("输入的用户名: %s\n", user.Username)
	fmt.Printf("输入的 user.UserID: %d\n", user.UserID)           // 调试输入的 UserID
	hashedPwd, err := dao.GetUserHashedPassword(user.Username) //调用dao层的方法
	if err != nil {
		return false, model.Tokens{}, err
	}
	result, err := utils.CompareHashPwdAndPwd(hashedPwd, user.Password)
	fmt.Println("结束CompareHashPwdAndPwd")
	//比较密码是否匹配
	if err != nil { //其他错误
		return false, tokens, err
	} else if !result { //密码不匹配
		return false, model.Tokens{}, respond.WrongPwd
	}
	fmt.Printf("准备getid")
	id, err := dao.GetUserID(user.Username) //获取用户id
	fmt.Println("出getid")
	if err != nil {
		fmt.Printf("err1")
		return false, model.Tokens{}, err
	}
	fmt.Printf("🔐 Service层: 准备调用 utils.GenerateTokens(%d)\n", id)
	tokens.AccessToken, tokens.RefreshToken, err = utils.GenerateTokens(id) //生成jwt key
	if err != nil {                                                         //其他错误
		fmt.Printf("err2")
		return false, model.Tokens{}, err
	}
	fmt.Println("service ok")
	return true, tokens, nil
	//var tokens model.Tokens
	//fmt.Printf("尝试登录用户: %s\n", user.Username)
	//hashedPwd, err := dao.GetUserHashedPassword(user.Username) //调用dao层的方法
	//fmt.Println("123")
	//if hashedPwd != "" {
	//	fmt.Printf(hashedPwd)
	//}
	//if err != nil {
	//	return false, model.Tokens{}, err
	//}
	//fmt.Println("456")
	//result, err := utils.CompareHashPwdAndPwd(hashedPwd, user.Password) //比较密码是否匹配
	//if err != nil {                                                     //其他错误
	//	return false, tokens, err
	//} else if !result { //密码不匹配
	//	return false, model.Tokens{}, respond.WrongPwd
	//}
	//id, err := dao.GetUserID(user.Username) //获取用户id
	//if err != nil {
	//	return false, model.Tokens{}, err
	//}
	//tokens.AccessToken, tokens.RefreshToken, err = utils.GenerateTokens(id) //生成jwt key
	//if err != nil {                                                         //其他错误
	//	return false, model.Tokens{}, err
	//}
	//return true, tokens, nil

	//dbUser, err := dao.GetUserByUsername(user.Username)
	//if err != nil {
	//	return false, model.Tokens{}, respond.WrongName
	//}
	//
	//// 2. 比较密码
	//result, err := utils.CompareHashPwdAndPwd(dbUser.Password, user.Password)
	//if err != nil {
	//	return false, tokens, err
	//} else if !result {
	//	return false, model.Tokens{}, respond.WrongPwd
	//}
	//
	//// 3. 生成Token
	//tokens.AccessToken, tokens.RefreshToken, err = utils.GenerateTokens(dbUser.ID)
	//if err != nil {
	//	return false, model.Tokens{}, err
	//}
	//
	//return true, tokens, nil
}
func AddCourse(course model.Course, handlerID int) error {
	fmt.Printf("Service: 收到课程添加请求，操作人ID = %d\n", handlerID)
	role, err := utils.CheckPermission(handlerID) //检查用户权限
	if err != nil {                               //如果出错
		return err
	}
	if role != "admin" { //如果不是商家
		return respond.ErrUnauthorized //返回错误
	}
	fmt.Printf("权限验证通过: 用户角色是 %s\n", role)
	return dao.AddCourse(course)
}
func GetAllCourse() (course []model.Course, err error) {
	course, err = dao.GetCourse()

	return
}
func Getchosecourse(uid int) (course []model.Course, err error) {

	course, err = dao.Getchosecourses(uid)
	return
}
func Qiangke(userid, courseid int) error {
	fmt.Printf("Service.Qiangke: userID=%d, courseID=%d\n", userid, courseid)
	var course model.Course
	if err := dao.Db.Where("course_id = ?", courseid).First(&course).Error; err != nil {
		return fmt.Errorf("课程不存在: %v", err)
	}
	var existing model.ZHONGJIAN
	if err := dao.Db.Where("user_id = ? AND course_id = ?", userid, courseid).First(&existing).Error; err == nil {
		return fmt.Errorf("已经选过这门课")
	}
	return dao.Qiangke(userid, courseid)

}
func DropCourse(userID, courseID int) (err error) {
	fmt.Printf("Service.DropCourse: userID=%d\n", userID)
	var a model.User
	var b model.Course
	if err := dao.Db.Table("users").Where("user_id = ?", userID).First(&a).Error; err != nil {
		return err
	}
	if err := dao.Db.Table("courses").Where("course_id = ?", courseID).First(&b).Error; err != nil {
		return err
	}
	err = dao.Dropcourse(userID, courseID)
	if err != nil {
		return err
	}
	return nil
}
