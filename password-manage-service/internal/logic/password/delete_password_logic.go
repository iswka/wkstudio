package password

import (
	commonutils "common-utils"
	"context"
	"errors"
	"password-manage-service/internal/model"
	"password-manage-service/internal/svc"
	"password-manage-service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePasswordLogic {
	return &DeletePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePasswordLogic) DeletePassword(req *types.DeletePasswordReq) (resp *types.CommonResp, err error) {
	// Get user ID from context
	userID, err := commonutils.GetUserIDFromContext(l.ctx)
	if err != nil {
		return nil, err
	}

	// Delete password record (ownership will be checked)
	if err := l.svcCtx.PasswordModel.Delete(req.ID, userID); err != nil {
		if errors.Is(err, model.ErrPasswordNotFound) {
			return nil, errors.New("password record not found")
		}
		logx.Errorf("Failed to delete password record: %v", err)
		return nil, errors.New("failed to delete password record")
	}

	return &types.CommonResp{
		Message: "Password deleted successfully",
	}, nil
}
