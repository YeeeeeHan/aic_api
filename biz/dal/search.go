package dal

// import (
// 	"context"
// 	"regexp"

// 	"gitlab.com/aic/aic_api/biz/model/gitlab.com/aic/http_idl_gen/gen/aic/data/models"
// 	"gitlab.com/aic/aic_api/logs"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// type TopicFields struct {
// 	Id            string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" form:"id" query:"id"`
// 	Creator       string                `protobuf:"bytes,2,rep,name=creator,proto3" json:"creator,omitempty" form:"creator" query:"creator"`
// 	TopicSettings *models.TopicSettings `protobuf:"bytes,30,opt,name=topic_settings,json=topicSettings,proto3" json:"topic_settings,omitempty" form:"topic_settings" query:"topic_settings"`
// 	ContentType   *models.ContentType   `protobuf:"varint,5,opt,name=content_type,json=contentType,proto3,enum=aic_api.ContentType" json:"content_type,omitempty" form:"content_type" query:"content_type"`
// 	Participants  []string              `protobuf:"bytes,10,rep,name=participants,proto3" json:"participants" form:"participants" query:"participants"`
// }

// type ProfileFields struct {
// 	Id              string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" form:"id" query:"id"`
// 	ProfileLinkedin *models.ProfileLinkedIn `protobuf:"bytes,3,opt,name=profile_linkedin,json=profileLinkedin,proto3" json:"profile_linkedin,omitempty" form:"profile_linkedin" query:"profile_linkedin"`
// 	Interests       []string                `protobuf:"bytes,40,rep,name=interests,proto3" json:"interests" form:"interests" query:"interests"`
// 	Header          string                  `protobuf:"bytes,21,opt,name=header,proto3" json:"header,omitempty" form:"header" query:"header"`
// 	Introduction    string                  `protobuf:"bytes,22,opt,name=introduction,proto3" json:"introduction,omitempty" form:"introduction" query:"introduction"`
// }

// type EventFields struct {
// 	Id                string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" form:"id" query:"id"`
// 	Title             string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty" form:"title" query:"title"`
// 	Tags              []string `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags" form:"tags" query:"tags"`
// 	CompanyPictureUrl string   `protobuf:"bytes,21,opt,name=company_picture_url,json=companyPictureUrl,proto3" json:"company_picture_url,omitempty" form:"company_picture_url" query:"company_picture_url"`
// 	CompanyName       string   `protobuf:"bytes,20,opt,name=company_name,json=companyName,proto3" json:"company_name,omitempty" form:"company_name" query:"company_name"`
// }

// func AllSearch(ctx context.Context, name string) ([]*models.Topic, []*models.Event, []*models.PublicProfile, error) {
// 	session, err := client.StartSession()
// 	if err != nil {
// 		return nil, nil, nil, err
// 	}
// 	defer session.EndSession(ctx)
// 	pattern := regexp.MustCompile("(?i)" + name)
// 	topics, err := session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
// 		var topics []*models.Topic

// 		filter := bson.M{}

// 		cursor, err := database.Collection(TOPIC_COLLECTION).Find(ctx, filter)
// 		if err != nil {
// 			logs.CtxError(ctx, "(All Search) Error while finding data from the Topic Collection : %v", err)
// 		}

// 		for cursor.Next(ctx) {
// 			var topic models.Topic

// 			var fields TopicFields

// 			err := cursor.Decode(&fields)
// 			if err != nil {
// 				logs.CtxError(ctx, "(All Search) Error from cursor.Decode while getting topic by validity: %v", err)
// 				return nil, err
// 			}

// 			if fields.TopicSettings == nil || !pattern.MatchString(fields.TopicSettings.Title) {
// 				continue
// 			}
// 			topic.Id = fields.Id
// 			topic.ContentType = *fields.ContentType
// 			if fields.TopicSettings != nil {
// 				topic.TopicSettings = &models.TopicSettings{
// 					Title: fields.TopicSettings.Title,
// 				}
// 			}

// 			topics = append(topics, &topic)
// 		}
// 		cursor.Close(ctx)
// 		return topics, nil
// 	})
// 	if err != nil {
// 		return nil, nil, nil, err
// 	}

