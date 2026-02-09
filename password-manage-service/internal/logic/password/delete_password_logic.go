package password

import (
	"context"
	"encoding/json"
	"errors"
	"password-manage-service/internal/model"
	"password-manage-service/internal/svc"
	"password-manage-service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePasswordLogic {
	return &DeletePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePasswordLogic) DeletePassword(req *types.DeletePasswordReq) (resp *types.CommonResp, err error) {
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

	// 删除密码记录（会检查所有权）
	if err := l.svcCtx.PasswordModel.Delete(req.ID, userID); err != nil {
		if errors.Is(err, model.ErrPasswordNotFound) {
			return nil, errors.New("密码记录不存在")
		}
		logx.Errorf("删除密码记录失败: %v", err)
		return nil, errors.New("删除密码记录失败")
	}

	return &types.CommonResp{
		Message: "删除成功",
	}, nil
}
