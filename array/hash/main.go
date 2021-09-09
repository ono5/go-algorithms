package main

type HashKeyVal struct {
	key   string
	value int
}

func NewHashKeyVal() *HashKeyVal {
	return &HashKeyVal{}
}

func (h *HashKeyVal) push(key string, value int) {
	h.key = key
	h.value = value
}

type HashTable struct {
	data map[int]*HashKeyVal
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		data: make(map[int]*HashKeyVal, size),
	}
}

func (h *HashTable) _hash(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash = (hash + int([]rune(string(key[i]))[0]))
	}

	return hash
}

func (h *HashTable) set(key string, value int) map[int]*HashKeyVal {
	address := h._hash(key)
	_, ok := h.data[address]
	if !ok {
		h.data[address] = NewHashKeyVal()
	}
	h.data[address].push(key, value)
	return h.data
}

func (h *HashTable) get(key string) {
	address := h._hash(key)
	currentBucket := h.data[address]
	if currentBucket != nil {

	}
}

func main() {
	myHashTable := NewHashTable(50)
	myHashTable.set("grapes", 1000)
	myHashTable.set("apples", 54)
	myHashTable.get("grapes")
}
