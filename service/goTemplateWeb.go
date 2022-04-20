package service

import (
	"context"
	types2 "go-template-web/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoTemplateWebLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *SrvContext
}

func NewGoTemplateWebLogic(ctx context.Context, svcCtx *SrvContext) *GoTemplateWebLogic {
	return &GoTemplateWebLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoTemplateWebLogic) GoTemplateWeb(req *types2.Request) (resp *types2.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
