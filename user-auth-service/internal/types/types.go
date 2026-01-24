package types

// RegisterReq 注册请求
type RegisterReq struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	Password string `json:"password" validate:"required,min=6,max=20"`
	Email    string `json:"email" validate:"required,email"`
	Mobile   string `json:"mobile,optional"`
}

// RegisterResp 注册响应
type RegisterResp struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

// LoginReq 登录请求
type LoginReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// LoginResp 登录响应
type LoginResp struct {
	ID          int64  `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
	ExpireTime  int64  `json:"expire_time"`
}

// UserInfoReq 用户信息请求
type UserInfoReq struct {
	UserID int64 `json:"user_id,optional"`
}

// UserInfoResp 用户信息响应
type UserInfoResp struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
}

// ChangePasswordReq 修改密码请求
type ChangePasswordReq struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=6,max=20"`
}

// CommonResp 通用响应
type CommonResp struct {
	Message string `json:"message"`
}
