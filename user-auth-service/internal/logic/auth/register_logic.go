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
	// 1. Check if username already exists
	exists, err := l.svcCtx.UserModel.CheckUsernameExists(req.Username)
	if err != nil {
		logx.Errorf("Failed to check username: %v", err)
		return nil, errors.New("registration failed, please try again later")
	}
	if exists {
		return nil, errors.New("username already exists")
	}

	// 2. Check if email already exists
	exists, err = l.svcCtx.UserModel.CheckEmailExists(req.Email)
	if err != nil {
		logx.Errorf("Failed to check email: %v", err)
		return nil, errors.New("registration failed, please try again later")
	}
	if exists {
		return nil, errors.New("email already in use")
	}

	// 3. Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		logx.Errorf("Failed to hash password: %v", err)
		return nil, errors.New("registration failed, please try again later")
	}

	// 4. Create user
	user := &model.User{
		Username: req.Username,
		Password: hashedPassword,
		Email:    req.Email,
		Mobile:   req.Mobile,
		Status:   1,
	}

	if err := l.svcCtx.UserModel.Create(user); err != nil {
		logx.Errorf("Failed to create user: %v", err)
		return nil, errors.New("registration failed, please try again later")
	}

	return &types.RegisterResp{
		ID:       user.ID,
		Username: user.Username,
		Message:  "Registration successful",
	}, nil
}
