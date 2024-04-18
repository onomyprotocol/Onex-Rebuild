package types

import (
	"slices"
	"strings"
)

const (
	// PoolKeyPrefix is the prefix to retrieve all Pool
	PoolKeyPrefix = "Pool/value/"
)

// PoolKey returns the store key to retrieve a Pool
func PoolKey(
	denomA string,
	denomB string,
) []byte {

	denoms := []string{denomA, denomB}
	// CoinAmsg and CoinBmsg pre-sort from raw msg

	slices.Sort(denoms)

	var key []byte

	pairBytes := []byte(strings.Join(denoms, ", "))
	key = append(key, pairBytes...)
	key = append(key, []byte("/")...)

	return key
}
