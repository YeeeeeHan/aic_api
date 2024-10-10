package dal

// import (
// 	"context"

// 	"gitlab.com/aic/aic_api/biz/model/gitlab.com/aic/http_idl_gen/gen/aic/data/models"
// 	"gitlab.com/aic/aic_api/cache"
// 	cacheHelper "gitlab.com/aic/aic_api/cache/helpers"
// 	"gitlab.com/aic/aic_api/logs"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// func ReportUser(ctx context.Context, report *models.Report) error {
// 	session, err := client.StartSession()
// 	if err != nil {
// 		return err
// 	}

// 	defer session.EndSession(ctx)

// 	_, err = session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {

// 		switch report.ContentType {
// 		case models.ContentType_CHAT:
// 			update := bson.M{
// 				"$set": bson.M{"isreported": true},
// 			}
// 			_, err := database.Collection(CHAT_COLLECTION).UpdateOne(ctx, bson.M{"id": bson.M{"$eq": report.GetChatId()}}, update)
// 			if err != nil {
// 				return nil, err
// 			}
// 		}

// 		_, err := database.Collection(REPORT_COLLECTION).InsertOne(ctx, report)
// 		if err != nil {
// 			logs.CtxError(ctx, "error creating new report: %v", err)
// 			return nil, err
// 		}

// 		_, err = database.Collection(BLOCKED_COLLECTION).
// 			UpdateOne(
// 				ctx,
// 				bson.M{"id": report.ReporterUserId},
// 				bson.M{"$addToSet": bson.M{"blocked": report.ReportedUserId}},
// 				options.Update().SetUpsert(true),
// 			)

// 		_, err = database.Collection(BLOCKED_COLLECTION).
// 			UpdateOne(
// 				ctx,
// 				bson.M{"id": report.ReportedUserId},
// 				bson.M{"$addToSet": bson.M{"blocked": report.ReporterUserId}},
// 				options.Update().SetUpsert(true),
// 			)

// 		cache.InvalidateCache(ctx, cacheHelper.GenerateCacheKey("GetMyProfileQuery.GetData", report.ReporterUserId, cache.Env))

// 		if err != nil {
// 			logs.CtxError(ctx, "error creating new block: %v", err)
// 			return nil, err
// 		}

// 		return nil, nil
// 	})
// 	return err
// }
