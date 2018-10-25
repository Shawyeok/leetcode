package main

func main() {
}

type Trie struct {
	nodes  map[byte]*Trie
	isLeaf bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	letter := word[0]
	if this.nodes == nil {
		this.nodes = make(map[byte]*Trie)
	}
	if this.nodes[letter] == nil {
		this.nodes[letter] = &Trie{}
	}
	if len(word) > 1 {
		this.nodes[letter].Insert(word[1:])
	} else {
		this.nodes[letter].isLeaf = true
	}
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	letter := word[0]
	trie := this.nodes[letter]
	if trie == nil {
		return false
	}
	if len(word) == 1 {
		return trie.isLeaf
	}
	return trie.Search(word[1:])
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	letter := prefix[0]
	trie := this.nodes[letter]
	if trie == nil {
		return false
	}
	if len(prefix) == 1 {
		return true
	}
	return trie.StartsWith(prefix[1:])
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
