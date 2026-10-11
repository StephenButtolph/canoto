//go:build !race

package canoto

import "sync/atomic"

// CacheUint64 atomically stores val into *addr.
func CacheUint64(addr *uint64, val uint64) {
	atomic.StoreUint64(addr, val)
}

// CacheUint32 atomically stores val into *addr.
func CacheUint32(addr *uint32, val uint32) {
	atomic.StoreUint32(addr, val)
}
