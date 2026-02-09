package password

import (
	"context"
	"encoding/json"
	"errors"
	"password-manage-service/internal/svc"
	"password-manage-service/internal/types"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPasswordListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPasswordListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPasswordListLogic {
	return &GetPasswordListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPasswordListLogic) GetPasswordList(req *types.GetPasswordListReq) (resp *types.GetPasswordListResp, err error) {
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

	// 设置默认值
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}

	// 查询密码列表
	passwords, err := l.svcCtx.PasswordModel.FindByUserID(userID)
	if err != nil {
		logx.Errorf("查询密码列表失败: %v", err)
		return nil, errors.New("获取密码列表失败")
	}

	// 转换为响应格式
	var list []types.PasswordItem
	for _, p := range passwords {
		list = append(list, types.PasswordItem{
			ID:          p.ID,
			Title:       p.Title,
			Description: p.Description,
			CreatedAt:   p.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   p.UpdatedAt.Format(time.RFC3339),
		})
	}

	// 简单的分页处理（实际应该数据库层面分页）
	total := int64(len(list))
	start := (req.Page - 1) * req.PageSize
	end := start + req.PageSize
	if start > total {
		list = []types.PasswordItem{}
	} else if end > total {
		list = list[start:]
	} else {
		list = list[start:end]
	}

	return &types.GetPasswordListResp{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}
