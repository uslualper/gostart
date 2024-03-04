package cache

import (
	"gostart/pkg/cache/memcached"
	"gostart/pkg/config"
)

var client Client

type Client interface {
	Get(key string) (string, error)
	Set(key string, value string) error
}

func SetupCache() {
	client = memcached.NewMemcacheClient([]string{config.GetString("MEMCACHED_HOST") + ":" + config.GetString("MEMCACHED_PORT")})
}

func Cache() Client {
	return client
}
