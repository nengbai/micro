package logic

import (
	"context"

	"micro/internal/svc"
	"micro/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MicroLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMicroLogic(ctx context.Context, svcCtx *svc.ServiceContext) MicroLogic {
	return MicroLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MicroLogic) Micro(req types.Request) (*types.Response, error) {
	// todo: add your logic here and delete this line

	return &types.Response{}, nil
}
