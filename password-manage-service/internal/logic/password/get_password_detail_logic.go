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
	// Get user ID from context
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

	// Query password record
	passwordRecord, err := l.svcCtx.PasswordModel.FindByID(req.ID)
	if err != nil {
		if errors.Is(err, model.ErrPasswordNotFound) {
			return nil, errors.New("password record not found")
		}
		logx.Errorf("Failed to query password record: %v", err)
		return nil, errors.New("failed to get password detail")
	}

	// Check ownership
	if passwordRecord.UserID != userID {
		return nil, errors.New("access denied")
	}

	// Decrypt password
	decryptedPassword, err := utils.DecryptPassword(passwordRecord.Password, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		logx.Errorf("Failed to decrypt password: %v", err)
		return nil, errors.New("failed to get password detail")
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
