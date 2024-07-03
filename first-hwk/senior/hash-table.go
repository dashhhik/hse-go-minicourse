package main

type Entry struct {
	Key   string
	Value int
}

type HashTable struct {
	Table [][]Entry
	Size  int
	Cap   int
}

func Hash(key string, cap int) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash = (31*hash + int(key[i])) % cap
	}
	return hash
}

func Init(cap int) *HashTable {
	return &HashTable{
		Table: make([][]Entry, cap),
		Size:  0,
		Cap:   cap,
	}
}

func (h *HashTable) Insert(key string, value int) {
	index := Hash(key, h.Cap)
	for i := range h.Table[index] {
		if h.Table[index][i].Key == key {
			h.Table[index][i].Value = value
			return
		}
	}
	h.Table[index] = append(h.Table[index], Entry{Key: key, Value: value})
	h.Size++
}

func (h *HashTable) Search(key string) (int, bool) {
	index := Hash(key, h.Cap)
	for _, entry := range h.Table[index] {
		if entry.Key == key {
			return entry.Value, true
		}
	}
	return 0, false
}

func (h *HashTable) Delete(key string) bool {
	index := Hash(key, h.Cap)
	for i, entry := range h.Table[index] {
		if entry.Key == key {
			h.Table[index] = append(h.Table[index][:i], h.Table[index][i+1:]...)
			h.Size--
			return true
		}
	}
	return false
}
