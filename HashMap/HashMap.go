package HashMap

import (
	"errors"
	"hash/fnv"
	"reflect"
	"strconv"
	"sync"
)

type Hashable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

type HashMap[K Hashable, V any] struct {
	mu            sync.Mutex
	buckets       []*bucket[K, V]
	size          int
	initSize      int
	itemsCount    int
	maxLoadFactor float64
}

// initialize new HashMap
func NewHashMap[K Hashable, V any](size int, maxLoadFactor float64) *HashMap[K, V] {
	hashmap := new(HashMap[K, V])
	hashmap.size = size
	hashmap.initSize = size
	hashmap.maxLoadFactor = maxLoadFactor
	hashmap.buckets = make([]*bucket[K, V], size)
	for i := 0; i < size; i++ {
		hashmap.buckets[i] = NewBucket[K, V]()
	}
	return hashmap
}

// Set - Add items to hashmap or update items in hashmap
func (h *HashMap[K, V]) Set(key K, value V) error {
	hash := h.hash(key)
	bucket := h.buckets[hash]

	// check if key already exists. If it does, update the value
	if _, err := bucket.Update(key, value); err == nil {
		return nil
	}

	// check if load factor is greater than max load factor
	if h.LoadFactor() >= h.maxLoadFactor {
		err := h.Resize(getPrime(h.size * 2))

		if err != nil {
			return err
		}
	}

	bucket.Add(key, value)
	h.itemsCount++

	return nil
}

// Get items from hashmap
func (h *HashMap[K, V]) Get(key K) (*node[K, V], error) {
	hash := h.hash(key)
	bucket := h.buckets[hash]

	found, err := bucket.Get(key)

	return found, err
}

// Delete - Remove items from hashmap
func (h *HashMap[K, V]) Delete(key K) (*node[K, V], error) {
	hash := h.hash(key)
	bucket := h.buckets[hash]

	node, err := bucket.Remove(key)

	if err != nil {
		return node, err
	}

	h.itemsCount--

	return node, err
}

// Reset - Clear hashmap
func (h *HashMap[K, V]) Reset() {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.itemsCount = 0
	h.size = h.initSize
	h.buckets = h.buckets[:h.size]

	for i := 0; i < h.size; i++ {
		h.buckets[i].Clear()
	}
}

// Resize hashmap
func (h *HashMap[K, V]) Resize(size int) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	bucketsToCopy := h.size
	h.size = size
	buckets := make([]*bucket[K, V], size)

	for i := 0; i < size; i++ {
		buckets[i] = NewBucket[K, V]()
	}

	bucketsCopied := copy(buckets, h.buckets)

	if bucketsCopied != bucketsToCopy {
		return errors.New("failed to resize hashmap")
	}

	h.buckets = buckets
	return nil
}

// Hash function.
func (h *HashMap[H, T]) hash(key H) uint {
	hashFunc := fnv.New64a()

	var strValue string

	switch reflect.TypeOf(key).Kind() {
	// integers
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		iValue := reflect.ValueOf(key).Int()
		strValue = strconv.FormatInt(iValue, 10)

	// unsigned integers
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		uValue := reflect.ValueOf(key).Uint()
		strValue = strconv.FormatUint(uValue, 10)

	// floats
	case reflect.Float32, reflect.Float64:
		floatValue := reflect.ValueOf(key).Float()
		strValue = strconv.FormatFloat(floatValue, 'f', -1, 64)

	// strings
	case reflect.String:
		strValue = reflect.ValueOf(key).String()
	}

	hashFunc.Write([]byte(strValue))

	return uint(hashFunc.Sum64()) % uint(h.size)
}

// Get size of hashmap
func (h *HashMap[K, V]) Size() int {
	return h.size
}

// Get count of items in hashmap
func (h *HashMap[K, V]) Count() int {
	return h.itemsCount
}

// Get load factor of hashmap
func (h *HashMap[K, V]) LoadFactor() float64 {
	return float64(h.itemsCount) / float64(h.size)
}

// Set load factor of hashmap (default is 0.75)
func (h *HashMap[K, V]) SetLoadFactor(loadFactor float64) {
	h.maxLoadFactor = loadFactor
}

func getPrime(min int) int {
	current := min
	for {
		if isPrime(current) {
			return current
		}
		current++
	}
}

func isPrime(n int) bool {
	if n <= 3 {
		return n >= 2
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i < n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}
