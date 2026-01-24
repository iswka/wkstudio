package user

import (
	"context"
	"encoding/json"
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
	// 从context中获取用户ID
	userIDValue := l.ctx.Value("user_id")
	if userIDValue == nil {
		return nil, errors.New("未授权")
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
		return nil, errors.New("无效的用户ID")
	}

	// 查询用户
	user, err := l.svcCtx.UserModel.FindByID(userID)
	if err != nil {
		if errors.Is(err, model.ErrUserNotFound) {
			return nil, errors.New("用户不存在")
		}
		logx.Errorf("查询用户失败: %v", err)
		return nil, errors.New("修改密码失败")
	}

	// 验证旧密码
	if !utils.CheckPasswordHash(req.OldPassword, user.Password) {
		return nil, errors.New("原密码错误")
	}

	// 加密新密码
	hashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		logx.Errorf("密码加密失败: %v", err)
		return nil, errors.New("修改密码失败")
	}

	// 更新密码
	if err := l.svcCtx.UserModel.UpdatePassword(userID, hashedPassword); err != nil {
		logx.Errorf("更新密码失败: %v", err)
		return nil, errors.New("修改密码失败")
	}

	return &types.CommonResp{
		Message: "密码修改成功",
	}, nil
}
