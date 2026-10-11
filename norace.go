//go:build !race

package canoto

import "sync/atomic"

// StoreUint64 atomically stores val into *addr; see [atomic.StoreUint64].
func StoreUint64(addr *uint64, val uint64) {
	atomic.StoreUint64(addr, val)
}

// StoreUint32 atomically stores val into *addr; see [atomic.StoreUint32].
func StoreUint32(addr *uint32, val uint32) {
	atomic.StoreUint32(addr, val)
}
