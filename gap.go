package main

import (
	"fmt"
)

const minimumGapIncrease = 16

type GapBuffer struct {
	data    []rune
	preLen  int
	postLen int
}

func (r *GapBuffer) SetText(s string) {
	r.data = []rune(s)
	r.preLen = 0
	r.postLen = len(r.data)
}

func (r *GapBuffer) GetText() string {
	ret := make([]rune, r.preLen + r.postLen)

	copy(ret, r.data)
	copy(ret[r.preLen:], r.data[r.postStart():])

	return string(ret)
}

func (r *GapBuffer) gapStart() int {
	return r.preLen
}

func (r *GapBuffer) gapLen() int {
	return r.postStart() - r.preLen
}

func (r *GapBuffer) postStart() int {
	return len(r.data) - r.postLen
}

func (r *GapBuffer) Insert(s string) {
	if r.gapLen() < len(s) {
		increaseBy := len(s) - r.gapLen()

		if increaseBy < minimumGapIncrease {
			increaseBy = minimumGapIncrease
		}

		newData := make([]rune, len(r.data) + increaseBy)
		copy(newData, r.data[:r.preLen])
		copy(newData[r.postStart() + increaseBy:],
			r.data[r.postStart():])

		r.data = newData
	}

	copy(r.data[r.gapStart():], []rune(s))
	r.preLen += len(s)
}

func (r *GapBuffer) Delete() {
	if r.postLen == 0 {
		return
	}

	r.postLen--
}

func (r *GapBuffer) Backspace() {
	if r.preLen == 0 {
		return
	}

	r.preLen--
}

func (r *GapBuffer) CursorNext() {
	if r.postLen == 0 {
		return
	}

	r.data[r.preLen] = r.data[r.postStart()]
	r.preLen++
	r.postLen--
}

func (r *GapBuffer) CursorPrevious() {
	if r.preLen == 0 {
		return
	}

	r.data[r.postStart() - 1] = r.data[r.preLen - 1]
	r.preLen--
	r.postLen++
}

func (r *GapBuffer) debugPrint() {
	for i := 0; i < len(r.data); i++ {
		if i >= r.gapStart() && i < r.gapStart() + r.gapLen() {
			fmt.Printf(" ")
		} else if i < r.preLen {
			preColor.Printf("%c", r.data[i])
		} else {
			postColor.Printf("%c", r.data[i])
		}
	}

	fmt.Printf("\n")
}
