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
	// 1. 查询用户
	user, err := l.svcCtx.UserModel.FindByUsername(req.Username)
	if err != nil {
		if errors.Is(err, model.ErrUserNotFound) {
			return nil, errors.New("用户名或密码错误")
		}
		logx.Errorf("查询用户失败: %v", err)
		return nil, errors.New("登录失败，请稍后重试")
	}

	// 2. 检查用户状态
	if user.Status != 1 {
		return nil, errors.New("账号已被禁用")
	}

	// 3. 验证密码
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("用户名或密码错误")
	}

	// 4. 生成JWT Token
	now := utils.GetCurrentTimestamp()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	accessToken, err := utils.GenerateJwtToken(
		l.svcCtx.Config.Auth.AccessSecret,
		now,
		accessExpire,
		user.ID,
	)
	if err != nil {
		logx.Errorf("生成Token失败: %v", err)
		return nil, errors.New("登录失败，请稍后重试")
	}

	return &types.LoginResp{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		AccessToken: accessToken,
		ExpireTime:  now + accessExpire,
	}, nil
}
