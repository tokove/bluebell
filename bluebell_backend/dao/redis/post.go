package redis

import (
	"bluebell_backend/model"
	"time"

	"github.com/go-redis/redis"
)

func CreatePost(postID uint64) error {
	pipeline := client.TxPipeline()
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  0,
		Member: postID,
	})
	_, err := pipeline.Exec()
	return err
}

func GetPostIDsInOrder(p *model.ParamPostList) ([]string, error) {
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == model.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	return client.ZRevRange(key, (p.Page-1)*p.Size, p.Page*p.Size-1).Result()
}
