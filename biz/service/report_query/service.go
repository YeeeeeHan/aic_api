package report_query

// import (
// 	"context"

// 	"gitlab.com/aic/aic_api/biz/auth"
// 	"gitlab.com/aic/aic_api/biz/dal"
// 	"gitlab.com/aic/aic_api/biz/model/gitlab.com/aic/http_idl_gen/gen/aic/data/models"
// 	"gitlab.com/aic/aic_api/biz/util/errors"
// 	"gitlab.com/aic/aic_api/biz/util/helpers"
// )

// type ReportQueryService interface {
// 	UpdateData(ctx context.Context, params *models.ReportRequestParams) (*models.ReportResponseData, error)
// 	Validate(ctx context.Context, params *models.ReportRequestParams) error
// }

// type ReportQueryServiceImpl struct{}

// func NewReportQueryService() ReportQueryService {
// 	return &ReportQueryServiceImpl{}
// }

// func (s *ReportQueryServiceImpl) UpdateData(ctx context.Context, params *models.ReportRequestParams) (*models.ReportResponseData, error) {
// 	params.Report.Id = helpers.GenerateID()
// 	err := dal.ReportUser(ctx, params.GetReport())

// 	if err != nil {
// 		return nil, err
// 	}
// 	return &models.ReportResponseData{
// 		Message: "",
// 	}, nil
// }

// func (s *ReportQueryServiceImpl) Validate(ctx context.Context, params *models.ReportRequestParams) error {
// 	if params == nil {
// 		return errors.NewInvalidParamsError("validation error")
// 	}
// 	if ok := auth.VerifyProfileHeader(ctx, params.GetProfileHeader()); !ok {
// 		return errors.NewAuthorisationError("profile header passed is not authorised")
// 	}

// 	switch params.Report.ContentType {
// 	case models.ContentType_CHAT:
// 		if params.Report.ChatId == "" {
// 			return errors.NewInvalidParamsError("missing chat id")
// 		}
// 	case models.ContentType_PROFILE:
// 		if params.Report.ProfileId == "" {
// 			return errors.NewInvalidParamsError("missing profile id")
// 		}
// 	case models.ContentType_OPENED_TOPIC:
// 		if params.Report.TopicId == "" {
// 			return errors.NewInvalidParamsError("missing topic id")
// 		}
// 	case models.ContentType_UNOPENED_TOPIC:
// 		if params.Report.TopicId == "" {
// 			return errors.NewInvalidParamsError("missing topic id")
// 		}
// 	}

// 	return nil
// }
