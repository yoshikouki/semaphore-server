package model

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type Semaphore struct {
	key  string
	user string
	ttl  time.Duration
}

func (m *Model) Lock(ctx context.Context, lockTarget, user string, ttl time.Duration) error {
	isLocked, err := m.redis.SetNX(ctx, lockTarget, user, ttl).Result()
	if err != nil {
		return err
	}

	remainingTTL, err := m.redis.TTL(ctx, lockTarget).Result()
	if err != nil {
		return err
	}

	expireDate := time.Now().Add(remainingTTL)

	if !isLocked {
		lockedUser, err := m.redis.Get(ctx, lockTarget).Result()
		if err != nil {
			return err
		}

		if user == lockedUser {
			return fmt.Errorf("%s is already locked. Expire: %s", lockTarget, expireDate)
		} else {
			return fmt.Errorf("%s is locked by %s. Expire: %s", lockTarget, user, expireDate)
		}
	}

	return nil
}

func (m *Model) Unlock(ctx context.Context, target string, user string) error {
	lockedUser, err := m.redis.Get(ctx, target).Result()
	if err == redis.Nil {
		return fmt.Errorf("%s haven't locked", target)
	} else if err != nil {
		return err
	}

	if user != lockedUser {
		return fmt.Errorf("%s don't release lock, because lock owner isn't %s", target, user)
	}

	_, err = m.redis.Del(ctx, target).Result()
	if err != nil {
		return err
	}

	return nil
}
