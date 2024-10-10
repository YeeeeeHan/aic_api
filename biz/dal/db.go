package dal

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var database *mongo.Database
var client *mongo.Client

func InitDB(isTestEnv bool) error {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	logrus.Fatal(err)
	// 	os.Exit(1)
	// }
	// client, err = mongo.NewClient(options.Client().ApplyURI(os.Getenv("URI")))
	// if err != nil {
	// 	logrus.Error("error when instantiating mongo client", err)
	// 	return err
	// }
	// ctx, _ := context.WithTimeout(context.Background(), 1000*time.Second)
	// err = client.Connect(ctx)
	// if err != nil {
	// 	logrus.Error("error when client connects to context after instantiating mongo client", err)
	// 	return err
	// }
	// err = client.Ping(ctx, readpref.Primary())
	// if err != nil {
	// 	logrus.Error("Fatal error when mongo client pings", err)
	// 	return err
	// }
	// databases, err := client.ListDatabaseNames(ctx, bson.M{})
	// if err != nil {
	// 	logrus.Error("Fatal error when enumerating databases", err)
	// 	return err
	// }
	// targetDb := "aic"
	// if isTestEnv {
	// 	targetDb = "aic_offline_environment" //TODO uncomment when deployed in production
	// }
	// database = client.Database(targetDb)
	// logs.CtxInfo(context.Background(), "databases are: %v", databases)
	// logs.CtxInfo(context.Background(), "connected to db: %s", targetDb)
	// logs.CtxInfo(context.Background(), "no errors in instantiating database clients")
	return nil
}

func GetDatabase() *mongo.Database {
	return database
}

func GetClient() *mongo.Client {
	return client
}
