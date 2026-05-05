package tree

const (
	alphabetSize = 26
)

type (
	trieNode struct {
		children  [alphabetSize]*trieNode
		isWordEnd bool
	}
	trie struct {
		root *trieNode
	}
)

// Autocompletion solves the problem in O(n) time and O(n) space.
func (t *trie) Autocompletion(word string) []string {
	panic("not implemented")

}

func newTrie(dict []string) *trie {
	panic("not implemented")

}

func (t *trie) insert(word string) {
	panic("not implemented")

}

func (t *trieNode) readWords() []string {
	panic("not implemented")

}
