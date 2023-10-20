package aho_corasick

import (
	"bufio"
	"encoding/hex"
	"os"
	"strings"
)

type state struct {
	id       int64           //节点id
	value    byte            //值
	parent   *state          //当前节点的父节点
	trans    map[byte]*state //当前节点子节点
	dict     int64           //表示当前节点是模式串的尾部，且模式串的长度
	failLink *state          //失败指针，指向当前节点的最长前缀
	dictLink *state          //该节点失败指针为另一模式串尾部，则指向尾部
	pattern  int64           //记录第几个pattern，只记录尾节点
}

//用于构建Tries
type TrieBuilder struct {
	states      []*state
	root        *state
	numPatterns int64
}

//创建并初始化一新的TrieBuilder
func NewTrieBuilder() *TrieBuilder {
	tb := &TrieBuilder{
		states:      make([]*state, 0),
		root:        nil,
		numPatterns: 0,
	}
	tb.addState(0, nil)
	tb.addState(0, nil)
	tb.root = tb.states[1]
	return tb
}

func (tb *TrieBuilder) addState(value byte, parent *state) *state {
	s := &state{
		id:       int64(len(tb.states)),
		value:    value,
		parent:   parent,
		trans:    make(map[byte]*state),
		dict:     0,
		failLink: nil,
		dictLink: nil,
		pattern:  0,
	}
	tb.states = append(tb.states, s)
	return s
}

//将字节切片加入到Trie中
func (tb *TrieBuilder) AddPattern(pattern []byte) *TrieBuilder {
	s := tb.root
	var t *state
	var ok bool

	for _, c := range pattern {
		if t, ok = s.trans[c]; !ok {
			t = tb.addState(c, s)
			s.trans[c] = t
		}
		s = t
	}

	s.dict = int64(len(pattern))
	s.pattern = tb.numPatterns
	tb.numPatterns++

	return tb
}

//将多个字节切片加入到Trie中
func (tb *TrieBuilder) AddPatterns(patterns [][]byte) *TrieBuilder {
	for _, pattern := range patterns {
		tb.AddPattern(pattern)
	}
	return tb
}

//将字符串加入到Trie中
func (tb *TrieBuilder) AddString(pattern string) *TrieBuilder {
	return tb.AddPattern([]byte(pattern))
}

//将多个字符串加入到Trie中
func (tb *TrieBuilder) AddStrings(patterns []string) *TrieBuilder {
	for _, pattern := range patterns {
		tb.AddString(pattern)
	}
	return tb
}

//从文件加载字节模式,需每行的十六进制形式
func (tb *TrieBuilder) LoadPatterns(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	for s.Scan() {
		str := strings.TrimSpace(s.Text())
		if len(str) != 0 {
			pattern, err := hex.DecodeString(str) //返回由十六进制字符串 s 表示的字节
			if err != nil {
				return err
			}
			tb.AddPattern(pattern)
		}
	}

	return s.Err()
}

//从文件加载字符串模式。 每行需要一个模式。
func (tb *TrieBuilder) LoadStrings(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	for s.Scan() {
		str := strings.TrimSpace(s.Text())
		if len(str) != 0 {
			tb.AddString(str)
		}
	}
	return s.Err()
}

//构造Trie
func (tb *TrieBuilder) Build() *Trie {
	//采用DFS计算前缀树所有节点的失败指针
	tb.computeFailLinks(tb.root)
	//采用DFS计算dictLink
	tb.computeDictLinks(tb.root)

	numStates := len(tb.states)

	/*创建全局变量，通过Trie返回。
	  这样可以方便地通过访问数组来匹配，避免访问树，数组的访问效率相对更高*/
	dict := make([]int64, numStates)
	trans := make([][256]int64, numStates)
	failLink := make([]int64, numStates)
	dictLink := make([]int64, numStates)
	pattern := make([]int64, numStates)

	for i, s := range tb.states {
		dict[i] = s.dict
		pattern[i] = s.pattern
		for c, t := range s.trans {
			trans[i][c] = t.id
		}
		if s.failLink != nil {
			failLink[i] = s.failLink.id
		}
		if s.dictLink != nil {
			dictLink[i] = s.dictLink.id
		}
	}
	return &Trie{dict, trans, failLink, dictLink, pattern}
}

func (tb *TrieBuilder) computeFailLinks(s *state) {
	//节点已有失败指针
	if s.failLink != nil {
		return
	}

	if s == tb.root || s.parent == tb.root {
		//节点为root或节点父节点为root
		s.failLink = tb.root
	} else {
		var ok bool

		for t := s.parent.failLink; t != tb.root; t = t.failLink {
			if t.failLink == nil {
				tb.computeFailLinks(t)
			}

			if s.failLink, ok = t.trans[s.value]; ok {
				break
			}
		}

		if s.failLink == nil {
			if s.failLink, ok = tb.root.trans[s.value]; !ok {
				s.failLink = tb.root
			}
		}
	}

	//节点子节点进行失败指针求值
	for _, t := range s.trans {
		tb.computeFailLinks(t)
	}
}

func (tb *TrieBuilder) computeDictLinks(s *state) {
	if s != tb.root {
		for t := s.failLink; t != tb.root; t = t.failLink {
			if t.dict != 0 {
				s.dictLink = t
				break
			}
		}
	}

	for _, t := range s.trans {
		tb.computeDictLinks(t)
	}
}
