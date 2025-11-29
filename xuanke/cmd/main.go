package main

import (
	"fmt"
	"xuanke/dao"
	"xuanke/routers"
)

//	func CompareHashPwdAndPwd(hashedPwd, pwd string) (bool, error) {
//		fmt.Printf("=== 密码比较开始 ===\n")
//		fmt.Printf("哈希密码: %s\n", hashedPwd)
//		fmt.Printf("输入密码: %s\n", pwd)
//		fmt.Printf("哈希密码长度: %d\n", len(hashedPwd))
//		fmt.Printf("输入密码长度: %d\n", len(pwd))
//
//		// 直接调用并打印错误
//		err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd))
//		fmt.Printf("bcrypt.CompareHashAndPassword 返回的错误: %v\n", err)
//		fmt.Printf("错误是否为nil: %t\n", err == nil)
//
//		if err != nil {
//			fmt.Printf("错误类型: %T\n", err)
//			if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
//				fmt.Printf("明确是密码不匹配错误\n")
//				return false, nil
//			}
//			fmt.Printf("其他错误，直接返回\n")
//			return false, err
//		}
//
//		fmt.Printf("密码匹配成功！\n")
//		return true, nil
//	}
func main() {
	err := dao.ConnectDB()
	if err != nil {
		fmt.Println(err)
	}
	//hashed1 := "$2a$10$eT/CrW.ZOytgK9aOaTPYdOkegW1XmmlmkBvoQCTpxXAjV6pQcyRWq"
	//password1 := "123456" // 假设这是原始密码
	//CompareHashPwdAndPwd(hashed1, password1)

	routers.RegisterRouters() //注册路由并启动服务

}
