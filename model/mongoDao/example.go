package mongoDao

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

type Example struct {
	ID   int64  `bson:"id"`
	Name string `bson:"name"`
}

const (
	DBName  = "go_template_mongo_1"
	ColName = "example"
)

var (
	colDao    *mongo.Collection //= GetMongoCol(DBName, ColName)
	once      sync.Once
	errGetCol error
)

func getCol(dBName, colName string) (*mongo.Collection, error) {
	once.Do(func() {
		//var err error
		colDao, errGetCol = GetMongoCol(dBName, colName)
		if errGetCol != nil {
			//todo: 日志
			fmt.Printf("GetMongoCol %+v\n", errGetCol)
		}
		if colDao == nil {
			errGetCol = fmt.Errorf("get col err,db:%v,col:%v\n", dBName, colName)
		}
	})
	return colDao, errGetCol
}
func (e *Example) Insert() error {
	col, err := getCol(DBName, ColName)
	if err != nil {
		fmt.Printf("get col %v\n", err)
		return err
	}
	ret, err := col.InsertOne(context.TODO(), e)
	if err != nil {
		fmt.Printf("insert one %v\n", err)
		return err
	}
	fmt.Printf("%+v\n", ret)
	return nil
}

func (e *Example) FindOne(q interface{}) (ret *Example, err error) {
	col, err := getCol(DBName, ColName)
	if err != nil {
		return nil, err
	}
	result := col.FindOne(context.TODO(), q)
	//if err != nil {
	//	return nil, err
	//}
	//ret.Err()
	err = result.Decode(&ret)
	if err != nil {
		return
	}
	return
}
