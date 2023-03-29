package memcached

import "github.com/bradfitz/gomemcache/memcache"

type memcacheClient struct {
	client *memcache.Client
}

func NewMemcacheClient(servers []string) *memcacheClient {
	return &memcacheClient{
		client: memcache.New(servers...),
	}
}

func (c *memcacheClient) Get(key string) (string, error) {
	item, err := c.client.Get(key)
	if err != nil {
		return "", err
	}
	return string(item.Value), nil
}

func (c *memcacheClient) Set(key string, value string) error {
	return c.client.Set(&memcache.Item{
		Key:   key,
		Value: []byte(value),
	})
}
