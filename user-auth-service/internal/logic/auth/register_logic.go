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

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// 1. 检查用户名是否已存在
	exists, err := l.svcCtx.UserModel.CheckUsernameExists(req.Username)
	if err != nil {
		logx.Errorf("检查用户名失败: %v", err)
		return nil, errors.New("注册失败，请稍后重试")
	}
	if exists {
		return nil, errors.New("用户名已存在")
	}

	// 2. 检查邮箱是否已存在
	exists, err = l.svcCtx.UserModel.CheckEmailExists(req.Email)
	if err != nil {
		logx.Errorf("检查邮箱失败: %v", err)
		return nil, errors.New("注册失败，请稍后重试")
	}
	if exists {
		return nil, errors.New("邮箱已被使用")
	}

	// 3. 密码加密
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		logx.Errorf("密码加密失败: %v", err)
		return nil, errors.New("注册失败，请稍后重试")
	}

	// 4. 创建用户
	user := &model.User{
		Username: req.Username,
		Password: hashedPassword,
		Email:    req.Email,
		Mobile:   req.Mobile,
		Status:   1,
	}

	if err := l.svcCtx.UserModel.Create(user); err != nil {
		logx.Errorf("创建用户失败: %v", err)
		return nil, errors.New("注册失败，请稍后重试")
	}

	return &types.RegisterResp{
		ID:       user.ID,
		Username: user.Username,
		Message:  "注册成功",
	}, nil
}
