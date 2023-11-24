package main

import "fmt"

const AplhabetSize = 26

// node
// node represents each node in the trie
type node struct {
	children [AplhabetSize]*node
	isEnd    bool
}

// Trie
// Trie represent a trie and has a pointer to the root node
type Trie struct {
	root *node
}

// Insert
// Insert will take in a word and add it to the trie
func (t *Trie) Insert(w string) {
	wordLength := len(w)
	currentNode := t.root

	for i := 0; i < wordLength; i++ {
		wordIndex := w[i] - 'a'

		if currentNode.children[wordIndex] == nil {
			currentNode.children[wordIndex] = &node{}
		}
		currentNode = currentNode.children[wordIndex]
	}
	currentNode.isEnd = true
}

// Search
// Search will take in a word and return true if that word included in the trie
func (t Trie) Search(w string) bool {
	wordLength := len(w)
	currentIndex := t.root

	for i := 0; i < wordLength; i++ {
		wordIndex := w[i] - 'a'

		if currentIndex.children[wordIndex] == nil {
			return false
		}

		currentIndex = currentIndex.children[wordIndex]
	}

	return currentIndex.isEnd
}

// Init
// Init will create a new Trie
func Init() *Trie {
	return &Trie{
		root: &node{},
	}
}

func main() {
	myTrie := Init()

	myTrie.Insert("alfian")
	myTrie.Insert("gading")

	fmt.Println("alfian is exist ", myTrie.Search("alfian"))
	fmt.Println("gading is exist ", myTrie.Search("gading"))
	fmt.Println("saputra is exist ", myTrie.Search("saputra"))
}
