package response

import (
	"context"
	"strconv"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var httpHeaderKey = "x-http-status"
var httpHeaderKeySetOnce sync.Once

func SetHttpHeaderKey(key string) {
	httpHeaderKeySetOnce.Do(func() {
		httpHeaderKey = key
	})
}

func GetHttpHeaderKey() string {
	return httpHeaderKey
}

func SetHeaders(ctx context.Context, headers metadata.MD) error {
	return grpc.SetHeader(ctx, headers)
}

func SetTrailers(ctx context.Context, trailers metadata.MD) error {
	return grpc.SetTrailer(ctx, trailers)
}

func SetHttpStatusHeader(ctx context.Context, code int) error {
	return SetHeaders(
		ctx,
		metadata.Pairs(httpHeaderKey, strconv.Itoa(code)),
	)
}
