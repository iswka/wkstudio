package password

import (
	commonutils "common-utils"
	"context"
	"errors"
	"password-manage-service/internal/model"
	"password-manage-service/internal/svc"
	"password-manage-service/internal/types"
	"password-manage-service/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePasswordLogic {
	return &UpdatePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePasswordLogic) UpdatePassword(req *types.UpdatePasswordReq) (resp *types.CommonResp, err error) {
	// Get user ID from context
	userID, err := commonutils.GetUserIDFromContext(l.ctx)
	if err != nil {
		return nil, err
	}

	// Query password record
	passwordRecord, err := l.svcCtx.PasswordModel.FindByID(req.ID)
	if err != nil {
		if errors.Is(err, model.ErrPasswordNotFound) {
			return nil, errors.New("password record not found")
		}
		logx.Errorf("Failed to query password record: %v", err)
		return nil, errors.New("failed to update password record")
	}

	// Check ownership
	if passwordRecord.UserID != userID {
		return nil, errors.New("access denied")
	}

	// Encrypt new password
	encryptedPassword, err := utils.EncryptPassword(req.Password, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		logx.Errorf("Failed to encrypt password: %v", err)
		return nil, errors.New("failed to update password record")
	}

	// Update password record
	passwordRecord.Title = req.Title
	passwordRecord.Description = req.Description
	passwordRecord.Password = encryptedPassword

	if err := l.svcCtx.PasswordModel.Update(passwordRecord); err != nil {
		logx.Errorf("Failed to update password record: %v", err)
		return nil, errors.New("failed to update password record")
	}

	return &types.CommonResp{
		Message: "Password updated successfully",
	}, nil
}
