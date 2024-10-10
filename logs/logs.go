package logs

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"gitlab.com/aic/aic_api/consts"
)

func LogWithContext(ctx context.Context) *logrus.Entry {
	return logrus.WithFields(logrus.Fields{
		string(consts.LogID): ctx.Value(consts.LogID),
	})
}

func CtxInfo(ctx context.Context, format string, args ...interface{}) {
	if len(args) == 0 {
		LogWithContext(ctx).Info(format)
		return
	}
	log := fmt.Sprintf(format, args...)
	LogWithContext(ctx).Info(log)
}

func CtxError(ctx context.Context, format string, args ...interface{}) {
	if len(args) == 0 {
		LogWithContext(ctx).Error(format)
		return
	}
	log := fmt.Sprintf(format, args...)
	LogWithContext(ctx).Error(log)
}

func CtxWarn(ctx context.Context, format string, args ...interface{}) {
	if len(args) == 0 {
		LogWithContext(ctx).Warn(format)
		return
	}
	log := fmt.Sprintf(format, args...)
	LogWithContext(ctx).Warn(log)
}
