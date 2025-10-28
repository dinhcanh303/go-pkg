package bloom

import (
	"context"
	_ "embed"
	"errors"
	"go-pkg/hash"
	"go-pkg/redis"
)

var (
	// ErrTooLargeOffset indicates the offset is too large in bitset.
	ErrTooLargeOffset = errors.New("too large offset")

	//go:embed set_script.lua
	setLuaScript string
	setScript    = redis.NewScript(setLuaScript)

	//go:embed get_script.lua
	getLuaScript string
	getScript    = redis.NewScript(getLuaScript)
)

type Filter struct {
	bitSet         bitSetProvider
	bits           uint
	kHashFunctions uint
}

func NewBloomFilter(store redis.Cache, key string, bits uint, kHashFunctions uint) *Filter {
	return &Filter{
		bits:           bits,
		bitSet:         newRedisBitSet(store, key, bits),
		kHashFunctions: kHashFunctions,
	}
}
func (f *Filter) getLocations(data []byte) []uint {
	locations := make([]uint, f.kHashFunctions)
	for i := uint(0); i < f.kHashFunctions; i++ {
		hashVal := hash.Hash(append(data, byte(i)))
		locations[i] = uint(hashVal % uint64(f.bits))
	}
	return locations
}

func (f *Filter) AddWithCtx(ctx context.Context, data []byte) error {
	locations := f.getLocations(data)
	return f.bitSet.set(ctx, locations)
}

func (f *Filter) Add(data []byte) error {
	return f.AddWithCtx(context.Background(), data)
}

func (f *Filter) ExistsWithCtx(ctx context.Context, data []byte) (bool, error) {
	locations := f.getLocations(data)
	isSet, err := f.bitSet.check(ctx, locations)
	if err != nil {
		return false, err
	}
	return isSet, nil
}

func (f *Filter) Exists(data []byte) (bool, error) {
	return f.ExistsWithCtx(context.Background(), data)
}
