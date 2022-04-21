package mongoDao

import (
	"context"
	"fmt"
	"go-template-web/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var mongoDBMap = map[string]*mongo.Client{}
var colMap = map[string]*mongo.Collection{}

// InitMongo 初始化mongodb
func InitMongo(c config.Config) {
	if len(c.MongoDB) <= 0 {
		return
	}
	for _, mongoConfig := range c.MongoDB {
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoConfig.Dsn))
		if err != nil || client == nil {
			panic(err)
		}
		fmt.Printf("%+v\n", mongoConfig)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		if err = client.Ping(ctx, readpref.Primary()); err != nil {
			panic(err)
		}
		fmt.Println("Mongo successfully connected and pinged.")
		mongoDBMap[mongoConfig.DBName] = client
		cancel()
	}
}

// CloseMongo mongo close
func CloseMongo() {
	for _, client := range mongoDBMap {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}

	}
}

func GetMongoCol(dBName, tableName string) (col *mongo.Collection, err error) {
	client, ok := mongoDBMap[dBName]
	fmt.Println(mongoDBMap)
	if !ok {
		err = fmt.Errorf("db: %s not exists", dBName)
		return
	}
	//设置database级别的优先从从库读
	//todo: 没必要每次创建database和collection对象
	col = client.Database(dBName, &options.DatabaseOptions{
		ReadPreference: readpref.SecondaryPreferred(),
	}).Collection(tableName)
	if col == nil {
		err = fmt.Errorf("get collection err, db:%s, col:%s", dBName, tableName)
		return
	}
	return
}
