package password

import (
	"context"
	"encoding/json"
	"errors"
	"password-manage-service/internal/svc"
	"password-manage-service/internal/types"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPasswordListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPasswordListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPasswordListLogic {
	return &GetPasswordListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPasswordListLogic) GetPasswordList(req *types.GetPasswordListReq) (resp *types.GetPasswordListResp, err error) {
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

	// Query password list
	passwords, err := l.svcCtx.PasswordModel.FindByUserID(userID)
	if err != nil {
		logx.Errorf("Failed to query password list: %v", err)
		return nil, errors.New("failed to get password list")
	}

	// Convert to response format
	var list []types.PasswordItem
	for _, p := range passwords {
		list = append(list, types.PasswordItem{
			ID:          p.ID,
			Title:       p.Title,
			Description: p.Description,
			CreatedAt:   p.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   p.UpdatedAt.Format(time.RFC3339),
		})
	}

	// Simple pagination handling (should be done at database level in production)
	total := int64(len(list))
	start := (req.Page - 1) * req.PageSize
	end := start + req.PageSize
	if start > total {
		list = []types.PasswordItem{}
	} else if end > total {
		list = list[start:]
	} else {
		list = list[start:end]
	}

	return &types.GetPasswordListResp{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}
