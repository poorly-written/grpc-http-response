package success

import (
	"context"

	"github.com/poorly-written/grpc-http-response"
	"github.com/poorly-written/grpc-http-response/codes"
)

func Send[T any](ctx context.Context, res T, code codes.Code) (T, error) {
	return res, response.SetHttpStatusHeader(ctx, code.HttpCode())
}

func Ok[T any](ctx context.Context, res T) (T, error) {
	return Send(ctx, res, codes.Ok)
}

func Created[T any](ctx context.Context, res T) (T, error) {
	return Send(ctx, res, codes.Created)
}

func Accepted[T any](ctx context.Context, res T) (T, error) {
	return Send(ctx, res, codes.Accepted)
}

func NoContent[T any](ctx context.Context, res T) (T, error) {
	return Send(ctx, res, codes.NoContent)
}
