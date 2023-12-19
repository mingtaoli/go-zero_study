package logic

import (
	"context"

	"firstapi/internal/svc"
	"firstapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FirstapiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFirstapiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FirstapiLogic {
	return &FirstapiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FirstapiLogic) Firstapi(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return &types.Response{
		Message: "hello",
	}, nil
}
