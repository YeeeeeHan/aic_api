package dal

// import (
// 	"context"

// 	"gitlab.com/aic/aic_api/biz/model/gitlab.com/aic/http_idl_gen/gen/aic/data/models"
// 	"gitlab.com/aic/aic_api/biz/util/helpers"
// 	"gitlab.com/aic/aic_api/cache"
// 	cacheHelper "gitlab.com/aic/aic_api/cache/helpers"
// 	"gitlab.com/aic/aic_api/logs"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// type TopicResult struct {
// 	Data []*models.Topic
// 	Page *models.Pagination
// 	Err  error
// }

// // ----------------------------GET------------------------------------
// func GetTopicsList(ctx context.Context, id string) ([]*models.Topic, error) {

// 	session, err := client.StartSession()
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer session.EndSession(ctx)

// 	topics, err := session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {

// 		var topics []*models.Topic

// 		opts := options.Find().SetSort(bson.M{"_id": -1})

// 		cursor, err := database.Collection(TOPIC_COLLECTION).Find(ctx, bson.M{"participants": bson.M{"$in": []string{id}}}, opts)

// 		if err != nil {
// 			logs.CtxError(ctx, "error while finding data from topics collection while getting topics in profile page: %v", err)
// 			return nil, err

// 		}

// 		for cursor.Next(ctx) {
// 			var topic models.Topic
// 			err := cursor.Decode(&topic)
// 			if err != nil {
// 				logs.CtxError(ctx, "error from cursor.Decode while getting topic by id: %v", err)
// 				return nil, err

// 			}

// 			topics = append(topics, &topic)
// 		}
// 		cursor.Close(ctx)

// 		return topics, nil
// 	})

// 	return topics.([]*models.Topic), err
// }

// func GetMultipleTopics(ctx context.Context, page int, id string, num int, output chan TopicResult) {

// 	session, err := client.StartSession()
// 	if err != nil {
// 		logs.CtxError(ctx, "Failed to start mongo client session: %v", err)
// 		return
// 	}
// 	defer session.EndSession(ctx)

// 	var returnedPage models.Pagination

// 	topics, err := session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {

// 		// Set the number of documents to display per page
// 		itemsPerPage := num

// 		count, err := database.Collection(TOPIC_COLLECTION).CountDocuments(context.Background(), bson.D{})
// 		if err != nil {
// 			logs.CtxError(ctx, "error getting count of Topic Collection: %v", err)
// 			return nil, err
// 		}

// 		// Calculate the number of documents to skip
// 		if count == 0 {
// 			count = 1
// 		}
// 		skip := (itemsPerPage * (page - 1)) % int(count)

// 		// Set up the options for the paginated query
// 		findOptions := options.Find().SetSkip(int64(skip)).SetLimit(int64(itemsPerPage)).SetSort((bson.M{"id": -1}))

// 		var history models.History

// 		database.Collection(HISTORY_COLLECTION).FindOne(ctx, bson.M{"id": id}).Decode(&history)

// 		history.History = append(history.History, id)

// 		blockedListMarshalled := &BlockedList{}

// 		cache.GetMarshalledCache(ctx, cacheHelper.GenerateCacheKey("blockedlist", id, cache.Env), blockedListMarshalled)

// 		cursor, err := database.
// 			Collection(TOPIC_COLLECTION).
// 			Find(
// 				ctx,
// 				bson.M{
// 					"creator": bson.M{"$nin": helpers.Ternary(len(blockedListMarshalled.Blocked) > 0, blockedListMarshalled.Blocked, []string{})},
// 				},
// 				findOptions,
// 			)

// 		if err != nil {
// 			logs.CtxError(ctx, "error while getting topics fetching topics from topic collection: %v", err)
// 			return nil, err

// 		}

// 		var topics []*models.Topic

// 		for cursor.Next(ctx) {
// 			var topic models.Topic
// 			err := cursor.Decode(&topic)
// 			if err != nil {
// 				logs.CtxError(ctx, "error from cursor.Decode while getting topics for connect: %v", err)
// 				return nil, err
// 			}

