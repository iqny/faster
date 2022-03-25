package service

import (
	"context"
	"github.com/opentracing/opentracing-go"
)

func (s *Service) GetName(ctx context.Context) string{
	addOrderSpan,_ := opentracing.StartSpanFromContext(ctx, "gethane")
	addOrderSpan.SetTag("GetName start", true)
	defer addOrderSpan.Finish()
	return "12312"
}