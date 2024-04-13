package types

const (
	// ModuleName defines the module name
	ModuleName = "finale"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_finale"
)

var (
	ParamsKey = []byte("p_finale")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
