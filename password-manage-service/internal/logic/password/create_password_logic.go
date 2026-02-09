package password

import (
	commonutils "common-utils"
	"context"
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
	// Get user ID from context
	userID, err := commonutils.GetUserIDFromContext(l.ctx)
	if err != nil {
		return nil, err
	}

	// Check if password already exists
	password, _ := l.svcCtx.PasswordModel.FindByTitleAndUserID(req.Title, userID)
	if password != nil {
		return nil, errors.New("password already exists")
	}

	// Encrypt password (using JWT secret as encryption key)
	encryptedPassword, err := utils.EncryptPassword(req.Password, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		logx.Errorf("Failed to encrypt password: %v", err)
		return nil, errors.New("failed to create password record")
	}

	// Create password record
	passwordRecord := &model.Password{
		Title:       req.Title,
		Description: req.Description,
		UserID:      userID,
		Password:    encryptedPassword,
	}

	if err := l.svcCtx.PasswordModel.Create(passwordRecord); err != nil {
		logx.Errorf("Failed to create password record: %v", err)
		return nil, errors.New("failed to create password record")
	}

	return &types.CreatePasswordResp{
		ID:      passwordRecord.ID,
		Title:   passwordRecord.Title,
		Message: "Password created successfully",
	}, nil
}
