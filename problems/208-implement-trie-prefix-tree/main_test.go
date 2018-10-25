package main

import "testing"

func Test1(t *testing.T) {
	trie := Constructor()
	actual := true
	actual = trie.Search("app") // returns false
	if actual {
		t.Errorf("expected: %v, actual: %v\n", false, actual)
	}
	actual = trie.StartsWith("app") // returns false
	if actual {
		t.Errorf("expected: %v, actual: %v\n", true, actual)
	}

	trie.Insert("apple")
	actual = trie.Search("apple") // returns true
	if !actual {
		t.Errorf("expected: %v, actual: %v\n", true, actual)
	}
	actual = trie.Search("app") // returns false
	if actual {
		t.Errorf("expected: %v, actual: %v\n", false, actual)
	}
	actual = trie.StartsWith("app") // returns true
	if !actual {
		t.Errorf("expected: %v, actual: %v\n", true, actual)
	}
	trie.Insert("app")
	actual = trie.Search("app") // returns true
	if !actual {
		t.Errorf("expected: %v, actual: %v\n", true, actual)
	}
}
