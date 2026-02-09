package user

import (
	"context"
	"encoding/json"
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
