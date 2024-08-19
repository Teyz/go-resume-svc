package service_v1

import (
	"context"
	"fmt"
	"time"

	pkg_cache "github.com/teyz/go-svc-template/pkg/cache"
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
	cache pkg_cache.Cache
}

func NewExampleStoreService(ctx context.Context, cache pkg_cache.Cache) (*service, error) {
	return &service{
		cache: cache,
	}, nil
}
