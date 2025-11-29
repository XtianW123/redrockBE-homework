package dao

import (
	"fmt"
	"xuanke/model"
	"xuanke/respond"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error
var Db *gorm.DB

func ConnectDB() error {
	dsn := "root:mima@tcp(127.0.0.1:3306)/xuanke?charset=utf8mb4&parseTime=True&loc=Local&timeout=30s"
	//第一部分：连接数据库，并检测其连接正常性

	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}) //链接数据库
	if err != nil {
		return err
	}
	Db.AutoMigrate(&model.User{}, &model.Course{}, &model.ZHONGJIAN{})
	return err
}
func GetUserByUsername(username string) (model.User, error) {
	var user model.User
	result := Db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}
func AddUser(user model.User) (err error) {
	//query := "INSERT INTO user (username, password, role, user_id) VALUES (?, ?, ?, ?)"
	//
	//_ = Db.Exec(query, user.Username, user.Password, user.Role, user.UserID)
	//return err
	result := Db.Create(&user)
	return result.Error
}
func AddCourse(course model.Course) (err error) {
	//query := "INSERT INTO course (course_id, course_name) VALUES (?, ?)"
	//_ = Db.Exec(query, course.CourseID, course.CourseID)
	//return err
	result := Db.Create(&course)
	return result.Error
}
func GetCourse() ([]model.Course, error) {
	var courses []model.Course
	Db.Find(&courses)
	return courses, nil
}
func Getchosecourses(uid int) ([]model.Course, error) {
	var courses []model.Course
	var courseIDs []int
	Db.Where("user_id = ?", uid).Table("zhongjians").Select("course_id").Find(&courseIDs)
	Db.Table("courses").Where("course_id IN (?)", courseIDs).Find(&courses)
	return courses, nil
}
func GetCourseByID(cid []int) ([]model.Course, error) {
	var courses []model.Course
	Db.Table("courses").Where("course_id = ?", cid).Find(&courses)
	return courses, nil
}
func GetUserHashedPassword(username string) (string, error) {
	var user model.User
	//result := Db.Where("username = ?", username).First(&password)
	result := Db.Table("users").Where("username = ?", username).Select("password").First(&user)
	if result.Error != nil {
		return "", result.Error
	}
	fmt.Println(user.Password)
	return user.Password, nil
}

func GetUserID(username1 string) (int, error) { //获取用户id
	fmt.Printf("🔍 === dao.GetUserID 开始 ===\n")
	fmt.Printf("查询用户名: %s\n", username1)
	fmt.Printf("jinrugetid")
	//var id int
	var user model.User
	//Db.Where("username = ?", username).First(&id)
	//Db.Table("users").Where("username = ?", username1).Select("user_id").First(&id)
	result := Db.Where("username = ?", username1).First(&user)
	if result.Error != nil {
		fmt.Printf(" 查询用户信息失败: %v\n", result.Error)
		return 0, result.Error
	}
	fmt.Printf("✅ 查询到的用户信息: ID=%d, UserID=%d, Username=%s, Role=%s\n",
		user.UserID, user.Username, user.Role)

	fmt.Printf("🔍 === dao.GetUserID 结束，返回: %d ===\n", user.UserID)
	//return id, nil
	return user.UserID, nil
	//////////////////////

	//var id int
	//query := "SELECT id FROM users WHERE username=?"
	//rows, err := Query(query, username)
	//if err != nil {
	//	return 0, err
	//}
	//if rows.Next() { //如果有这个用户
	//	err = rows.Scan(&id) //将用户id赋值给id
	//	if err != nil {
	//		return 0, err
	//	}
	//	return id, nil
	//}
	//return 0, respond.WrongName //找不到用户
}
func GetUserInfoByID(id int) (model.User, error) {
	var user model.User
	Db.Where("id=?", id).First(&user)

	return user, respond.Ok

}

func Qiangke(userID, courseID int) error {
	//fmt.Printf("DAO.Qiangke: 创建选课记录 userID=%d, courseID=%d\n", userID, courseID)
	//Db.Transaction(func(tx *gorm.DB) error {
	//	var mu sync.Mutex
	//
	//	mu.Lock()
	//	tx.Create(&model.ZHONGJIAN{ID: userID, CourseID: courseID})
	//	mu.Unlock()
	//	return nil
	//})
	//return nil
	return Db.Transaction(func(tx *gorm.DB) error {
		// 不需要手动加锁，数据库事务会自动处理
		tx.Create(&model.ZHONGJIAN{UserID: userID, CourseID: courseID})
		fmt.Printf("选课成功")
		return nil
	})
}
func Dropcourse(userID, courseID int) error {
	Db.Where("user_id=?", userID).Where("course_id=?", courseID).Delete(&model.ZHONGJIAN{})
	return nil
}
