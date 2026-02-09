package password

import (
	"context"
	"encoding/json"
	"errors"
	"password-manage-service/internal/model"
	"password-manage-service/internal/svc"
	"password-manage-service/internal/types"
	"password-manage-service/internal/utils"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPasswordDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPasswordDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPasswordDetailLogic {
	return &GetPasswordDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPasswordDetailLogic) GetPasswordDetail(req *types.GetPasswordDetailReq) (resp *types.GetPasswordDetailResp, err error) {
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
		return nil, errors.New("获取密码详情失败")
	}

	// 检查所有权
	if passwordRecord.UserID != userID {
		return nil, errors.New("无权访问")
	}

	// 解密密码
	decryptedPassword, err := utils.DecryptPassword(passwordRecord.Password, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		logx.Errorf("密码解密失败: %v", err)
		return nil, errors.New("获取密码详情失败")
	}

	return &types.GetPasswordDetailResp{
		ID:          passwordRecord.ID,
		Title:       passwordRecord.Title,
		Description: passwordRecord.Description,
		Password:    decryptedPassword,
		CreatedAt:   passwordRecord.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   passwordRecord.UpdatedAt.Format(time.RFC3339),
	}, nil
}
