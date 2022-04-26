package mongoDao

import (
	"fmt"
	"go-template-web/dao/mongoDao"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

// Example xiugai
type Example struct {
	ID   int64  `bson:"id"`
	Name string `bson:"name"`
}

// DBName 自己改
const (
	DBName  = "go_template_mongo_1"
	ColName = "example"
)

var (
	colDao    *mongo.Collection //= GetMongoCol(DBName, ColName)
	once      sync.Once
	errGetCol error
)

// GetCol 固定写法别动
func GetCol(dBName, colName string) (*mongo.Collection, error) {
	once.Do(func() {
		//var err error
		colDao, errGetCol = mongoDao.GetMongoCol(dBName, colName)
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
