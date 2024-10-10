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

// type ProfileResult struct {
// 	Data []*models.PublicProfile
// 	Page *models.Pagination
// 	Err  error
// }

// // -----------------------------GET-----------------------------------

// // func NewProfile(ctx context.Context, p *models.Profile) (primitive.ObjectID, error) {

// // 	session, err := client.StartSession()
// // 	if err != nil {
// // 		return primitive.NilObjectID, err
// // 	}
// // 	defer session.EndSession(ctx)

// // 	r, err := session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {

// // 		res, err := database.Collection("profile").InsertOne(ctx, p)
// // 		if err != nil {
// // 			logs.CtxError(ctx,"Insert Error : ", err)
// // 			return primitive.NilObjectID, err
// // 		}

// // 		logrus.Info("New profile added into mongodb", p.ProfileLinkedin.LinkedinId)

// // 		r := (res.InsertedID.(primitive.ObjectID))

// // 		_, err = database.Collection("history").InsertOne(ctx,
// // 			bson.M{"id": res.InsertedID.(primitive.ObjectID),
// // 				"history": []string{}})

// // 		return r, err
// // 	})
// // 	return r.(primitive.ObjectID), err
// // }

// type BlockedList struct {
// 	Id      string   `json:"id"`
// 	Blocked []string `json:"blockedlist"`
// }

// func GetProfileById(ctx context.Context, id string) (*models.Profile, error) {
// 	var result models.Profile

// 	err := database.Collection(PROFILE_COLLECTION).FindOne(ctx, bson.M{"id": id}).Decode(&result)
// 	if err != nil {
// 		logs.CtxError(ctx, "error while getting profile by id: %v", err)
// 		return nil, err
// 	}
// 	return &result, nil
// }

// func GetBlockedListById(ctx context.Context, id string) *BlockedList {
// 	var blockedList BlockedList

// 	err := database.Collection(BLOCKED_COLLECTION).FindOne(ctx, bson.M{"id": id}).Decode(&blockedList)
// 	if err != nil {
// 		logs.CtxWarn(ctx, "user has no blocked list %v", err)
// 		return nil
// 	}
// 	return &blockedList
// }

// // func GetProfileByMongoId(ctx context.Context, id primitive.ObjectID) (*models.Profile, error) {
// // 	var result models.Profile
// // 	err := database.Collection("profile").FindOne(ctx, bson.M{"id": id}).Decode(&result)
// // 	if err != nil {
// // 		logs.CtxError(ctx,"Error while getting profile by id: ", err)
// // 		return nil, err
// // 	}
// // 	return &result, nil
// // }

// func GetMultipleProfiles(ctx context.Context, page int, id string, num int, output chan ProfileResult) {

// 	session, err := client.StartSession()
// 	if err != nil {
// 		output <- ProfileResult{nil, nil, err}
// 		return
// 	}
// 	defer session.EndSession(ctx)

// 	var returnedPage models.Pagination

// 	profiles, err := session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {

// 		// Set the number of documents to display per page
// 		itemsPerPage := num

// 		count, err := database.Collection(PROFILE_COLLECTION).CountDocuments(context.Background(), bson.D{})
// 		if err != nil {
// 			logs.CtxError(ctx, "error getting count of Profile Collection: %v", err)
// 			return nil, err
// 		}

// 		// Calculate the number of documents to skip
// 		if count == 0 {
// 			count = 1
// 		}
// 		skip := (itemsPerPage * (page - 1)) % int(count)

// 		// Set up the options for the paginated query
// 		findOptions := options.Find().SetSkip(int64(skip)).SetLimit(int64(itemsPerPage)).SetSort((bson.M{"id": -1}))

// 		blockedListMarshalled := &BlockedList{}

// 		cache.GetMarshalledCache(ctx, cacheHelper.GenerateCacheKey("blockedlist", id, cache.Env), blockedListMarshalled)

