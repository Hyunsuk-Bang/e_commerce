package redisDB

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func userLikesKey(userId string) string {
	return fmt.Sprintf("users:likes#%s", userId)
}

func userLikesitem(ctx context.Context, c *redis.Client, userId string, itemId string) (bool, error) {
	return c.SIsMember(ctx, userLikesKey(userId), itemId).Result()
}

func likeItem(ctx context.Context, c *redis.Client, userId string, itemId string) {
	_, err := c.SAdd(ctx, userLikesKey(userId), itemId).Result()
	if err != nil {
		log.Fatal(err)
	}
	c.HIncrBy(ctx, ItemLikesKey(itemId), "likes", 1)
}

func unlikeItem(ctx context.Context, c *redis.Client, userId string, itemId string) {
	_, err := c.SRem(ctx, userLikesKey(userId), itemId).Result()
	if err != nil {
		log.Fatal(err)
	}

	if GetItemLikes(ctx, c, itemId) == 0 {
		c.HSet(ctx, ItemLikesKey(itemId), "likes", "0")
	}
	c.HIncrBy(ctx, ItemLikesKey(itemId), "likes", -1)
}

func likedItem(ctx context.Context, c *redis.Client, userId string) []string {
	likedItems, err := c.SMembers(ctx, userLikesKey(userId)).Result()
	if err != nil {
		log.Fatal(err)
	}
	return likedItems
}
