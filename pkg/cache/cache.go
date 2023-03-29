package cache

import (
	"go-start/pkg/cache/memcached"
	"go-start/pkg/config"
)

var client Client

type Client interface {
	Get(key string) (string, error)
	Set(key string, value string) error
}

func SetupCache() {
	client = memcached.NewMemcacheClient([]string{config.Config("MEMCACHED_HOST") + ":" + config.Config("MEMCACHED_PORT")})
}

func Cache() Client {
	return client
}
