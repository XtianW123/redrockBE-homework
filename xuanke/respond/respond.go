package respond

type Response struct { //响应结构体
	Status string `json:"status"`
	Info   string `json:"info"`
}

func (r Response) Error() string { // 实现 error 接口
	return r.Info
}

var (
	Weishuru = Response{
		Status: "1111",
		Info:   "Weishuru",
	}
	Ok = Response{ //正常
		Status: "10000",
		Info:   "success",
	}

	WrongName = Response{ //用户名错误
		Status: "40001",
		Info:   "wrong username",
	}
	WrongParamType = Response{ //参数错误
		Status: "40005",
		Info:   "wrong param type",
	}
	WrongPwd = Response{ //密码错误
		Status: "40002",
		Info:   "wrong password",
	}

	InvalidName = Response{ //用户名无效
		Status: "40003",
		Info:   "the username already exists",
	}

	WrongUsernameOrPwd = Response{ //用户名或密码错误
		Status: "40007",
		Info:   "wrong username or password",
	}

	MissingToken = Response{ //缺少token
		Status: "40009",
		Info:   "missing token",
	}

	InvalidTokenSingingMethod = Response{ //jwt token签名方法无效
		Status: "40010",
		Info:   "invalid signing method",
	}

	InvalidToken = Response{ //无效token
		Status: "40011",
		Info:   "invalid token",
	}

	InvalidClaims = Response{ //无效声明
		Status: "40012",
		Info:   "invalid claims",
	}

	WrongUserID = Response{ //用户ID错误
		Status: "40013",
		Info:   "wrong userid",
	}

	ErrUnauthorized = Response{ //未授权，没有权限
		Status: "40014",
		Info:   "unauthorized",
	}

	ErrProductNotExists = Response{ //商品不存在
		Status: "40017",
		Info:   "product not exists",
	}

	CantFindProduct = Response{ //找不到商品
		Status: "40018",
		Info:   "can't find product",
	}

	EmptyProductList = Response{ //商品列表为空
		Status: "40019",
		Info:   "product list is empty",
	}
	InvalidRefreshToken = Response{ //刷新令牌无效
		Status: "40020",
		Info:   "invalid refresh token",
	}

	ErrQuantityTooLarge = Response{ //数量太大
		Status: "40022",
		Info:   "quantity too large",
	}

	WrongTokenType = Response{ //无效令牌类型
		Status: "40033",
		Info:   "wrong token type",
	}
)