// 	profiles, err := session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
// 		var profiles []*models.PublicProfile

// 		//TODO include the regex here instead of fitlering after retrieving everything from DB
// 		filter := bson.M{}

// 		cursor, err := database.Collection(PROFILE_COLLECTION).Find(ctx, filter, options.Find())
// 		if err != nil {
// 			logs.CtxError(ctx, "(All Search) Error while finding data from the Profile Collection: %v", err)
// 		}

// 		for cursor.Next(ctx) {
// 			var profile models.PublicProfile

// 			type ProfileFields struct {
// 				Id              string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" form:"id" query:"id"`
// 				ProfileLinkedin *models.ProfileLinkedIn `protobuf:"bytes,3,opt,name=profile_linkedin,json=profileLinkedin,proto3" json:"profile_linkedin,omitempty" form:"profile_linkedin" query:"profile_linkedin"`
// 			}

// 			var fields ProfileFields

// 			err := cursor.Decode(&fields)
// 			if err != nil {
// 				logs.CtxError(ctx, "(All Search) Error from cursor.Decode while getting profile by validity: %v", err)
// 				return nil, err
// 			}
// 			fullName := fields.ProfileLinkedin.LinkedinFirstName + " " + fields.ProfileLinkedin.LinkedinLastName
// 			if !pattern.MatchString(fullName) {
// 				continue
// 			}

// 			profile.Id = fields.Id
// 			if fields.ProfileLinkedin != nil {
// 				profile.ProfileLinkedin = &models.ProfileLinkedIn{
// 					LinkedinFirstName:  fields.ProfileLinkedin.LinkedinFirstName,
// 					LinkedinLastName:   fields.ProfileLinkedin.LinkedinLastName,
// 					LinkedinProfileUrl: fields.ProfileLinkedin.LinkedinProfileUrl,
// 				}
// 			}
// 			profiles = append(profiles, &profile)
// 		}
// 		cursor.Close(ctx)
// 		return profiles, nil
// 	})
// 	if err != nil {
// 		return nil, nil, nil, err
// 	}

// 	events, err := session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
// 		var events []*models.Event

// 		filter := bson.M{
// 			"title": bson.M{
// 				"$regex":   name,
// 				"$options": "i",
// 			},
// 		}

// 		cursor, err := database.Collection(EVENT_COLLECTION).Find(ctx, filter)
// 		if err != nil {
// 			logs.CtxError(ctx, "(All Search) Error while finding daata from the Event Collection %v", err)
// 		}

// 		for cursor.Next(ctx) {
// 			var event models.Event
// 			var fields EventFields

// 			err := cursor.Decode(&fields)
// 			if err != nil {
// 				logs.CtxError(ctx, "(All Search) Error from cursor.Decode while getting event by validity: %v", err)
// 				return nil, err
// 			}
// 			event.Title = fields.Title
// 			event.Id = fields.Id
// 			events = append(events, &event)
// 		}
// 		cursor.Close(ctx)
// 		return events, nil
// 	})
// 	if err != nil {
// 		return nil, nil, nil, err
// 	}

// 	return topics.([]*models.Topic), events.([]*models.Event), profiles.([]*models.PublicProfile), nil
// }

// func SearchProfile(ctx context.Context, name string, page int) ([]*models.PublicProfile, error) {
// 	if name == "" {
// 		var nullResult []*models.PublicProfile
// 		return nullResult, nil
// 	}

// 	session, err := client.StartSession()
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer session.EndSession(ctx)
// 	filter := bson.M{
// 		"$or": []bson.M{
// 			{"profilelinkedin.linkedinfirstname": bson.M{
// 				"$regex":   name,
// 				"$options": "i",
// 			}},
// 			{"profilelinkedin.linkedinlastname": bson.M{
// 				"$regex":   name,
// 				"$options": "i",
// 			}},
// 		},
// 	}
// 	profiles, err := session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
// 		profiles, err := findProfileWithFilter(ctx, filter, page)
// 		if err != nil {
// 			return nil, err
// 		}
// 		return profiles, nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return profiles.([]*models.PublicProfile), nil
// }

