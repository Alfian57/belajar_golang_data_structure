package main

import "fmt"

const ArraySize = 7

// HashTable Structure
type HashTable struct {
	array [ArraySize]*bucket
}

// Bucket Structure
type bucket struct {
	head *bucketNode
}

// BucketNode Structure
// bucketNode is a linked list node that holds the key
type bucketNode struct {
	key  string
	next *bucketNode
}

// Insert (Hash Table)
// Insert will take in a key and add it to the hash table array
func (h *HashTable) Insert(k string) {
	index := hash(k)
	h.array[index].insert(k)
}

// Search (Hash Table)
// Search will take in a key and return true if that key is stored in the hash table
func (h *HashTable) Search(k string) bool {
	index := hash(k)
	return h.array[index].search(k)
}

// Delete (Hash Table)
// Delete will take in a key and delete it from the hash table
func (h *HashTable) Delete(k string) {
	index := hash(k)
	h.array[index].delete(k)
}

// insert (Bucket)
// insert will take in a key, create a node with the key and
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k, next: b.head}
		b.head = newNode
	} else {
		fmt.Println(k, " already exist")
	}
}

// search (Bucket)
// search will take in a key and return true if the bucket has
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}

	return false
}

// delete (Bucket)
func (b *bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}

	toDelete := b.head
	for toDelete.next.key != k {
		if toDelete.next.next == nil {
			return
		}
		toDelete = toDelete.next
	}
	toDelete.next = toDelete.next.next
}

// hash
func hash(k string) int {
	sum := 0

	for _, v := range k {
		sum += int(v)
	}

	return sum % ArraySize
}

// Init
// Init will create a bucket in each slog of the hash table
func Init() *HashTable {
	hashTable := &HashTable{}

	for i := range hashTable.array {
		hashTable.array[i] = &bucket{}
	}

	return hashTable
}

func main() {
	myHashTable := Init()
	list := []string{
		"ALFIAN",
		"GADING",
		"SAPUTRA",
		"CHOIRIL",
		"LINTANG",
		"ALFARIZKY",
		"RASYID",
		"AKMAL",
	}

	for _, v := range list {
		myHashTable.Insert(v)
	}

	fmt.Println("GADING is exist : ", myHashTable.Search("GADING"))
	myHashTable.Delete("ALFIAN")
	fmt.Println("ALFIAN is exist : ", myHashTable.Search("ALFIAN"))
}
