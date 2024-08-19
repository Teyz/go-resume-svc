package service_v1

import (
	"context"
	"fmt"
	"time"
)

const (
	exampleCacheDuration = time.Hour * 24
)

func generateExampleCacheKeyWithID(id string) string {
	return fmt.Sprintf("go-svc-template:example:id:%v", id)
}

func generateExamplesCacheKey() string {
	return "go-svc-template:examples"
}

type service struct {
}

func NewExampleStoreService(ctx context.Context) (*service, error) {
	return &service{}, nil
}
