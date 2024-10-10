package dal

// import (
// 	"context"

// 	"github.com/sirupsen/logrus"
// 	"gitlab.com/aic/aic_api/biz/model/gitlab.com/aic/http_idl_gen/gen/aic/data/models"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// func AddNewFeedback(ctx context.Context, feedback *models.Feedback) error {
// 	session, err := client.StartSession()
// 	if err != nil {
// 		return err

// 	}
// 	defer session.EndSession(ctx)

// 	_, err = session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {

// 		_, err := database.Collection(FEEDBACK_COLLECTION).InsertOne(ctx, feedback)
// 		if err != nil {
// 			logrus.Error("Insert Error : ", err)
// 			return nil, err
// 		}
// 		return nil, nil
// 	})

// 	return err
// }