// func SearchTopic(ctx context.Context, name string, page int) ([]*models.Topic, error) {
// 	if name == "" {
// 		var nullResult []*models.Topic
// 		return nullResult, nil
// 	}
// 	session, err := client.StartSession()
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer session.EndSession(ctx)

// 	filter := bson.M{"topicsettings.title": bson.M{
// 		"$regex":   name,
// 		"$options": "i",
// 	}}

// 	topics, err := session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
// 		topics, err := findTopicWithFilter(ctx, filter, page)
// 		if err != nil {
// 			return nil, err
// 		}
// 		return topics, nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return topics.([]*models.Topic), nil
// }

// func SearchEvent(ctx context.Context, name string, page int) ([]*models.Event, error) {
// 	if name == "" {
// 		var nullResult []*models.Event
// 		return nullResult, nil
// 	}
// 	session, err := client.StartSession()
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer session.EndSession(ctx)
// 	filter := bson.M{
// 		"title": bson.M{
// 			"$regex":   name,
// 			"$options": "i",
// 		},
// 	}
// 	events, err := session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
// 		events, err := findEventWithFilter(ctx, filter, page)
// 		if err != nil {
// 			return nil, err
// 		}
// 		return events, nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return events.([]*models.Event), nil
// }

// func SearchTags(ctx context.Context, tagName string, page int) ([]*models.Topic, []*models.Event, []*models.PublicProfile, error) {
// 	session, err := client.StartSession()
// 	if err != nil {
// 		return nil, nil, nil, err
// 	}

// 	filter := bson.M{"topicsettings.tags": bson.M{"$in": []interface{}{
// 		primitive.Regex{Pattern: tagName, Options: "i"},
// 	}}}

// 	defer session.EndSession(ctx)
// 	topics, err := session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
// 		topics, err := findTopicWithFilter(ctx, filter, page)
// 		if err != nil {
// 			return nil, err
// 		}
// 		return topics, nil
// 	})

// 	if err != nil {
// 		return nil, nil, nil, err
// 	}

// 	filter = bson.M{"$or": []bson.M{
// 		{
// 			"interests": bson.M{
// 				"$in": []interface{}{
// 					primitive.Regex{
// 						Pattern: tagName, Options: "i",
// 					},
// 				},
// 			},
// 		},
// 		{
// 			"modules": bson.M{
// 				"$in": []interface{}{
// 					primitive.Regex{
// 						Pattern: tagName, Options: "i",
// 					},
// 				},
// 			},
// 		},
// 		{
// 			"hobbies": bson.M{
// 				"$in": []interface{}{
// 					primitive.Regex{
// 						Pattern: tagName, Options: "i",
// 					},
// 				},
// 			},
// 		},
// 	},
// 	}
// 	profiles, err := session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
// 		profiles, err := findProfileWithFilter(ctx, filter, page)
// 		if err != nil {
// 			return nil, err
// 		}
// 		return profiles, nil
// 	})
// 	if err != nil {
// 		return nil, nil, nil, err
// 	}
// 	filter = bson.M{"tags": bson.M{
// 		"$in": []interface{}{
// 			primitive.Regex{Pattern: tagName, Options: "i"},
// 		},
// 	}}

// 	events, err := session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
// 		events, err := findEventWithFilter(ctx, filter, page)
// 		if err != nil {
// 			return nil, err
// 		}
// 		return events, nil
// 	})
// 	if err != nil {
// 		return nil, nil, nil, err
// 	}

// 	return topics.([]*models.Topic), events.([]*models.Event), profiles.([]*models.PublicProfile), nil
// }

// func findProfileWithFilter(ctx context.Context, filter bson.M, page int) ([]*models.PublicProfile, error) {
// 	var profiles []*models.PublicProfile

// 	opts := options.Find()

// 	itemsPerPage := 4
// 	skipCount := (page - 1) * itemsPerPage
// 	opts.SetSkip(int64(skipCount))
// 	opts.SetLimit(int64(itemsPerPage))

