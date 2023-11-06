package svc

import (
	"beyond-learn-go-zero/application/article/api/internal/config"
	"beyond-learn-go-zero/application/article/rpc/article"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/zeromicro/go-zero/zrpc"
)

const (
	defaultCosTimeout = 10
)

type ServiceContext struct {
	Config     config.Config
	CosClient  *cos.Client
	ArticleRPC article.Article
}

func NewServiceContext(c config.Config) *ServiceContext {
	if c.Cos.TimeoutSec == 0 {
		c.Cos.TimeoutSec = defaultCosTimeout
	}
	timeout := time.Duration(c.Cos.TimeoutSec) * time.Second
	u, _ := url.Parse(c.Cos.BucketURL)
	su, _ := url.Parse(c.Cos.ServiceURL)
	b := &cos.BaseURL{BucketURL: u, ServiceURL: su}
	cc := cos.NewClient(b, &http.Client{
		Timeout: timeout,
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv(c.Cos.SecretID),
			SecretKey: os.Getenv(c.Cos.SecretKey),
		},
	})

	return &ServiceContext{
		Config:     c,
		CosClient:  cc,
		ArticleRPC: article.NewArticle(zrpc.MustNewClient(c.ArticleRPC)),
	}
}
