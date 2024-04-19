package types

const (
	// ModuleName defines the module name
	ModuleName = "market"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_market"
)

var (
	ParamsKey = []byte("p_market")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func UidKey() []byte {
	return []byte("uid_market")
}
