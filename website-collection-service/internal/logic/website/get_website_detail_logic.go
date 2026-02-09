package website

import (
	commonutils "common-utils"
	"context"
	"errors"
	"time"
	"website-collection-service/internal/model"
	"website-collection-service/internal/svc"
	"website-collection-service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWebsiteDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetWebsiteDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWebsiteDetailLogic {
	return &GetWebsiteDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetWebsiteDetailLogic) GetWebsiteDetail(req *types.GetWebsiteDetailReq) (resp *types.GetWebsiteDetailResp, err error) {
	// Get user ID from context
	userID, err := commonutils.GetUserIDFromContext(l.ctx)
	if err != nil {
		return nil, err
	}

	// Query website record
	websiteRecord, err := l.svcCtx.WebsiteModel.FindByID(req.ID)
	if err != nil {
		if errors.Is(err, model.ErrWebsiteNotFound) {
			return nil, errors.New("website not found")
		}
		logx.Errorf("Failed to query website record: %v", err)
		return nil, errors.New("failed to get website detail")
	}

	// Check ownership
	if websiteRecord.UserID != userID {
		return nil, errors.New("access denied")
	}

	return &types.GetWebsiteDetailResp{
		ID:          websiteRecord.ID,
		Title:       websiteRecord.Title,
		Icon:        websiteRecord.Icon,
		Description: websiteRecord.Description,
		URL:         websiteRecord.URL,
		CreatedAt:   websiteRecord.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   websiteRecord.UpdatedAt.Format(time.RFC3339),
	}, nil
}
