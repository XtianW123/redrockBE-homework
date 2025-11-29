package utils

import (
	"errors"
	"fmt"
	"net/http"
	"time"
	"xuanke/dao"
	"xuanke/model"
	"xuanke/respond"

	"github.com/gin-gonic/gin"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var accessKey = AccessKey
var (
	AccessKey  = []byte("access_secret_example_change_me")
	RefreshKey = []byte("refresh_secret_example_change_me")
	issuer     = "demo.jwt.singlefile"
	accessTTL  = 15 * time.Minute   // 访问令牌有效期
	refreshTTL = 7 * 24 * time.Hour // 刷新令牌有效期
)

// 自定义声明（访问/刷新可共用），并用 Type 区分 token 类型
type CustomClaims struct {
	UserID int64  `json:"uid"`
	Role   string `json:"role"`
	Type   string `json:"type"` // "access" or "refresh"
	jwt.RegisteredClaims
}

var JWTKey = []byte("your-secret-key-here")

func GenerateTokens(userID int) (string, string, error) {
	// 创建访问令牌
	fmt.Printf("🎯 === GenerateTokens 开始 ===\n")
	fmt.Printf("接收到的参数 userID = %d\n", userID)
	fmt.Printf("参数类型: %T\n", userID)
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    userID,                                  // 获取用户ID
		"exp":        time.Now().Add(40 * time.Minute).Unix(), // 设置访问令牌过期时间为 15 分钟
		"token_type": "access_token",                          // 令牌类型为访问令牌
	})

	// 使用密钥签名访问令牌
	accessTokenString, err := accessToken.SignedString(JWTKey)
	if err != nil {
		return "", "", err
	}

	// 创建刷新令牌
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    userID,                                    // 获取用户ID
		"exp":        time.Now().Add(7 * 24 * time.Hour).Unix(), // 设置刷新令牌过期时间为 7 天
		"token_type": "refresh_token",                           // 令牌类型为刷新令牌
	})

	// 使用密钥签名刷新令牌
	refreshTokenString, err := refreshToken.SignedString(JWTKey)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

func ValidateRefreshToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 检查签名方法是否为 HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, respond.InvalidTokenSingingMethod
		}
		// 返回用于验证的密钥
		return RefreshKey, nil
	})
	if err != nil {
		return nil, err
	}

	// 进一步检查载荷中 token_type 是否正确
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, respond.InvalidClaims
	}
	// 检查 token_type 是否是 refresh_token
	if claimType, ok := claims["token_type"].(string); !ok || claimType != "refresh_token" {
		return nil, respond.WrongTokenType
	}
	return token, nil
}
func HashPassword(pwd string) (string, error) {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPwd), nil
}

// CompareHashPwdAndPwd 用于比较哈希密码和密码是否匹配
func CompareHashPwdAndPwd(hashedPwd, pwd string) (bool, error) {
	fmt.Println("进入")
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) { //密码不匹配
		fmt.Println("q")
		return false, nil
	} else if err != nil { //其他错误
		fmt.Println("w")
		return false, err
	} else { //密码匹配
		fmt.Printf("ok")
		return true, nil
	}
}

// func Comparemima(mima string) (bool, error) {
//
// }
func CheckPermission(userID int) (string, error) {
	//user, err := dao.GetUserInfoByID(handlerID) //通过handlerID获取用户信息
	//if err != nil {
	//	return "", err
	//}
	//return user.Role, nil
	var user model.User
	result := dao.Db.Where("user_id = ?", userID).First(&user) // 注意：查询 user_id 字段
	if result.Error != nil {
		return "", result.Error
	}
	return user.Role, nil
}

func JWTTokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取 Authorization 字段
		tokenString := c.GetHeader("Authorization")
		fmt.Printf("=== JWT 中间件开始 ===\n")
		fmt.Printf("Authorization Header: %s\n", tokenString)
		fmt.Printf("Header 长度: %d\n", len(tokenString))
		if tokenString == "" { // 没有token
			c.JSON(http.StatusUnauthorized, respond.MissingToken)
			c.Abort() // 中断后续流程
			return
		}

		// 解析并验证 Token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 确保签名方法是我们支持的 HMAC
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				fmt.Printf(" 不支持的签名方法: %v\n", token.Method)
				return nil, respond.InvalidTokenSingingMethod
			}
			fmt.Println(" 签名方法正确")
			return JWTKey, nil
		})

		if err != nil { // token无效
			fmt.Printf(" Token 解析错误: %v\n", err)
			c.JSON(http.StatusUnauthorized, respond.InvalidToken)
			c.Abort() // 中断后续流程
			return
		}
		if !token.Valid {
			fmt.Println(" Token 无效")
			c.JSON(http.StatusUnauthorized, respond.InvalidToken)
			c.Abort()
			return
		}
		fmt.Printf(" Token 验证成功\n")

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			fmt.Printf("JWT Claims: %+v\n", claims)
			tokenType, ok := claims["token_type"].(string)
			if !ok {
				fmt.Println(" 无法获取 token_type")
				c.JSON(401, respond.InvalidClaims)
				c.Abort()
				return
			}

			fmt.Printf("Token 类型: %s\n", tokenType)

			if tokenType != "access_token" {
				fmt.Printf(" 错误的 token 类型: %s\n", tokenType)
				c.JSON(401, respond.WrongTokenType)
				c.Abort()
				return
			}

			// 正确获取 user_id
			userIDFloat, ok := claims["user_id"].(float64)
			if !ok {
				fmt.Printf(" 无法获取 user_id, 类型: %T, 值: %v\n", claims["user_id"], claims["user_id"])
				c.JSON(401, respond.InvalidClaims)
				c.Abort()
				return
			}

			userID := int(userIDFloat)
			fmt.Printf("JWT 中间件: 提取到 user_id = %d\n", userID)

			// 设置到 Gin 上下文
			c.Set("user_id", userID)
			c.Next()
		} else {
			c.JSON(401, respond.InvalidClaims)
			c.Abort()
		}
	}
}
