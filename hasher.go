package fnv1a

import (
	"math"
)

const (
	zeroHash   = Hash(2166136261)
	multiplier = 16777619
	mask       = Hash(0xFF)
)

// NewHash initializes a new Hash (it is not recommended to use the zero value of the Hash)
func NewHash() Hash {
	return zeroHash
}

// Hash provides the set of functions needed to compute a running fnv1a of the supported rest.li primitives
type Hash uint32

func (h *Hash) addUint32(v uint32) {
	hash := *h
	hash ^= Hash(v) & mask
	hash *= multiplier
	hash ^= Hash(v>>8) & mask
	hash *= multiplier
	hash ^= Hash(v>>16) & mask
	hash *= multiplier
	hash ^= Hash(v>>24) & mask
	hash *= multiplier
	*h = hash
}

func (h *Hash) addUint64(v uint64) {
	hash := *h
	hash ^= Hash(v) & mask
	hash *= multiplier
	hash ^= Hash(v>>8) & mask
	hash *= multiplier
	hash ^= Hash(v>>16) & mask
	hash *= multiplier
	hash ^= Hash(v>>24) & mask
	hash *= multiplier
	hash ^= Hash(v>>32) & mask
	hash *= multiplier
	hash ^= Hash(v>>40) & mask
	hash *= multiplier
	hash ^= Hash(v>>48) & mask
	hash *= multiplier
	hash ^= Hash(v>>56) & mask
	hash *= multiplier
	*h = hash
}

// AddByte increments the current hasher with the given byte
func (h *Hash) AddByte(v byte) {
	hash := *h
	hash ^= Hash(v)
	hash *= multiplier
	*h = hash
}

// AddInt16 increments the current hasher with the given int16
func (h *Hash) AddInt16(v int16) {
	hash := *h
	hash ^= Hash(v) & mask
	hash *= multiplier
	hash ^= Hash(v>>8) & mask
	hash *= multiplier
	*h = hash
}

// AddInt32 increments the current hasher with the given int32
func (h *Hash) AddInt32(v int32) {
	h.addUint32(uint32(v))
}

// AddInt64 increments the current hasher with the given int64
func (h *Hash) AddInt64(v int64) {
	h.addUint64(uint64(v))
}

// AddFloat32 increments the current hasher with the given float32
func (h *Hash) AddFloat32(v float32) {
	h.addUint32(math.Float32bits(v))
}

// AddFloat64 increments the current hasher with the given float64
func (h *Hash) AddFloat64(v float64) {
	h.addUint64(math.Float64bits(v))
}

// AddBool increments the current hasher with the given bool
func (h *Hash) AddBool(v bool) {
	hash := *h
	var b Hash
	if v {
		b = 1
	} else {
		b = 0
	}
	hash ^= b
	hash *= multiplier
	*h = hash
}

// AddString increments the current hasher with the given string
func (h *Hash) AddString(v string) {
	h.AddBytes([]byte(v))
}

// AddBytes increments the current hasher with the given byte slice
func (h *Hash) AddBytes(v []byte) {
	hash := *h
	for _, b := range v {
		h.AddByte(b)
	}
	*h = hash
}

// Add increments the current hasher with the value of another hasher
func (h *Hash) Add(other Hash) {
	h.addUint32(uint32(other))
}

// HashByte initializes a Hash with the given byte
func HashByte(v byte) Hash {
	h := NewHash()
	h.AddByte(v)
	return h
}

// HashInt16 initializes a Hash with the given int16
func HashInt16(v int16) Hash {
	h := NewHash()
	h.AddInt16(v)
	return h
}

// Hash HashInt32 initializes a Hash with the given int32
func HashInt32(v int32) Hash {
	h := NewHash()
	h.AddInt32(v)
	return h
}

// Hash HashInt64 initializes a Hash with the given int64
func HashInt64(v int64) Hash {
	h := NewHash()
	h.AddInt64(v)
	return h
}

// HashFloat32 initializes a Hash with the given float32
func HashFloat32(v float32) Hash {
	h := NewHash()
	h.AddFloat32(v)
	return h
}

// HashFloat64 initializes a Hash with the given float64
func HashFloat64(v float64) Hash {
	h := NewHash()
	h.AddFloat64(v)
	return h
}

// HashBool initializes a Hash with the given bool
func HashBool(v bool) Hash {
	h := NewHash()
	h.AddBool(v)
	return h
}

// HashString initializes a Hash with the given string
func HashString(v string) Hash {
	h := NewHash()
	h.AddString(v)
	return h
}

// HashBytes initializes a Hash with the given bytes
func HashBytes(v []byte) Hash {
	h := NewHash()
	h.AddBytes(v)
	return h
}
