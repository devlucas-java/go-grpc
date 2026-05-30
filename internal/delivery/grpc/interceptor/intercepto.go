package interceptor

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/time/rate"
	"google.golang.org/grpc"
)

func LogIntercepto(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	log.Printf("grpc method: %s, request: %v", info.FullMethod, req)
	resp, err := handler(ctx, req)
	if err != nil {
		log.Printf("grpc method: %s, error: %v", info.FullMethod, err)
	}

	return resp, err

}

var limit = rate.NewLimiter(rate.Every(1*time.Second), 1)

func RateLimiter(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	if !limit.Allow() {
		return nil, fmt.Errorf("error rate limit execed")
	}

	return handler(ctx, req)
}
