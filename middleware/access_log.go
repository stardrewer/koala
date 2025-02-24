package middleware

import (
	"context"
	"time"

	"github.com/ibinarytree/koala/logs"
	"github.com/ibinarytree/koala/meta"
	"google.golang.org/grpc/status"
)

func AccessLogMiddleware(next MiddlewareFunc) MiddlewareFunc {
	return func(ctx context.Context, req interface{}) (resp interface{}, err error) {

		startTime := time.Now()
		resp, err = next(ctx, req)

		serverMeta := meta.GetServerMeta(ctx)
		errStatus, _ := status.FromError(err)

		cost := time.Since(startTime).Nanoseconds() / 1000
		logs.AddField(ctx, "cost_us", cost)
		logs.AddField(ctx, "method", serverMeta.Method)

		logs.AddField(ctx, "cluster", serverMeta.Cluster)
		logs.AddField(ctx, "env", serverMeta.Env)
		logs.AddField(ctx, "server_ip", serverMeta.ServerIP)
		logs.AddField(ctx, "client_ip", serverMeta.ClientIP)
		logs.AddField(ctx, "idc", serverMeta.IDC)
		logs.Access(ctx, "result=%v", errStatus.Code())

		return
	}
}
