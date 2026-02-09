package types

// CreatePasswordReq 创建密码请求
type CreatePasswordReq struct {
	Title       string `json:"title" validate:"required,max=200"`
	Description string `json:"description,optional"`
	Password    string `json:"password" validate:"required"`
}

// CreatePasswordResp 创建密码响应
type CreatePasswordResp struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

// GetPasswordListReq 获取密码列表请求
type GetPasswordListReq struct {
	Page     int64 `json:"page,optional,default=1"`
	PageSize int64 `json:"page_size,optional,default=20"`
}

// PasswordItem 密码项
type PasswordItem struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// GetPasswordListResp 获取密码列表响应
type GetPasswordListResp struct {
	List     []PasswordItem `json:"list"`
	Total    int64          `json:"total"`
	Page     int64          `json:"page"`
	PageSize int64          `json:"page_size"`
}

// GetPasswordDetailReq 获取密码详情请求
type GetPasswordDetailReq struct {
	ID int64 `json:"id" validate:"required"`
}

// GetPasswordDetailResp 获取密码详情响应
type GetPasswordDetailResp struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Password    string `json:"password"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// UpdatePasswordReq 更新密码请求
type UpdatePasswordReq struct {
	ID          int64  `json:"id" validate:"required"`
	Title       string `json:"title" validate:"required,max=200"`
	Description string `json:"description,optional"`
	Password    string `json:"password" validate:"required"`
}

// CommonResp 通用响应
type CommonResp struct {
	Message string `json:"message"`
}

// DeletePasswordReq 删除密码请求
type DeletePasswordReq struct {
	ID int64 `json:"id" validate:"required"`
}
