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

type DeleteWebsiteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteWebsiteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteWebsiteLogic {
	return &DeleteWebsiteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteWebsiteLogic) DeleteWebsite(req *types.DeleteWebsiteReq) (resp *types.CommonResp, err error) {
	// Get user ID from context
	userID, err := commonutils.GetUserIDFromContext(l.ctx)
	if err != nil {
		return nil, err
	}

	// Delete website record (ownership will be checked)
	if err := l.svcCtx.WebsiteModel.Delete(req.ID, userID); err != nil {
		if errors.Is(err, model.ErrWebsiteNotFound) {
			return nil, errors.New("website not found")
		}
		logx.Errorf("Failed to delete website record: %v", err)
		return nil, errors.New("failed to delete website record")
	}

	return &types.CommonResp{
		Message: "Website deleted successfully",
	}, nil
}
