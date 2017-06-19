package main

import (
	"fmt"
)

func reverseRunes(src []rune) []rune {
	ret := make([]rune, len(src))
	j := 0
	for i := len(src) - 1; i >= 0; i-- {
		ret[j] = src[i]
		j++
	}
	return ret
}

type SplitBuffer struct {
	pre  []rune
	post []rune
}

func (r *SplitBuffer) SetText(s string) {
	r.pre = []rune{}
	r.post = reverseRunes([]rune(s))
}

func (r *SplitBuffer) GetText() string {
	ret := make([]rune, len(r.pre) + len(r.post))

	copy(ret, r.pre)
	copy(ret[len(r.pre):], reverseRunes(r.post))

	return string(ret)
}

func (r *SplitBuffer) Insert(s string) {
	r.pre = append(r.pre, []rune(s)...)
}

func (r *SplitBuffer) Delete() {
	if len(r.post) == 0 {
		return
	}

	r.post = r.post[:len(r.post) - 1]
}

func (r *SplitBuffer) Backspace() {
	if len(r.pre) == 0 {
		return
	}

	r.pre = r.pre[:len(r.pre) - 1]
}

func (r *SplitBuffer) CursorNext() {
	if len(r.post) == 0 {
		return
	}

	r.pre = append(r.pre, r.post[len(r.post) - 1])
	r.post = r.post[:len(r.post) - 1]
}

func (r *SplitBuffer) CursorPrevious() {
	if len(r.pre) == 0 {
		return
	}

	r.post = append(r.post, r.pre[len(r.pre) - 1])
	r.pre = r.pre[:len(r.pre) - 1]
}

func (r *SplitBuffer) debugPrint() {
	for _, c := range r.pre {
		preColor.Printf("%c", c)
	}

	for i := len(r.post) - 1; i >= 0; i-- {
		postColor.Printf("%c", r.post[i])
	}

	fmt.Printf("\n")
}
