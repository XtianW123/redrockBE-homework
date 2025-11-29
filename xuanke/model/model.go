package model

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AddUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	UserID   int    `json:"user_id"`
}

//	type LoginUser struct {
//		Username string `json:"username"`
//		Password string `json:"password"`
//	}
type JWTKey struct {
	Key string `json:"key"`
}

//	type GetAllCourse struct {
//		CourseName string `json:"course_name"`
//		CourseID   int    `json:"course_id"`
//		Rongliang  int    `json:"rongliang"`
//	}
type Course struct {
	CourseID   int    `gorm:"primaryKey" json:"course_id"`
	CourseName string `json:"course_name"`
	Rongliang  int    `json:"rongliang"` //CourseID   int    `json:"course_id"`
	//ID         int    `gorm:"primaryKey" json:"id"`
	//CourseID   int    `json:"course_id"`
	//Rongliang  int    `json:"rongliang"`
	//	UserSelections []ZHONGJIAN `gorm:"foreignKey:CourseID"`
}
type User struct {
	UserID   int    `gorm:"primaryKey" json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	//ID       int    `gorm:"primaryKey" json:"id"`
	//UserID   int    `json:"user_id"`
	//Username string `json:"username"`
	//Password string `json:"password"`
	//Role     string `json:"role"`
	//UserSelections []ZHONGJIAN `gorm:"foreignKey:UserID"`
}
type ZHONGJIAN struct {
	ID       int `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID   int `gorm:"index" json:"user_id"`
	CourseID int `gorm:"index" json:"course_id"`
	//USER     User   `gorm:"foreignkey:UserID;references:UserID;"`
	//COURSE   Course `gorm:"foreignkey:CourseID;references: CourseID;"`
}
