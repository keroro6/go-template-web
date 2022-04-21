package service

import (
	"context"
	"fmt"
	"go-template-web/model/mongoDao"
	types2 "go-template-web/types"
	"go.mongodb.org/mongo-driver/bson"

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

	e := &mongoDao.Example{
		ID:   1,
		Name: "wgy",
	}
	err = e.Insert()
	if err != nil {
		return nil, fmt.Errorf("insert err:[%w]", err)
	}

	data, err := e.FindOne(bson.M{"id": e.ID})
	if err != nil {
		return nil, fmt.Errorf("findOne by id:%v,err:%v", e.ID, err)
	}
	//bson.D{}
	resp = &types2.Response{
		Message: "success",
		ID:      data.ID,
		Name:    data.Name,
	}
	return
}
