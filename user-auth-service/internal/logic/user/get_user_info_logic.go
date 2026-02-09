package user

import (
	commonutils "common-utils"
	"context"
	"errors"
	"user-auth-service/internal/model"
	"user-auth-service/internal/svc"
	"user-auth-service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	// Get user ID from context (injected by JWT middleware)
	userID, err := commonutils.GetUserIDFromContext(l.ctx)
	if err != nil {
		return nil, err
	}

	// Query user information
	user, err := l.svcCtx.UserModel.FindByID(userID)
	if err != nil {
		if errors.Is(err, model.ErrUserNotFound) {
			return nil, errors.New("user not found")
		}
		logx.Errorf("Failed to find user: %v", err)
		return nil, errors.New("failed to get user information")
	}

	return &types.UserInfoResp{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Mobile:   user.Mobile,
	}, nil
}
