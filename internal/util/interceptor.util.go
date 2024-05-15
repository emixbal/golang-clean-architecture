package util

import (
	"context"
	"golang-clean-architecture/domain"
	"time"
)

func ResponseInterceptor(ctx context.Context, resp *domain.ApiResponse) {
	resp.Timestamp = time.Now()
	resp.TraceID = ""
}
