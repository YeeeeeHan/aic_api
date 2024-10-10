// Code generated by hertz generator.

package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	aic_api "gitlab.com/aic/aic_api/biz/model/aic_api"
)

// UploadQuestionsQuery .
// @router /api/v:version/aic/questions/upload [POST]
func UploadQuestionsQuery(ctx context.Context, c *app.RequestContext) {
	var err error
	var req aic_api.UploadQuestionsRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(aic_api.UploadQuestionsResponse)

	c.JSON(consts.StatusOK, resp)
}
