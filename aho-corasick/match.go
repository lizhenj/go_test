package aho_corasick

import (
	"bytes"
	"fmt"
)

type Match struct {
	pos     int64  //结束位置
	pattern int64  //
	match   []byte //
}

func newMatch(pos, pattern int64, match []byte) *Match {
	return &Match{pos, pattern, match}
}

func newMatchString(pos, pattern int64, match string) *Match {
	return &Match{pos: pos, pattern: pattern, match: []byte(match)}
}

func (m *Match) String() string {
	return fmt.Sprintf("{%d %d %q}", m.pos, m.pattern, m.match)
}

func (m *Match) Pos() int64 {
	return m.pos
}

func (m *Match) Pattern() int64 {
	return m.pattern
}

func (m *Match) Match() []byte {
	return m.match
}

func (m *Match) MatchString() string {
	return string(m.match)
}

func MatchEqual(a, b *Match) bool {
	return a.pos == b.pos && a.pattern == b.pattern && bytes.Equal(a.match, b.match)
}
