package trie

import (
	"testing"
)

func TestTrie(t *testing.T) {
	trie := New()
	trie.Insert("apple")
	if !trie.Search("apple") {
		t.Errorf("search apple failed")
	}
	if trie.Search("app") {
		t.Errorf("search app failed")
	}
	if !trie.StartsWith("app") {
		t.Errorf("startsWith app failed")
	}
	trie.Insert("app")
	if !trie.Search("app") {
		t.Errorf("search app failed")
	}
}
