package logic

import (
	"beyond-learn-go-zero/application/article/rpc/internal/model"
	"beyond-learn-go-zero/application/article/rpc/internal/svc"
	"beyond-learn-go-zero/application/article/rpc/pb"
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishLogic {
	return &PublishLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishLogic) Publish(in *pb.PublishRequest) (*pb.PublishResponse, error) {
	ret, err := l.svcCtx.ArticleModel.Insert(l.ctx, &model.Article{
		Title:       in.Title,
		Content:     in.Content,
		Cover:       in.Cover,
		Description: in.Description,
		AuthorId:    in.UserId,
		PublishTime: time.Now(),
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
	})
	if err != nil {
		l.Logger.Errorf("Publish Insert req: %v error: %v", in, err)
		return nil, err
	}

	articleId, err := ret.LastInsertId()
	if err != nil {
		l.Logger.Errorf("LastInsertId error: %v", err)
		return nil, err
	}

	return &pb.PublishResponse{ArticleId: articleId}, nil
}
