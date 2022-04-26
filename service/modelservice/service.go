package modelservice

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-template-web/service"
)

type ModelSrvLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *service.SrvContext
}

func NewModelSrvLogic(ctx context.Context, svcCtx *service.SrvContext) *ModelSrvLogic {
	return &ModelSrvLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
