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

type CreatePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePasswordLogic {
	return &CreatePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePasswordLogic) CreatePassword(req *types.CreatePasswordReq) (resp *types.CreatePasswordResp, err error) {
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

	// 加密密码（使用JWT密钥作为加密密钥）
	encryptedPassword, err := utils.EncryptPassword(req.Password, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		logx.Errorf("密码加密失败: %v", err)
		return nil, errors.New("创建密码记录失败")
	}

	// 创建密码记录
	passwordRecord := &model.Password{
		Title:       req.Title,
		Description: req.Description,
		UserID:      userID,
		Password:    encryptedPassword,
	}

	if err := l.svcCtx.PasswordModel.Create(passwordRecord); err != nil {
		logx.Errorf("创建密码记录失败: %v", err)
		return nil, errors.New("创建密码记录失败")
	}

	return &types.CreatePasswordResp{
		ID:      passwordRecord.ID,
		Title:   passwordRecord.Title,
		Message: "创建成功",
	}, nil
}
