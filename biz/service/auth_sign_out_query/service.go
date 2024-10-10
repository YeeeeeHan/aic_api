package sign_out_query

import (
	"context"

	"gitlab.com/aic/aic_api/biz/auth"
	"gitlab.com/aic/aic_api/biz/model/github.com/aic/http_idl_gen/gen/aic/data/models"
	"gitlab.com/aic/aic_api/biz/redis"
	"gitlab.com/aic/aic_api/biz/util/errors"
)

type SignOutQueryService interface {
	GetData(ctx context.Context, params *models.SignOutRequestParams) (*models.SignOutResponseData, error)
	Validate(ctx context.Context, params *models.SignOutRequestParams) error
}

type SignOutQueryServiceImpl struct{}

func NewSignOutQueryService() SignOutQueryService {
	return &SignOutQueryServiceImpl{}
}

func (s *SignOutQueryServiceImpl) GetData(ctx context.Context, params *models.SignOutRequestParams) (*models.SignOutResponseData, error) {
	authHeader := params.GetProfileHeader()
	redis.RedisClient.Del(ctx, authHeader)
	return &models.SignOutResponseData{
		Message: "success",
	}, nil
}

func (s *SignOutQueryServiceImpl) Validate(ctx context.Context, params *models.SignOutRequestParams) error {
	if params == nil || params.GetProfileHeader() == "" {
		return errors.NewInvalidParamsError("validation error")
	}
	if ok := auth.VerifyProfileHeader(ctx, params.GetProfileHeader()); !ok {
		return errors.NewAuthorisationError("profile header passed is not authorised")
	}
	return nil
}
