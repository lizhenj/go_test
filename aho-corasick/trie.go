package aho_corasick

const (
	rootState int64 = 1
	nilState  int64 = 0
)

//Trie 表示按照 Aho-Corasick 算法具有额外链接的模式字典树
type Trie struct {
	dict     []int64      //i:序号 v:模式串的长度
	trans    [][256]int64 //i:序号 j:byte v:节点id
	failLink []int64      //i:序号 v:失败指针节点id
	dictLink []int64      //i:序号 v:失败指针为模式串的尾时的节点id
	pattern  []int64      //i:序号 v:第几个模式串 只在尾点记录
}

// Walk 在任何匹配上调用此函数，给出结束位置、匹配字节的长度， 和模式编号。
type WalkFn func(end, n, pattern int64) bool

// Walk 在给定的输出上运行算法，在每个输出上调用提供的回调函数
// 匹配。 如果回调函数返回 false，算法将终止。
func (tr *Trie) Walk(input []byte, fn WalkFn) {
	s := rootState

	for i, c := range input {
		t := tr.trans[s][c]
		// 字符c不是s节点的子节点
		if t == nilState {
			//从失败指针开始继续匹配
			for u := tr.failLink[s]; u != rootState; u = tr.failLink[u] {
				if t = tr.trans[u][c]; t != nilState {
					break
				}
			}

			if t == nilState {
				if t = tr.trans[rootState][c]; t == nilState {
					t = rootState
				}
			}
		}

		s = t
		//s节点是个pattern的尾部
		if tr.dict[s] != 0 {
			//调用回调函数，保存匹配信息
			if !fn(int64(i), tr.dict[s], tr.pattern[s]) {
				return
			}
		}

		//s节点还是另一个pattern的尾部
		if tr.dictLink[s] != nilState {
			for u := tr.dictLink[s]; u != nilState; u = tr.dictLink[u] {
				if !fn(int64(i), tr.dict[u], tr.pattern[u]) {
					return
				}
			}
		}
	}
}

//Match 对字节输入运行 Aho-Corasick 字符串搜索算法。
func (tr *Trie) Match(input []byte) []*Match {
	matches := make([]*Match, 0)
	tr.Walk(input, func(end, n, pattern int64) bool {
		pos := end - n + 1
		matches = append(matches, newMatch(pos, pattern, input[pos:pos+n]))
		return true
	})
	return matches
}

//MatchFirst 与 Match 相同，但在第一次成功匹配后返回。
func (tr *Trie) MatchFirst(input []byte) *Match {
	var match *Match
	tr.Walk(input, func(end, n, pattern int64) bool {
		pos := end - n + 1
		match = &Match{pos: pos, match: input[pos : pos+n]}
		return false
	})
	return match
}

//MatchString 对字符串输入运行 Aho-Corasick 字符串搜索算法。
func (tr *Trie) MatchString(input string) []*Match {
	return tr.Match([]byte(input))
}

//MatchFirstString 与 MatchString 相同，但在第一次成功匹配后返回。
func (tr *Trie) MatchFirstString(input string) *Match {
	return tr.MatchFirst([]byte(input))
}
