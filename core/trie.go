package core

const ALPHABET_SIZE int = 26

type TrieNode struct {
	children    [ALPHABET_SIZE]*TrieNode
	isEndOfWord bool
}

func GetNode() *TrieNode {
	node := new(TrieNode)
	node.isEndOfWord = false

	for i := 0; i < ALPHABET_SIZE; i++ {
		node.children[i] = nil
	}
	return node
}

func Insert(node *TrieNode, key string) {
	for i := 0; i < len(key); i++ {
		index := key[i] - 'a'
		if node.children[index] == nil {
			node.children[index] = GetNode()
		}
		node = node.children[index]
	}
	// mark last node as leaf
	node.isEndOfWord = true
}

// Return true if key exists in the Trie, false otherwise
func Search(node *TrieNode, key string) bool {
	for i := 0; i < len(key); i++ {
		index := key[i] - 'a'
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}
	return (node != nil && node.isEndOfWord)
}
