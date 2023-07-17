package BloomFilter

import (
	"hash/fnv"
	"reflect"
	"strconv"
	"sync"
)

type Hashable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

type BloomFilter[T Hashable] struct {
	mu     sync.RWMutex
	bitmap []byte
	count  int
}

// Create a new Bloom Filter
func NewBloomFilter[T Hashable](size int) *BloomFilter[T] {
	byteLen := (size + 7) / 8
	return &BloomFilter[T]{
		bitmap: make([]byte, byteLen),
		count:  0,
	}
}

// Add elements to the filter
func (b *BloomFilter[T]) Add(value T) {
	b.mu.Lock()
	defer b.mu.Unlock()

	hash := b.hash(value)

	// Locate the bit to switch on
	byteIdx, bitIdx := hash/8, hash%8

	// left shift the binary equivalent of 1 by bitIdx positions and
	// perfom a bitwise OR with the byte at index [byteIdx]
	b.bitmap[byteIdx] |= 1 << bitIdx

	b.count++
}

// Check if an element is in the filter
func (b *BloomFilter[T]) Check(value T) bool {
	b.mu.RLock()
	defer b.mu.RUnlock()

	hash := b.hash(value)

	targetByte := b.bitmap[hash/8]

	// left shift the binary equivalent of 1 by bitIdx positions and
	// perfom a bitwise AND with the targetByte.
	// If the result is equal 1*2^bitIdx, then the bit is set
	return targetByte&(1<<(hash%8)) == 1<<(hash%8)
}

// Clear the filter and reset the count
func (b *BloomFilter[T]) Reset() {
	b.mu.Lock()
	defer b.mu.Unlock()

	byteLen := len(b.bitmap)
	b.bitmap = make([]byte, byteLen)
	b.count = 0
}

func (b *BloomFilter[T]) hash(value T) uint64 {
	hashFunc := fnv.New64a()

	var data string

	switch reflect.TypeOf(value).Kind() {
	// integers
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		iValue := reflect.ValueOf(value).Int()
		data = strconv.FormatInt(iValue, 10)

	// unsigned integers
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		uValue := reflect.ValueOf(value).Uint()
		data = strconv.FormatUint(uValue, 10)

	// floats
	case reflect.Float32, reflect.Float64:
		floatValue := reflect.ValueOf(value).Float()
		data = strconv.FormatFloat(floatValue, 'f', -1, 64)

	// strings
	case reflect.String:
		data = reflect.ValueOf(value).String()
	}

	hashFunc.Write([]byte(data))

	return hashFunc.Sum64() % uint64(len(b.bitmap))
}
