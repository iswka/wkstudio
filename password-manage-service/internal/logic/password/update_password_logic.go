package password

import (
	"context"
	"encoding/json"
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

	// 查询密码记录
	passwordRecord, err := l.svcCtx.PasswordModel.FindByID(req.ID)
	if err != nil {
		if errors.Is(err, model.ErrPasswordNotFound) {
			return nil, errors.New("密码记录不存在")
		}
		logx.Errorf("查询密码记录失败: %v", err)
		return nil, errors.New("更新密码记录失败")
	}

	// 检查所有权
	if passwordRecord.UserID != userID {
		return nil, errors.New("无权访问")
	}

	// 加密新密码
	encryptedPassword, err := utils.EncryptPassword(req.Password, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		logx.Errorf("密码加密失败: %v", err)
		return nil, errors.New("更新密码记录失败")
	}

	// 更新密码记录
	passwordRecord.Title = req.Title
	passwordRecord.Description = req.Description
	passwordRecord.Password = encryptedPassword

	if err := l.svcCtx.PasswordModel.Update(passwordRecord); err != nil {
		logx.Errorf("更新密码记录失败: %v", err)
		return nil, errors.New("更新密码记录失败")
	}

	return &types.CommonResp{
		Message: "更新成功",
	}, nil
}
