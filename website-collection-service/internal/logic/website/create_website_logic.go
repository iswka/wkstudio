package website

import (
	commonutils "common-utils"
	"context"
	"errors"
	"website-collection-service/internal/model"
	"website-collection-service/internal/svc"
	"website-collection-service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateWebsiteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateWebsiteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateWebsiteLogic {
	return &CreateWebsiteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateWebsiteLogic) CreateWebsite(req *types.CreateWebsiteReq) (resp *types.CreateWebsiteResp, err error) {
	// Get user ID from context
	userID, err := commonutils.GetUserIDFromContext(l.ctx)
	if err != nil {
		return nil, err
	}

	// Create website record
	websiteRecord := &model.Website{
		Title:       req.Title,
		Icon:        req.Icon,
		Description: req.Description,
		URL:         req.URL,
		UserID:      userID,
	}

	if err := l.svcCtx.WebsiteModel.Create(websiteRecord); err != nil {
		logx.Errorf("Failed to create website record: %v", err)
		return nil, errors.New("failed to create website record")
	}

	return &types.CreateWebsiteResp{
		ID:      websiteRecord.ID,
		Title:   websiteRecord.Title,
		Message: "Website created successfully",
	}, nil
}
