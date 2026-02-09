package types

// CreateWebsiteReq represents a create website request
type CreateWebsiteReq struct {
	Title       string `json:"title" validate:"required,max=200"`
	Icon        string `json:"icon,optional"`
	Description string `json:"description,optional"`
	URL         string `json:"url" validate:"required,url"`
}

// CreateWebsiteResp represents a create website response
type CreateWebsiteResp struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

// GetWebsiteListReq represents a get website list request
type GetWebsiteListReq struct {
	Page     int64 `json:"page,optional,default=1"`
	PageSize int64 `json:"page_size,optional,default=20"`
}

// WebsiteItem represents a website item
type WebsiteItem struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
	URL         string `json:"url"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// GetWebsiteListResp represents a get website list response
type GetWebsiteListResp struct {
	List     []WebsiteItem `json:"list"`
	Total    int64         `json:"total"`
	Page     int64         `json:"page"`
	PageSize int64         `json:"page_size"`
}

// GetWebsiteDetailReq represents a get website detail request
type GetWebsiteDetailReq struct {
	ID int64 `json:"id" validate:"required"`
}

// GetWebsiteDetailResp represents a get website detail response
type GetWebsiteDetailResp struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
	URL         string `json:"url"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// UpdateWebsiteReq represents an update website request
type UpdateWebsiteReq struct {
	ID          int64  `json:"id" validate:"required"`
	Title       string `json:"title" validate:"required,max=200"`
	Icon        string `json:"icon,optional"`
	Description string `json:"description,optional"`
	URL         string `json:"url" validate:"required,url"`
}

// CommonResp represents a common response
type CommonResp struct {
	Message string `json:"message"`
}

// DeleteWebsiteReq represents a delete website request
type DeleteWebsiteReq struct {
	ID int64 `json:"id" validate:"required"`
}
