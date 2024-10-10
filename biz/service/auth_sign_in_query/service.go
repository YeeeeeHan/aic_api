package sign_in_query

import (
	"context"

	"gitlab.com/aic/aic_api/biz/model/github.com/aic/http_idl_gen/gen/aic/data/models"
	"gitlab.com/aic/aic_api/biz/util/errors"
)

type SignInQueryService interface {
	GetData(ctx context.Context, params *models.SignInRequestParams) (*models.SignInResponseData, error)
	Validate(ctx context.Context, params *models.SignInRequestParams) error
}

type SignInQueryServiceImpl struct{}

func NewSignInQueryService() SignInQueryService {
	return &SignInQueryServiceImpl{}
}

func (s *SignInQueryServiceImpl) GetData(ctx context.Context, params *models.SignInRequestParams) (*models.SignInResponseData, error) {

	// token, err := auth.OauthConf.Exchange(ctx, params.GetCode())
	// if err != nil {
	// 	return nil, err
	// }
	// accessToken := token.AccessToken

	// client := auth.OauthConf.Client(ctx, token)
	// // Getting data of name and profile picture
	// response, err := client.Get("https://api.linkedin.com/v2/me")
	// if err != nil {
	// 	return nil, err
	// }
	// defer response.Body.Close()

	// content, err := ioutil.ReadAll(response.Body)

	// if err != nil {
	// 	return nil, err
	// }

	// var data models.LinkedInData
	// err = json.Unmarshal(content, &data)
	// if err != nil {
	// 	return nil, err
	// }

	// profile, err := dal.CheckProfileByLinkedInId(ctx, data.Id)

	// var (
	// 	id          string
	// 	linkedInId  string
	// 	newUserBool bool
	// )

	// if err != nil { // New User

	// 	newProfile, newId, err := dal.RegisterNewUser(ctx, &data, accessToken)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	profile = newProfile
	// 	id = newId
	// 	linkedInId = newProfile.ProfileLinkedin.GetLinkedinId()
	// 	newUserBool = true
	// } else {
	// 	linkedInId = profile.ProfileLinkedin.GetLinkedinId()
	// 	id = profile.GetId()
	// 	newUserBool = false
	// }

	// jWT := auth.CreateJWT(profile.GetId(), data.GetId(), accessToken)
	// err = redis.RedisClient.Set(ctx, jWT, id, 730*time.Hour).Err()
	// if err != nil {
	// 	return nil, err
	// }

	// return &models.SignInResponseData{
	// 	Jwttoken:   jWT,
	// 	LinkedinId: linkedInId,
	// 	Id:         id,
	// 	IsNewUser:  newUserBool,
	// }, nil

	return &models.SignInResponseData{}, nil
}

func (s *SignInQueryServiceImpl) Validate(ctx context.Context, params *models.SignInRequestParams) error {
	if params == nil || params.GetCode() == "" {
		return errors.NewInvalidParamsError("validation error")
	}
	return nil
}
