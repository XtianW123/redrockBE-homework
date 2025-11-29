package main

//这个也离谱,get不报错但出错,post报错但不出错
import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type STUDENT struct {
	ID    uint `gorm:"primarykey"`
	Name  string
	Likes []ZHONGJIAN `gorm:"foreignKey:STUDENTID"`
}
type COURSE struct {
	ID    uint `gorm:"primarykey"`
	Name  string
	Likes []ZHONGJIAN `gorm:"foreignKey:COURSEID"`
}
type ZHONGJIAN struct {
	STUDENTID uint    `gorm:"primarykey"`
	COURSEID  uint    `gorm:"primarykey"`
	STUDENT   STUDENT `gorm:"foreignkey:STUDENTID;references:ID;"`
	COURSE    COURSE  `gorm:"foreignkey:COURSEID;references:ID;"`
}

func main() {
	r := gin.Default()
	r.POST("/xuanke", x)
	r.Run()
}
func x(c *gin.Context) {
	db, _ := gorm.Open(mysql.Open("root:mima@tcp(localhost:3306)/school?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	db.AutoMigrate(&STUDENT{}, &COURSE{}, &ZHONGJIAN{})
	db.Create(&STUDENT{ID: 1, Name: "孙"})
	db.Create(&STUDENT{ID: 2, Name: "王"})
	db.Create(&COURSE{ID: 1, Name: "数学"})
	db.Create(&COURSE{ID: 2, Name: "c"})
	db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&STUDENT{ID: 3, Name: "于"}).Error; err != nil {
			return err
		}
		if err := tx.Create(&COURSE{ID: 3, Name: "思政"}).Error; err != nil {
			return err
		}

		var count int64
		tx.Model(&ZHONGJIAN{}).Where("COURSEID = ?", 1).Count(&count)
		if count < 50 {
			guanxi := []ZHONGJIAN{
				{STUDENTID: 1, COURSEID: 1}, {STUDENTID: 3, COURSEID: 2},
			}
			if err := tx.Create(&guanxi).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
