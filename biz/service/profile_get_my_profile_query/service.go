package get_my_profile_query

import (
	"context"

	"gitlab.com/aic/aic_api/biz/auth"
	"gitlab.com/aic/aic_api/biz/model/github.com/aic/http_idl_gen/gen/aic/data/models"
	"gitlab.com/aic/aic_api/biz/util/errors"
)

type GetMyProfileQueryService interface {
	GetData(ctx context.Context, params *models.GetMyProfileRequestParams) (*models.GetMyProfileResponseData, error)
	Validate(ctx context.Context, params *models.GetMyProfileRequestParams) error
}

type GetMyProfileQueryServiceImpl struct{}

func NewGetMyProfileQueryService() GetMyProfileQueryService {
	return &GetMyProfileQueryServiceImpl{}
}

func (s *GetMyProfileQueryServiceImpl) GetData(ctx context.Context, params *models.GetMyProfileRequestParams) (*models.GetMyProfileResponseData, error) {
	// profileHeader := params.GetProfileHeader()
	// profile, err := dal.GetProfileById(ctx, profileHeader)

	// blockedList := dal.GetBlockedListById(ctx, profileHeader)

	// if blockedList != nil {
	// 	cache.AddToCache(ctx, cacheHelper.GenerateCacheKey("blockedlist", profileHeader, cache.Env), *blockedList)
	// }

	// return &models.GetMyProfileResponseData{
	// 	Profile: profile,
	// }, err
	return &models.GetMyProfileResponseData{}, nil
}

func (s *GetMyProfileQueryServiceImpl) Validate(ctx context.Context, params *models.GetMyProfileRequestParams) error {
	if params == nil || params.GetProfileHeader() == "" {
		return errors.NewInvalidParamsError("validation error")
	}
	if ok := auth.VerifyProfileHeader(ctx, params.GetProfileHeader()); !ok {
		return errors.NewAuthorisationError("profile header passed is not authorised")
	}
	if ok := auth.VerifyProfileHeader(ctx, params.GetProfileHeader()); !ok {
		return errors.NewAuthorisationError("profile header passed is not authorised")
	}
	return nil
}
