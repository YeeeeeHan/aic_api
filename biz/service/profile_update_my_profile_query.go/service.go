package update_my_profile_query

// import (
// 	"context"

// 	"gitlab.com/aic/aic_api/biz/auth"
// 	"gitlab.com/aic/aic_api/biz/dal"
// 	"gitlab.com/aic/aic_api/biz/model/gitlab.com/aic/http_idl_gen/gen/aic/data/models"
// 	"gitlab.com/aic/aic_api/biz/util/errors"
// )

// type UpdateMyProfileQueryService interface {
// 	GetData(ctx context.Context, params *models.UpdateMyProfileRequestParams) (*models.UpdateMyProfileResponseData, error)
// 	Validate(ctx context.Context, params *models.UpdateMyProfileRequestParams) error
// }

// type UpdateMyProfileQueryServiceImpl struct{}

// func NewUpdateMyProfileQueryService() UpdateMyProfileQueryService {
// 	return &UpdateMyProfileQueryServiceImpl{}
// }

// func (s *UpdateMyProfileQueryServiceImpl) GetData(ctx context.Context, params *models.UpdateMyProfileRequestParams) (*models.UpdateMyProfileResponseData, error) {
// 	profileHeader := params.GetProfileHeader()
// 	newProfile := params.GetNewProfile()
// 	err := dal.UpdateProfile(ctx, newProfile, profileHeader)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &models.UpdateMyProfileResponseData{
// 		Message: "",
// 	}, err
// }

// func (s *UpdateMyProfileQueryServiceImpl) Validate(ctx context.Context, params *models.UpdateMyProfileRequestParams) error {
// 	if params == nil || params.GetProfileHeader() == "" {
// 		return errors.NewInvalidParamsError("validation error")
// 	}
// 	if ok := auth.VerifyProfileHeader(ctx, params.GetProfileHeader()); !ok {
// 		return errors.NewAuthorisationError("profile header passed is not authorised")
// 	}
// 	return nil
// }
