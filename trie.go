package datt

func NewTrie() *Trie {
	return &Trie{
		root: &trieNode{
			hash: make(map[rune]*trieNode),
		},
	}
}

type trieNode struct {
	hash map[rune]*trieNode
}

func (tn *trieNode) Append(r rune) *trieNode {
	nextNode := &trieNode{
		hash: make(map[rune]*trieNode),
	}

	tn.hash[r] = nextNode

	return nextNode
}

func (tn *trieNode) AppendEnd() {
	tn.hash['*'] = nil
}

func (tn *trieNode) Contains(r rune) *trieNode {
	node, ok := tn.hash[r]
	if ok {
		return node
	} else {
		return nil
	}
}

type Trie struct {
	root *trieNode
}

func (t *Trie) Insert(word string) {
	currentNode := t.root

	for _, r := range word {
		nextNode := currentNode.Contains(r)

		if nextNode != nil {
			currentNode = nextNode
		} else {
			next := currentNode.Append(r)
			currentNode = next
		}
	}

	currentNode.AppendEnd()
}

func (t *Trie) AutoComplete(prefix string) []string {
	currentNode := t.search(prefix)
	if currentNode == nil {
		return nil
	}
	var words allWords
	t.collectAllWords(currentNode, nil, &words)

	return words.words
}

func (t *Trie) GetAllWords() []string {
	var words allWords
	t.collectAllWords(t.root, nil, &words)

	return words.words
}

func (t *Trie) Contains(word string) bool {
	return t.search(word) != nil
}

type allWords struct {
	words []string
}

func (a *allWords) add(t string) {
	a.words = append(a.words, t)
}

func (t *Trie) collectAllWords(node *trieNode, word []rune, words *allWords) {
	for r, nextNode := range node.hash {
		if r == '*' {
			wordString := string(word)
			words.add(wordString)
		} else {
			wordCopy := make([]rune, len(word)+1)
			copy(wordCopy, word)
			wordCopy = append(wordCopy, r)

			t.collectAllWords(nextNode, wordCopy, words)
		}
	}
}

func (t *Trie) search(word string) *trieNode {
	currentNode := t.root

	for _, r := range word {
		nextNode := currentNode.Contains(r)

		if nextNode == nil {
			return nil
		} else {
			currentNode = nextNode
		}
	}

	return currentNode
}
