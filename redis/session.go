package redisDB

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

const Layout = "2006-01-02 15:04:05 (PST)"

type Session struct {
	UserId   string `redis:"userId"`
	Username string `redis:"username"`
	Expiry   string `redis:"expiry"`
}

func pageCacheKey(id string) string {
	return fmt.Sprintf("pagecache#%s", id)
}
func sessionsKey(id string) string {
	return fmt.Sprintf("sessions#%s", id)
}

func (s *Session) serialize() map[string]interface{} {
	return map[string]interface{}{
		"userId":   s.UserId,
		"username": s.Username,
		"expiry":   time.Now().Add(120 * time.Second).Format(Layout),
	}
}
func SaveSession(ctx context.Context, c *redis.Client, sessionToken string, s *Session) {
	c.HSet(ctx, sessionsKey(sessionToken), s.serialize())
}

func GetSession(ctx context.Context, c *redis.Client, id string) Session {
	var session Session
	err := c.HGetAll(ctx, sessionsKey(id)).Scan(&session)
	if err != nil {
		log.Fatal(err)
	}
	return session
}
