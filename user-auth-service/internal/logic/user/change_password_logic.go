package user

import (
	commonutils "common-utils"
	"context"
	"errors"
	"user-auth-service/internal/model"
	"user-auth-service/internal/svc"
	"user-auth-service/internal/types"
	"user-auth-service/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordLogic {
	return &ChangePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangePasswordLogic) ChangePassword(req *types.ChangePasswordReq) (resp *types.CommonResp, err error) {
	// Get user ID from context
	userID, err := commonutils.GetUserIDFromContext(l.ctx)
	if err != nil {
		return nil, err
	}

	// Query user
	user, err := l.svcCtx.UserModel.FindByID(userID)
	if err != nil {
		if errors.Is(err, model.ErrUserNotFound) {
			return nil, errors.New("user not found")
		}
		logx.Errorf("Failed to find user: %v", err)
		return nil, errors.New("failed to change password")
	}

	// Verify old password
	if !utils.CheckPasswordHash(req.OldPassword, user.Password) {
		return nil, errors.New("incorrect old password")
	}

	// Hash new password
	hashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		logx.Errorf("Failed to hash password: %v", err)
		return nil, errors.New("failed to change password")
	}

	// Update password
	if err := l.svcCtx.UserModel.UpdatePassword(userID, hashedPassword); err != nil {
		logx.Errorf("Failed to update password: %v", err)
		return nil, errors.New("failed to change password")
	}

	return &types.CommonResp{
		Message: "Password changed successfully",
	}, nil
}
