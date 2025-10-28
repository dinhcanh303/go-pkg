package bloom

import (
	"context"
	"errors"
	"go-pkg/redis"
	"strconv"
)

type redisBitSet struct {
	store redis.Cache
	key   string
	bits  uint
}

func newRedisBitSet(store redis.Cache, key string, bits uint) *redisBitSet {
	return &redisBitSet{
		store: store,
		key:   key,
		bits:  bits,
	}
}

func (r *redisBitSet) buildOffsetArgs(offsets []uint) ([]string, error) {
	args := make([]string, 0, len(offsets))

	for _, offset := range offsets {
		if offset >= r.bits {
			return nil, ErrTooLargeOffset
		}
		args = append(args, strconv.FormatUint(uint64(offset), 10))
	}
	return args, nil
}

func (r *redisBitSet) check(ctx context.Context, offsets []uint) (bool, error) {
	args, err := r.buildOffsetArgs(offsets)
	if err != nil {
		return false, err
	}
	// Execute the Lua script to check bits
	resp, err := r.store.ScriptRun(ctx, getScript, []string{r.key}, args)
	if errors.Is(err, redis.Nil) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	exists, ok := resp.(int64)
	if !ok {
		return false, nil
	}
	return exists == 1, nil

}

func (r *redisBitSet) del(ctx context.Context) error {
	_, err := r.store.Del(ctx, r.key)
	return err
}

func (r *redisBitSet) set(ctx context.Context, offsets []uint) error {
	args, err := r.buildOffsetArgs(offsets)
	if err != nil {
		return err
	}
	// Execute the Lua script to set bits
	_, err = r.store.ScriptRun(ctx, setScript, []string{r.key}, args)
	if errors.Is(err, redis.Nil) {
		return nil
	}

	return err
}

func (r *redisBitSet) expire(ctx context.Context, seconds int) (bool, error) {
	return r.store.Expire(ctx, r.key, seconds)
}
