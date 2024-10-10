package profiles_search_query

// import (
// 	"context"

// 	"gitlab.com/aic/aic_api/biz/dal"
// 	"gitlab.com/aic/aic_api/biz/model/gitlab.com/aic/http_idl_gen/gen/aic/data/models"
// 	"gitlab.com/aic/aic_api/biz/util/errors"
// 	"gitlab.com/aic/aic_api/logs"
// )

// type ProfilesSearchQueryService interface {
// 	GetData(ctx context.Context, params *models.ProfilesSearchRequestParams) (*models.ProfilesSearchResponseData, error)
// 	Validate(ctx context.Context, params *models.ProfilesSearchRequestParams) error
// 	Filter(ctx context.Context, profile *models.Profile)
// }

// type ProfilesSearchQueryServiceImpl struct{}

// func NewProfilesSearchQueryService() ProfilesSearchQueryService {
// 	return &ProfilesSearchQueryServiceImpl{}
// }

// func (s *ProfilesSearchQueryServiceImpl) GetData(ctx context.Context, params *models.ProfilesSearchRequestParams) (*models.ProfilesSearchResponseData, error) {
// 	id := params.GetId()
// 	logs.CtxInfo(ctx, "getting data from other profile")
// 	profile, err := dal.GetProfileById(ctx, id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	s.Filter(ctx, profile)
// 	return &models.ProfilesSearchResponseData{
// 		Profile: profile,
// 	}, err
// }

// func (s *ProfilesSearchQueryServiceImpl) Validate(ctx context.Context, params *models.ProfilesSearchRequestParams) error {
// 	if params == nil || params.GetId() == "" {
// 		return errors.NewInvalidParamsError("validation error")
// 	}
// 	return nil
// }

// func (s *ProfilesSearchQueryServiceImpl) Filter(ctx context.Context, profile *models.Profile) {
// 	profile.Email = ""
// 	profile.ProfilePretexts = nil
// 	profile.Telegram = ""
// 	profile.AccountType = ""
// }
