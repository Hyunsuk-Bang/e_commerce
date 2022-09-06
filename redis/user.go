package redisDB

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type User struct {
	Username     string `redis:"username"`
	Password     string `redis:"password"`
	SessionToken string `redis:"sessionToken"`
}

func userKey(userID string) string {
	return fmt.Sprintf("users#%s", userID)
}
func usernameUniqueKey() string {
	return "usernames:unique"
}

func isUniqueUserName(ctx context.Context, c *redis.Client, userName string) bool {
	isDuplicate, err := c.SIsMember(ctx, usernameUniqueKey(), userName).Result()
	if err != nil {
		log.Fatal(err)
	}
	return !isDuplicate
}

func (u *User) serialize() map[string]interface{} {
	return map[string]interface{}{
		"username":     u.Username,
		"password":     u.Password,
		"sessionToken": uuid.NewString(),
	}
}
func CreateUser(ctx context.Context, c *redis.Client, u *User) (string, error) {
	if !isUniqueUserName(ctx, c, u.Username) {
		return "", errors.New("given user name is already exsists")
	}
	id := uuid.NewString()
	c.HSet(ctx, userKey(id), u.serialize())
	c.SAdd(ctx, usernameUniqueKey(), u.Username)
	return id, nil
}

func GetUserById(ctx context.Context, c *redis.Client, id string) User {
	var user User
	if err := c.HGetAll(ctx, userKey(id)).Scan(&user); err != nil {
		log.Fatal(err)
	}
	return user
}
