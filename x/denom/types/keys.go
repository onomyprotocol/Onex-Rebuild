package types

const (
	// ModuleName defines the module name
	ModuleName = "denom"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_denom"

	// Version defines the current version the IBC module supports
	Version = "denom-1"

	// PortID is the default port id that module binds to
	PortID = "denom"
)

var (
	ParamsKey = []byte("p_denom")
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("denom-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
