package types

const (
	// ModuleName defines the module name
	ModuleName = "shop"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_shop"
)

var (
	ParamsKey = []byte("p_shop")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	PostKey = "Post/value/"
)

const (
	PostCountKey = "Post/count/"
)