// 		cursor, err := database.
// 			Collection(PROFILE_COLLECTION).
// 			Find(
// 				ctx,
// 				bson.M{
// 					"id": bson.M{"$ne": id, "$nin": helpers.Ternary(len(blockedListMarshalled.Blocked) > 0, blockedListMarshalled.Blocked, []string{})},
// 				},
// 				findOptions,
// 			)

// 		if err != nil {
// 			logs.CtxError(ctx, "error while fetching profiles from profile collection: %v", err)
// 			return nil, err

// 		}

// 		var profiles []*models.PublicProfile

// 		for cursor.Next(ctx) {
// 			var profile models.PublicProfile
// 			err := cursor.Decode(&profile)
// 			if err != nil {
// 				logs.CtxError(ctx, "error from cursor.Decode while getting profiles for connect: %v", err)
// 				return nil, err
// 			}
// 			profiles = append(profiles, &profile)
// 		}

// 		cursor.Close(ctx)

// 		if len(profiles) < num {
// 			// Set up the options for the paginated query
// 			findOptions := options.Find().SetLimit(int64(itemsPerPage - len(profiles))).SetSort((bson.M{"id": -1}))

// 			cursor, err := database.
// 				Collection(PROFILE_COLLECTION).
// 				Find(
// 					ctx,
// 					bson.M{
// 						"id": bson.M{"$ne": id},
// 					},
// 					findOptions,
// 				)

// 			if err != nil {
// 				logs.CtxError(ctx, "error while finding profiles from profile collection: %v", err)
// 				return nil, err
// 			}

// 			for cursor.Next(ctx) {
// 				var profile models.PublicProfile
// 				err := cursor.Decode(&profile)
// 				if err != nil {
// 					logs.CtxError(ctx, "error from cursor.Decode while getting profiles for homepage: %v", err)
// 					return nil, err
// 				}

// 				profiles = append(profiles, &profile)

// 			}
// 			cursor.Close(ctx)
// 		}

// 		if err := cursor.Err(); err != nil {
// 			logs.CtxError(ctx, "error from cursor while getting profile by id: %v", err)
// 		}

// 		cursor.Close(ctx)

// 		returnedPage = models.Pagination{
// 			Page:     int32(page),
// 			PerPage:  int32(itemsPerPage),
// 			LastPage: int32(count) / int32(itemsPerPage),
// 		}

// 		return profiles, nil
// 	})

// 	if err != nil {
// 		output <- ProfileResult{nil, &returnedPage, err}
// 		return
// 	}
// 	output <- ProfileResult{profiles.([]*models.PublicProfile), &returnedPage, nil}
// 	return
// }

// func GetReducedProfileById(ctx context.Context, id string) (*models.ReducedProfile, error) {
// 	var result models.ReducedProfile
// 	err := database.Collection(PROFILE_COLLECTION).FindOne(ctx, bson.M{"id": id}).Decode(&result)
// 	if err != nil {
// 		logs.CtxError(ctx, "error while getting profile by id: %v", err)
// 		return nil, err
// 	}
// 	return &result, nil
// }

// func GetProfileLinkedIn(ctx context.Context, id string) (*models.ProfileLinkedIn, error) {
// 	var result models.Profile
// 	// projection := bson.M{
// 	// 	"profilelinkedin": 1, // Include profilelinkedin field
// 	// }

// 	err := database.Collection(PROFILE_COLLECTION).
// 		FindOne(ctx, bson.M{"id": id}).
// 		Decode(&result)
// 	if err != nil {
// 		logs.CtxError(ctx, "error while getting profile by id: %v", err)
// 		return nil, err
// 	}
// 	return result.ProfileLinkedin, nil
// }

// func CheckProfileByLinkedInId(ctx context.Context, id string) (*models.Profile, error) {
// 	var p models.Profile
// 	err := database.Collection(PROFILE_COLLECTION).FindOne(ctx, bson.M{"profilelinkedin.linkedinid": id}).Decode(&p)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &p, nil
// }

