package types

// CreatePasswordReq represents a create password request
type CreatePasswordReq struct {
	Title       string `json:"title" validate:"required,max=200"`
	Description string `json:"description,optional"`
	Password    string `json:"password" validate:"required"`
}

// CreatePasswordResp represents a create password response
type CreatePasswordResp struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

// GetPasswordListReq represents a get password list request
type GetPasswordListReq struct {
	Page     int64 `json:"page,optional,default=1"`
	PageSize int64 `json:"page_size,optional,default=20"`
}

// PasswordItem represents a password item
type PasswordItem struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// GetPasswordListResp represents a get password list response
type GetPasswordListResp struct {
	List     []PasswordItem `json:"list"`
	Total    int64          `json:"total"`
	Page     int64          `json:"page"`
	PageSize int64          `json:"page_size"`
}

// GetPasswordDetailReq represents a get password detail request
type GetPasswordDetailReq struct {
	ID int64 `json:"id" validate:"required"`
}

// GetPasswordDetailResp represents a get password detail response
type GetPasswordDetailResp struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Password    string `json:"password"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// UpdatePasswordReq represents an update password request
type UpdatePasswordReq struct {
	ID          int64  `json:"id" validate:"required"`
	Title       string `json:"title" validate:"required,max=200"`
	Description string `json:"description,optional"`
	Password    string `json:"password" validate:"required"`
}

// CommonResp represents a common response
type CommonResp struct {
	Message string `json:"message"`
}

// DeletePasswordReq represents a delete password request
type DeletePasswordReq struct {
	ID int64 `json:"id" validate:"required"`
}
