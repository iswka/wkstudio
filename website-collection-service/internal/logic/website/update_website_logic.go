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

type UpdateWebsiteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateWebsiteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateWebsiteLogic {
	return &UpdateWebsiteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateWebsiteLogic) UpdateWebsite(req *types.UpdateWebsiteReq) (resp *types.CommonResp, err error) {
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
		return nil, errors.New("failed to update website record")
	}

	// Check ownership
	if websiteRecord.UserID != userID {
		return nil, errors.New("access denied")
	}

	// Update website record
	websiteRecord.Title = req.Title
	websiteRecord.Icon = req.Icon
	websiteRecord.Description = req.Description
	websiteRecord.URL = req.URL

	if err := l.svcCtx.WebsiteModel.Update(websiteRecord); err != nil {
		logx.Errorf("Failed to update website record: %v", err)
		return nil, errors.New("failed to update website record")
	}

	return &types.CommonResp{
		Message: "Website updated successfully",
	}, nil
}
