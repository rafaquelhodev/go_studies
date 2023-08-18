package cache

type InMemoryCache struct {
	Data map[string]string
}

func (c *InMemoryCache) get(key string) (string, bool) {
	value, found := c.Data[key]
	return value, found
}