// 			for _, profileId := range topic.Participants {
// 				var profile models.ReducedProfile
// 				_ = database.Collection(PROFILE_COLLECTION).FindOne(ctx, bson.M{"id": profileId}).Decode(&profile)
// 				topic.ParticipantsData = append(topic.ParticipantsData, &profile)
// 			}

// 			topics = append(topics, &topic)
// 		}

// 		cursor.Close(ctx)

// 		if len(topics) < num {
// 			// Set up the options for the paginated query
// 			findOptions := options.Find().SetLimit(int64(itemsPerPage - len(topics))).SetSort((bson.M{"id": -1}))

// 			cursor, err := database.
// 				Collection(TOPIC_COLLECTION).
// 				Find(
// 					ctx,
// 					bson.M{
// 						"creator": bson.M{"$nin": helpers.Ternary(len(blockedListMarshalled.Blocked) > 0, blockedListMarshalled.Blocked, []string{})},
// 					},
// 					findOptions,
// 				)

// 			if err != nil {
// 				logs.CtxError(ctx, "error while finding topics from topic collection: %v", err)
// 				return nil, err
// 			}

// 			for cursor.Next(ctx) {
// 				var topic models.Topic
// 				err := cursor.Decode(&topic)
// 				if err != nil {
// 					logs.CtxError(ctx, "error from cursor.Decode while getting topics for connect: %v", err)
// 					return nil, err
// 				}

// 				for _, profileId := range topic.Participants {
// 					var profile models.ReducedProfile
// 					_ = database.Collection(PROFILE_COLLECTION).FindOne(ctx, bson.M{"id": profileId}).Decode(&profile)
// 					topic.ParticipantsData = append(topic.ParticipantsData, &profile)
// 				}

// 				topics = append(topics, &topic)

// 			}
// 			cursor.Close(ctx)
// 		}
// 		returnedPage = models.Pagination{
// 			Page:     int32(page),
// 			PerPage:  int32(itemsPerPage),
// 			LastPage: int32(count) / int32(itemsPerPage),
// 		}
// 		return topics, nil
// 	})
// 	if err != nil {
// 		output <- TopicResult{nil, &returnedPage, err}
// 		return
// 	}
// 	output <- TopicResult{topics.([]*models.Topic), &returnedPage, nil}
// 	return
// }

// func GetTopicsInConnect(ctx context.Context, page int, id string) ([]*models.Topic, *models.Pagination, error) {
// 	var topics []*models.Topic

// 	// Set the number of documents to display per page
// 	itemsPerPage := 3

// 	// Calculate the number of documents to skip
// 	skip := itemsPerPage * (page - 1)

// 	// Set up the options for the paginated query
// 	findOptions := options.Find().SetSkip(int64(skip)).SetLimit(int64(itemsPerPage)).SetSort((bson.M{"id": -1}))

// 	var history models.History

// 	database.Collection(HISTORY_COLLECTION).FindOne(ctx, bson.M{"id": id}).Decode(&history)

// 	// history.History = append(history.History)

// 	cursor, err := database.
// 		Collection(TOPIC_COLLECTION).
// 		Find(
// 			ctx,
// 			bson.M{
// 				"participants": bson.M{"$nin": []string{id}},
// 				"id":           bson.M{"$nin": history.History},
// 			},
// 			findOptions,
// 		)

// 	if err != nil {
// 		logs.CtxError(ctx, "error while finding data from profiles collection while getting profile in connect page: %v", err)
// 		return nil, nil, err

// 	}

// 	count, err := database.Collection(TOPIC_COLLECTION).CountDocuments(ctx, bson.M{})

// 	if err != nil {
// 		logs.CtxError(ctx, "error while finding data from profiles collection while counting profiles in connect page: %v", err)
// 		return nil, nil, err

// 	}

// 	for cursor.Next(ctx) {
// 		var topic models.Topic
// 		err := cursor.Decode(&topic)
// 		if err != nil {
// 			logs.CtxError(ctx, "error from cursor.Decode while getting profile by id: %v", err)
// 			return nil, nil, err

