package store

import (
	"github.com/go-redis/cache"
	"github.com/go-redis/redis"
	"github.com/vmihailenco/msgpack"
)

type Store struct {
	codec *cache.Codec
}

func NewStore(redisClient *redis.Client) *Store {
	codec := &cache.Codec{
		Redis: redisClient,

		Marshal: msgpack.Marshal,

		Unmarshal: msgpack.Unmarshal,
	}
	return &Store{codec: codec}
}
