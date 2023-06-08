package trie

import "fmt"

type Node struct {
	end      bool // Indicates that a word does end here
	children [26]*Node
}

func NewNode() *Node {
	return &Node{
		children: [26]*Node{},
	}
}

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{
		root: NewNode(),
	}
}

func (t *Trie) Insert(word string) {

	node := t.root
	for _, r := range word {
		i := int(r - 'a')
		fmt.Println(r, " index ", i)
		if node.children[i] == nil {
			node.children[i] = NewNode()
		}

		node = node.children[i]
	}

	node.end = true
}

// Search returns true if the word has been inserted in the trie
func (t *Trie) Search(str string) bool {
	node := t.root

	for _, r := range str {
		i := int(r - 'a')
		if node.children[i] == nil {
			return false
		}
		node = node.children[i]
	}
	return node.end
}

// StartsWith returns true if there is a word in the trie that starts with the given string
func (t *Trie) StartsWith(str string) bool {
	node := t.root

	for _, r := range str {
		i := int(r - 'a')
		if node.children[i] == nil {
			return false
		}
		node = node.children[i]
	}
	return true

}
