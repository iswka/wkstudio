package types

// RegisterReq represents a registration request
type RegisterReq struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	Password string `json:"password" validate:"required,min=6,max=20"`
	Email    string `json:"email" validate:"required,email"`
	Mobile   string `json:"mobile,optional"`
}

// RegisterResp represents a registration response
type RegisterResp struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

// LoginReq represents a login request
type LoginReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// LoginResp represents a login response
type LoginResp struct {
	ID          int64  `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
	ExpireTime  int64  `json:"expire_time"`
}

// UserInfoReq represents a user info request
type UserInfoReq struct {
	UserID int64 `json:"user_id,optional"`
}

// UserInfoResp represents a user info response
type UserInfoResp struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
}

// ChangePasswordReq represents a change password request
type ChangePasswordReq struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=6,max=20"`
}

// CommonResp represents a common response
type CommonResp struct {
	Message string `json:"message"`
}
