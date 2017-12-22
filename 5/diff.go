package main

import (
	"math"
	"fmt"
)

func diff(a, b string, unified bool) string {
	var ap, bp string
	if unified {
		ap = "- "
		bp = "+ "
	} else {
		ap = "< "
		bp = "> "
	}

	codedA, codedB, codes := code(a, b)

	dp := DP{
		a: codedA,
		b: codedB,
	}
	diff := dp.distance(len(codedA), len(codedB))

	return prettyText(diff, codes, bp, ap)
}

type Kind int8

const (
	Insert Kind = iota + 1
	Delete
	Equal
)

type Op struct {
	Kind Kind
	Text rune
}

type DP struct {
	a, b []rune
}

func (dp DP) distance(i, j int) []Op {
	if i == 0 && j == 0 {
		return []Op{}
	}
	if i == 0 {
		return append(dp.distance(0, j-1), Op{Kind: Insert, Text: dp.b[j-1]})
	}
	if j == 0 {
		return append(dp.distance(i-1, 0), Op{Kind: Delete, Text: dp.a[i-1]})
	}

	ac, bc := dp.a[i-1], dp.b[j-1]

	if ac == bc {
		return minD(
			append(dp.distance(i-1, j), Op{Kind: Delete, Text: ac}),
			append(dp.distance(i, j-1), Op{Kind: Insert, Text: bc}),
			append(dp.distance(i-1, j-1), Op{Kind: Equal, Text: ac}),
		)
	} else {
		return minD(
			append(dp.distance(i-1, j), Op{Kind: Delete, Text: ac}),
			append(dp.distance(i, j-1), Op{Kind: Insert, Text: bc}),
		)
	}
}

// minD は引数の中から最も小さい編集距離を持つ []Op を返します。
func minD(opLists ...[]Op) []Op {
	md := math.MaxInt32
	var ans []Op

	for _, ops := range opLists {
		d := d(ops)
		if md > d {
			md = d
			ans = ops
		}
	}

	return ans
}

// d は []Op の編集距離を返します。
func d(ops []Op) int {
	res := 0
	for _, op := range ops {
		if op.Kind == Delete || op.Kind == Insert {
			res++
		}
	}

	return res
}

func code(a, b string) ([]rune, []rune, map[rune]string) {
	var i, j []rune

	m := make(map[string]rune)

	aLines := splitToLines(a)
	bLines := splitToLines(b)

	count := 0
	for _, line := range aLines {
		_, ok := m[line]
		if !ok {
			m[line] = rune(count)
			count++
		}
		i = append(i, m[line])
	}

	for _, line := range bLines {
		_, ok := m[line]
		if !ok {
			m[line] = rune(count)
			count++
		}
		j = append(j, m[line])
	}

	ans := make(map[rune]string)
	for k, v := range m {
		ans[v] = k
	}

	return i, j, ans
}

func splitToLines(str string) []string {
	rs := []rune(str)
	ans := []string{}
	buf := ""
	for _, r := range rs {
		buf += string(r)
		if r == '\n' {
			ans = append(ans, buf)
			buf = ""
		}
	}
	return ans
}

func prettyText(ops []Op, codes map[rune]string, insertPrefix, deletePrefix string) string {
	res := ""

	for _, op := range ops {
		var line string

		switch op.Kind {
		case Insert:
			line = fmt.Sprintf("%s%s", insertPrefix, codes[op.Text])
		case Delete:
			line = fmt.Sprintf("%s%s", deletePrefix, codes[op.Text])
		case Equal:
			line = fmt.Sprintf("  %s", codes[op.Text])
		}
		res = fmt.Sprintf("%s%s", res, line)
	}

	return res
}
