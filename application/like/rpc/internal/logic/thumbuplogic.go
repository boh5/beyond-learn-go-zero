package logic

import (
	"beyond-learn-go-zero/application/like/rpc/internal/types"
	"context"
	"encoding/json"

	"beyond-learn-go-zero/application/like/rpc/internal/svc"
	"beyond-learn-go-zero/application/like/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
)

type ThumbUpLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewThumbUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThumbUpLogic {
	return &ThumbUpLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ThumbUpLogic) ThumbUp(in *service.ThumbUpRequest) (*service.ThumbUpResponse, error) {
	// todo: 暂时忽略逻辑
	// 1. 查询是否点赞过
	// 2. 计算当前内容的总点赞数和点踩数

	msg := &types.ThumbUpMsg{
		BizId:    in.BizId,
		ObjId:    in.ObjId,
		UserId:   in.UserId,
		LikeType: in.LikeType,
	}

	// 发送kafka消息，异步
	threading.GoSafe(func() {
		data, err := json.Marshal(msg)
		if err != nil {
			l.Logger.Errorf("[ThumbUp] marshal msg: %v error: %v", msg, err)
			return
		}
		err = l.svcCtx.KqPusherClient.Push(string(data))
		if err != nil {
			l.Logger.Errorf("[ThumbUp] kq push data: %s error: %v", data, err)
		}
	})

	return &service.ThumbUpResponse{}, nil
}
