package redisDB

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/go-redis/redis/v8"
)

type Item struct {
	itemId      string `redis:itemId`
	name        string `redis:"itemName"`
	description string `redis:"description"`
	likes       uint   `redis:"likes"`
}

func ItemLikesKey(itemId string) string {
	return fmt.Sprintf("item#%s", itemId)
}

func GetItemLikes(ctx context.Context, c *redis.Client, itemId string) int64 {
	likes, err := c.HGet(ctx, ItemLikesKey(itemId), "likes").Result()
	if err != nil {
		log.Fatal(err)
	}
	ret, err := strconv.ParseInt(likes, 0, 32)
	if err != nil {
		log.Fatal(err)
	}
	return ret
}
