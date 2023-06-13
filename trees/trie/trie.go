package trie

type Node struct {
	children [26]*Node // 字符集只包含小写字母，所以使用长度为 26 的数组来存储子节点
	isEnd    bool      // 标记是否是某个单词的结尾
}

type Trie struct {
	root *Node
}

func New() *Trie {
	return &Trie{
		root: &Node{},
	}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, c := range word {
		idx := c - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &Node{}
		}
		node = node.children[idx]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.root
	for _, c := range word {
		idx := c - 'a'
		if node.children[idx] == nil {
			return false
		}
		node = node.children[idx]
	}
	return node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, c := range prefix {
		idx := c - 'a'
		if node.children[idx] == nil {
			return false
		}
		node = node.children[idx]
	}
	return true
}
