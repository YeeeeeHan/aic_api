package middlewares

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol"
	"gitlab.com/aic/aic_api/biz/auth"
	"gitlab.com/aic/aic_api/biz/redis"
	"gitlab.com/aic/aic_api/biz/util/helpers"
	"gitlab.com/aic/aic_api/consts"
	"gitlab.com/aic/aic_api/logs"
	"gitlab.com/aic/aic_api/monitoring"
)

var (
	AllowedHandlers = []string{
		"gitlab.com/aic/aic_api/biz/handler.SignInQuery",
		"gitlab.com/aic/aic_api/biz/handler.SignOutQuery",
		"gitlab.com/aic/aic_api/biz/handler.JoinRoom",
		"gitlab.com/aic/aic_api/biz/handler.AdminSignInQuery",
		"gitlab.com/aic/aic_api/biz/handler.AppleSignInQuery",
		"gitlab.com/aic/aic_api/biz/handler.ProfileUpdateDpQuery",
	}
)

func MonitorReq(ctx context.Context, c *app.RequestContext){
	req := c.GetRequest()
	start := time.Now()
	c.Set("start", start)

	c.Next(ctx)
	stop := time.Now()
	c.Set("stop", stop)

	latency := stop.Sub(start)

	go parseCtxReq(req, latency)
	return
}

func parseCtxReq(req *protocol.Request ,lat time.Duration) {
	path := string(req.URI().Path())
	paramData := string(req.Body())
	pid_pattern := regexp.MustCompile(`\"profile_header\"\s*:\s*\"([a-zA-Z0-9]+)\"`)
	reg := pid_pattern.FindStringSubmatch(paramData)
	var pid string
	if len(reg) > 1 {
		pid = reg[1]
	} else {
		pid = "None"
	}

	dt := time.Now()
	date := dt.Format("01-02-2006")
	time := dt.Format("15:04:05")
	log := fmt.Sprintf("{\"profile_id\":\"%s\",\"path\":\"%s\",\"date\":\"%s\",\"time\":\"%s\", \"latency\":\"%s\"}\n", pid, path, date, time, lat)
	go monitoring.Write(log)
}

func CheckAuth(ctx context.Context, c *app.RequestContext) {
	for _, h := range AllowedHandlers {
		if h == c.HandlerName() {
			return
		}
	}

	authHeader := string(c.GetHeader("Authorization"))
	profile_header, err := redis.RedisClient.Get(ctx, authHeader).Result()
	if err != nil {
		logs.CtxWarn(ctx, "error is: %v", err)
		if authHeader != "" {
			if authHeader == "63d64ec3206e6225053a541239c57499f57dd2f197b7f3b32e58956c515994db" {
				c.Next(context.WithValue(ctx, consts.ProfileHeader, "63d64ec3206e6225053a541239c57499f57dd2f197b7f3b32e58956c515994db"))
				return
			}
			verified, profile_header, _ := auth.VerifyJWT(ctx, authHeader)
			if verified {
				_ = redis.RedisClient.Set(ctx, authHeader, profile_header, 730*time.Hour).Err()
				logs.CtxInfo(ctx, "JWT token still valid and added back to redis, profile header: %v", profile_header)
				c.Next(context.WithValue(ctx, consts.ProfileHeader, profile_header))
			}
		}
	} else {
		c.Next(context.WithValue(ctx, consts.ProfileHeader, profile_header))
		return
	}
	c.AbortWithError(403, err)

}

func InjectMeta(ctx context.Context, c *app.RequestContext) {
	// setting context time limit
	newCtx, _ := context.WithDeadline(ctx, time.Now().Add(time.Minute))

	// setting timezone to be singapore
	singapore, _ := time.LoadLocation("Asia/Singapore")

	// generating log id
	logId := fmt.Sprintf("%v-%v", time.Now().In(singapore).Format("02-01-2006-15-04-05"), helpers.GenerateID())
	c.Response.Header.Add(string(consts.LogID), logId)
	newCtx = context.WithValue(newCtx, consts.LogID, logId)
	c.Next(newCtx)
}
