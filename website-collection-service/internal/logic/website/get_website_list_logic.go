package website

import (
	"context"
	"encoding/json"
	"errors"
	"time"
	"website-collection-service/internal/svc"
	"website-collection-service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWebsiteListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetWebsiteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWebsiteListLogic {
	return &GetWebsiteListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetWebsiteListLogic) GetWebsiteList(req *types.GetWebsiteListReq) (resp *types.GetWebsiteListResp, err error) {
	// Get user ID from context
	userIDValue := l.ctx.Value("user_id")
	if userIDValue == nil {
		return nil, errors.New("unauthorized")
	}

	var userID int64
	switch v := userIDValue.(type) {
	case float64:
		userID = int64(v)
	case json.Number:
		userID, _ = v.Int64()
	case int64:
		userID = v
	default:
		return nil, errors.New("invalid user ID")
	}

	// Set default values
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}

	// Query website list
	websites, err := l.svcCtx.WebsiteModel.FindByUserID(userID)
	if err != nil {
		logx.Errorf("Failed to query website list: %v", err)
		return nil, errors.New("failed to get website list")
	}

	// Convert to response format
	var list []types.WebsiteItem
	for _, w := range websites {
		list = append(list, types.WebsiteItem{
			ID:          w.ID,
			Title:       w.Title,
			Icon:        w.Icon,
			Description: w.Description,
			URL:         w.URL,
			CreatedAt:   w.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   w.UpdatedAt.Format(time.RFC3339),
		})
	}

	// Simple pagination handling (should be done at database level in production)
	total := int64(len(list))
	start := (req.Page - 1) * req.PageSize
	end := start + req.PageSize
	if start > total {
		list = []types.WebsiteItem{}
	} else if end > total {
		list = list[start:]
	} else {
		list = list[start:end]
	}

	return &types.GetWebsiteListResp{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}
