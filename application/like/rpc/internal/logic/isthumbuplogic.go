package logic

import (
	"context"

	"beyond-learn-go-zero/application/like/rpc/internal/svc"
	"beyond-learn-go-zero/application/like/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsThumbUpLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsThumbUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsThumbUpLogic {
	return &IsThumbUpLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsThumbUpLogic) IsThumbUp(in *service.IsThumbUpRequest) (*service.IsThumbUpResponse, error) {
	// todo: add your logic here and delete this line

	return &service.IsThumbUpResponse{}, nil
}