// // -----------------------------CREATE-----------------------------------
// func newProfile(ctx context.Context, p *models.Profile, id string) error {
// 	session, err := client.StartSession()
// 	if err != nil {
// 		return err
// 	}
// 	defer session.EndSession(ctx)

// 	_, err = session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
// 		// Perform multiple operations within the transaction using sessCtx
// 		_, err := database.Collection(PROFILE_COLLECTION).
// 			InsertOne(ctx, p)
// 		if err != nil {
// 			logs.CtxError(ctx, "error creating new profile: %v", err)
// 			return "", err
// 		}

// 		logs.CtxInfo(ctx, "new profile added into mongodb, profile linked in id: %s", p.ProfileLinkedin.LinkedinId)

// 		database.Collection(HISTORY_COLLECTION).
// 			InsertOne(ctx,
// 				bson.M{"id": id,
// 					"history": []string{}})
// 		return nil, nil
// 	},
// 	)

// 	return err
// }

// // -----------------------------DELETE-----------------------------------
// // -----------------------------UPDATE-----------------------------------
// func UpdateProfile(ctx context.Context, p *models.Profile, id string) error {

// 	filter := bson.M{"id": id}

// 	update := bson.M{"$set": bson.M{
// 		"studies":                          helpers.CleanText(p.Studies),
// 		"header":                           helpers.CleanText(p.Header),
// 		"introduction":                     helpers.CleanText(p.Introduction),
// 		"interests":                        p.Interests,
// 		"modules":                          p.Modules,
// 		"profilepretexts":                  p.ProfilePretexts,
// 		"telegram":                         helpers.CleanText(p.Telegram),
// 		"profilelinkedin.linkedinheadline": helpers.CleanText(p.Header),
// 		"linkedinlink":                     helpers.CleanText(p.LinkedInLink),
// 		"networkingintentions":             p.NetworkingIntentions,
// 		"clubssocieties":                   p.ClubsSocieties,
// 		"hobbies":                          p.Hobbies,
// 	}}
// 	if p.ProfileLinkedin != nil {
// 		update = bson.M{"$set": bson.M{
// 			"profilelinkedin.linkedinfirstname": p.ProfileLinkedin.LinkedinFirstName,
// 			"profilelinkedin.linkedinlastname":  p.ProfileLinkedin.LinkedinLastName,
// 			"studies":                           helpers.CleanText(p.Studies),
// 			"header":                            helpers.CleanText(p.Header),
// 			"introduction":                      helpers.CleanText(p.Introduction),
// 			"interests":                         p.Interests,
// 			"modules":                           p.Modules,
// 			"profilepretexts":                   p.ProfilePretexts,
// 			"telegram":                          helpers.CleanText(p.Telegram),
// 			"profilelinkedin.linkedinheadline":  helpers.CleanText(p.Header),
// 			"linkedinlink":                      helpers.CleanText(p.LinkedInLink),
// 			"networkingintentions":              p.NetworkingIntentions,
// 			"clubssocieties":                    p.ClubsSocieties,
// 			"hobbies":                           p.Hobbies,
// 		}}
// 	}

// 	_, err := database.Collection(PROFILE_COLLECTION).UpdateOne(ctx, filter, update)

// 	if err != nil {
// 		logs.CtxError(ctx, "error updating profile: %v", err)
// 		return err
// 	}

// 	logs.CtxInfo(ctx, "profile data updated for id: %s", id)
// 	return nil
// }

// func UpdateProfileDp(ctx context.Context, id, dpLink string) error {

// 	filter := bson.M{"id": id}

// 	update := bson.M{"$set": bson.M{
// 		"profilelinkedin.linkedinprofileurl": dpLink,
// 	}}

// 	_, err := database.Collection(PROFILE_COLLECTION).UpdateOne(ctx, filter, update)

// 	if err != nil {
// 		logs.CtxError(ctx, "error updating profile: %v", err)
// 		return err
// 	}

// 	return nil
// }
