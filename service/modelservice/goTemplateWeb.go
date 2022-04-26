package modelservice

import (
	"go-template-web/internalapi"
)


func (l *ModelSrvLogic) GoTemplateWeb(req *internalapi.Request) (resp *internalapi.Response, err error) {
	// todo: add your logic here and delete this line

	//e := &mongoDao.Example{
	//	ID:   1,
	//	Name: "wgy",
	//}
	//err = e.Insert()
	//if err != nil {
	//	return nil, fmt.Errorf("insert err:[%w]", err)
	//}
	//
	//data, err := e.FindOne(bson.M{"id": e.ID})
	//if err != nil {
	//	return nil, fmt.Errorf("findOne by id:%v,err:%v", e.ID, err)
	//}
	////bson.D{}
	//resp = &types2.Response{
	//	Message: "success",
	//	ID:      data.ID,
	//	Name:    data.Name,
	//}
	return
}