// 		}

// 		for _, profileId := range topic.Participants {
// 			var profile models.ReducedProfile
// 			_ = database.Collection(PROFILE_COLLECTION).FindOne(ctx, bson.M{"id": profileId}).Decode(&profile)
// 			topic.ParticipantsData = append(topic.ParticipantsData, &profile)
// 		}

// 		topics = append(topics, &topic)
// 	}

// 	cursor.Close(ctx)

// 	if err := cursor.Err(); err != nil {
// 		logs.CtxError(ctx, "error from cursor while getting profile by id: %v", err)
// 	}

// 	cursor.Close(ctx)

// 	pagination := models.Pagination{
// 		Page:     int32(page),
// 		PerPage:  int32(itemsPerPage),
// 		LastPage: int32(count) / int32(itemsPerPage),
// 	}

// 	return topics, &pagination, nil
// }

// func GetTopicById(ctx context.Context, id string) (*models.Topic, error) {
// 	var topic models.Topic

// 	err := database.Collection(TOPIC_COLLECTION).FindOne(ctx, bson.M{"id": id}).Decode(&topic)

// 	for _, profileId := range topic.Participants {
// 		var profile models.ReducedProfile
// 		_ = database.Collection(PROFILE_COLLECTION).FindOne(ctx, bson.M{"id": profileId}).Decode(&profile)
// 		topic.ParticipantsData = append(topic.ParticipantsData, &profile)
// 	}
// 	if err != nil {
// 		logs.CtxError(ctx, "error while getting profile by id: %v", err)
// 		return nil, err
// 	}
// 	return &topic, nil
// }

// // ----------------------------CREATE------------------------------------

// func AddNewTopic(ctx context.Context, topic *models.NewTopic) error {
// 	_, err := database.Collection(TOPIC_COLLECTION).InsertOne(ctx, topic)
// 	if err != nil {
// 		logs.CtxError(ctx, "error creating new topic: %v", err)
// 		return err
// 	}

// 	return err
// }

// // ----------------------------DELETE------------------------------------

// func DeleteTopic(ctx context.Context, idList []string) error {

// 	session, err := client.StartSession()
// 	if err != nil {

// 		return err
// 	}
// 	defer session.EndSession(ctx)

// 	_, err = session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {

// 		res, err := database.Collection(INTERACTION_COLLECTION).DeleteMany(ctx, bson.M{"topicid": bson.M{"$in": idList}})
// 		if err != nil {
// 			return nil, err
// 		}
// 		logs.CtxInfo(ctx, "deleted topics, res: %v", res)

// 		res, err = database.Collection(TOPIC_COLLECTION).DeleteMany(ctx, bson.M{"id": bson.M{"$in": idList}})
// 		if err != nil {
// 			return nil, err
// 		}
// 		logs.CtxInfo(ctx, "deleted topics, res: %v", res)

// 		return nil, nil
// 	})

// 	return err
// }

// func DeleteTopicsByProfileId(ctx context.Context, profileHeader string) error {
// 	filter := bson.M{"creator": profileHeader}

// 	result, err := database.Collection(COFFEE_REQUEST_COLLECTION).DeleteMany(ctx, filter)
// 	logs.CtxInfo(ctx, "deleted topics results:%v", result)
// 	if err != nil {
// 		logs.CtxError(ctx, "error deleting topics: %v", err)
// 	}
// 	return nil
// }

// // -----------------------------UPDATE-----------------------------------
// func UpdateTopic(ctx context.Context, topic *models.Topic) error {
// 	filter := bson.M{"id": topic.Id}
// 	// Note: be careful when using ReplaceOne, ensure that all the fields in the intended topic are sent, if it is missing, it will be updated as null
// 	_, err := database.Collection(TOPIC_COLLECTION).ReplaceOne(ctx, filter, topic)

// 	if err != nil {
// 		logs.CtxError(ctx, "error updating topic: %v", err)
// 		return err
// 	}

// 	logs.CtxInfo(ctx, "topic updated for id: %s", topic.Id)
// 	return nil
// }