// 	cursor, err := database.Collection(PROFILE_COLLECTION).Find(ctx, filter, opts)
// 	if err != nil {
// 		logs.CtxError(ctx, "Error while finding data from the Profile Collection: %v", err)
// 	}

// 	for cursor.Next(ctx) {
// 		var profile models.PublicProfile
// 		var fields ProfileFields

// 		err := cursor.Decode(&fields)
// 		if err != nil {
// 			logs.CtxError(ctx, "Error from cursor.Decode while getting profile by validity: %v", err)
// 			return nil, err
// 		}

// 		profile.Id = fields.Id
// 		profile.Interests = fields.Interests
// 		profile.Header = fields.Header
// 		profile.Introduction = fields.Introduction
// 		if fields.ProfileLinkedin != nil {
// 			profile.ProfileLinkedin = &models.ProfileLinkedIn{
// 				LinkedinId:         fields.ProfileLinkedin.Id,
// 				LinkedinFirstName:  fields.ProfileLinkedin.LinkedinFirstName,
// 				LinkedinLastName:   fields.ProfileLinkedin.LinkedinLastName,
// 				LinkedinProfileUrl: fields.ProfileLinkedin.LinkedinProfileUrl,
// 				LinkedinHeadline:   fields.ProfileLinkedin.LinkedinHeadline,
// 			}
// 		}
// 		profiles = append(profiles, &profile)
// 	}
// 	cursor.Close(ctx)
// 	return profiles, nil
// }

// func findTopicWithFilter(ctx context.Context, filter bson.M, page int) ([]*models.Topic, error) {
// 	var topics []*models.Topic

// 	opts := options.Find()

// 	itemsPerPage := 4
// 	skipCount := (page - 1) * itemsPerPage
// 	opts.SetSkip(int64(skipCount))
// 	opts.SetLimit(int64(itemsPerPage))

// 	cursor, err := database.Collection(TOPIC_COLLECTION).Find(ctx, filter, opts)
// 	if err != nil {
// 		logs.CtxError(ctx, "Error while finding data from the Topic Collection: %v", err)
// 	}

// 	for cursor.Next(ctx) {
// 		var topic models.Topic
// 		var fields TopicFields

// 		err := cursor.Decode(&fields)
// 		if err != nil {
// 			logs.CtxError(ctx, "Error from cursor.Decode while getting topic by validity: %v", err)
// 			return nil, err
// 		}

// 		topic.Id = fields.Id
// 		topic.ContentType = *fields.ContentType
// 		topic.Participants = fields.Participants
// 		if fields.TopicSettings != nil {
// 			topic.TopicSettings = &models.TopicSettings{
// 				Title: fields.TopicSettings.Title,
// 				Tags:  fields.TopicSettings.Tags,
// 			}
// 		}

// 		topics = append(topics, &topic)
// 	}
// 	cursor.Close(ctx)
// 	return topics, nil
// }

// func findEventWithFilter(ctx context.Context, filter bson.M, page int) ([]*models.Event, error) {
// 	var events []*models.Event

// 	opts := options.Find()

// 	itemsPerPage := 4
// 	skipCount := (page - 1) * itemsPerPage
// 	opts.SetSkip(int64(skipCount))
// 	opts.SetLimit(int64(itemsPerPage))

// 	cursor, err := database.Collection(EVENT_COLLECTION).Find(ctx, filter, opts)
// 	if err != nil {
// 		logs.CtxError(ctx, "Error while finding data from the Event Collection: %v", err)
// 	}

// 	for cursor.Next(ctx) {
// 		var event models.Event

// 		var fields EventFields

// 		err := cursor.Decode(&fields)
// 		if err != nil {
// 			logs.CtxError(ctx, "Error from cursor.Decode while getting event by validity: %v", err)
// 			return nil, err
// 		}
// 		event.Title = fields.Title
// 		event.Id = fields.Id
// 		event.CompanyPictureUrl = fields.CompanyPictureUrl
// 		event.CompanyName = fields.CompanyName
// 		event.Tags = fields.Tags
// 		events = append(events, &event)
// 	}
// 	cursor.Close(ctx)
// 	return events, nil
// }
