package auth

import (
	"context"
	"errors"
	"user-auth-service/internal/model"
	"user-auth-service/internal/svc"
	"user-auth-service/internal/types"
	"user-auth-service/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// 1. Find user by username
	user, err := l.svcCtx.UserModel.FindByUsername(req.Username)
	if err != nil {
		if errors.Is(err, model.ErrUserNotFound) {
			return nil, errors.New("invalid username or password")
		}
		logx.Errorf("Failed to find user: %v", err)
		return nil, errors.New("login failed, please try again later")
	}

	// 2. Check user status
	if user.Status != 1 {
		return nil, errors.New("account has been disabled")
	}

	// 3. Verify password
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("invalid username or password")
	}

	// 4. Generate JWT Token
	now := utils.GetCurrentTimestamp()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	accessToken, err := utils.GenerateJwtToken(
		l.svcCtx.Config.Auth.AccessSecret,
		now,
		accessExpire,
		user.ID,
	)
	if err != nil {
		logx.Errorf("Failed to generate token: %v", err)
		return nil, errors.New("login failed, please try again later")
	}

	return &types.LoginResp{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		AccessToken: accessToken,
		ExpireTime:  now + accessExpire,
	}, nil
}
