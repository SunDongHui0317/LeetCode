package ismatch

import (
	"fmt"
	"log"
	"strings"
)

func debug(v ...interface{}) {
	log.Println(v...)
}

/* 有限状态机（正则表达式的实现） */
func isMatch(s string, p string) bool {
	begin := new(Node)
	begin.C = '>'
	begin.Size = generatePattern(begin, p, 0)
	debug("begin:", begin.String())
	return check(begin, s, 0)
}

type Node struct {
	C        byte
	Parent   *Node
	Children map[byte][]*Node
	End      bool
	Size     int
}

func (n *Node) String() string {
	return n.StringLevel(0, make(map[*Node]bool))
}

func (n *Node) StringLevel(level int, finishNodes map[*Node]bool) string {
	r := make([]string, 0)
	if n.End {
		r = append(r, fmt.Sprintf("  id%s{%v};", toString(n), string(n.C)))
	} else {
		r = append(r, fmt.Sprintf("  id%s(%v);", toString(n), string(n.C)))
	}
	finishNodes[n] = true
	for k, v := range n.Children {
		for _, c := range v {
			if _, ok := finishNodes[c]; !ok {
				r = append(r, c.StringLevel(level+1, finishNodes))
			}
			r = append(r, fmt.Sprintf("  id%s -- %s --> id%s;", toString(n), string(k), toString(c)))
		}
	}
	return strings.Join(r, "\n")
}

func (n *Node) Append(c byte, child *Node) {
	m := n.Children
	if m == nil {
		m = make(map[byte][]*Node)
		n.Children = m
	}
	list := m[c]
	if list == nil {
		list = make([]*Node, 0)
	}
	for _, v := range list {
		if v == child {
			m[c] = list
			return
		}
	}
	list = append(list, child)
	m[c] = list
}

func generatePattern(now *Node, str string, idx int) int {
	if len(str) <= idx {
		now.End = true
		return now.Size
	}
	vnow := now
	switch str[idx] {
	case '*':
		now.Size = 0
		now.Append(now.C, now)
	default:
		node := new(Node)
		node.C = str[idx]
		now.Append(str[idx], node)
		node.Parent = now
		node.Size = 1
		vnow = node
	}
	ret := generatePattern(vnow, str, idx+1)
	if ret == 0 {
		now.End = true
	}
	addParent := now
	for addParent.Parent != nil {
		if addParent.Size == 0 {
			debug(toString(vnow), " -> ", toString(addParent.Parent))
			addParent.Parent.Append(vnow.C, vnow)
			addParent = addParent.Parent
		} else {
			break
		}
	}
	return now.Size + ret
}

func check(now *Node, str string, idx int) bool {
	if len(str) <= idx {
		return now.End
	}
	list := now.Children['.']
	for _, v := range now.Children[str[idx]] {
		list = append(list, v)
	}
	for _, v := range list {
		r := check(v, str, idx+1)
		if r {
			return true
		}
	}
	return false
}

func toString(i interface{}) string {
	switch i.(type) {
	case int:
		return fmt.Sprintf("%v", i)
	case string:
		return fmt.Sprintf("%v", i)
	case bool:
		return fmt.Sprintf("%v", i)
	default:
		return fmt.Sprintf("%p", i)
	}
}

/*Runtime: 0 ms, faster than 100.00% of Go online submissions for Regular Expression Matching.
Memory Usage: 2.2 MB, less than 56.82% of Go online submissions for Regular Expression Matching.
*/
func ismatch(s string, p string) bool {
	n := len(s)
	m := len(p)
	dp := make([][]bool, n + 1)
	for i := range(dp) {
		dp[i] = make([]bool, m + 1)
	}
	dp[0][0] = true
	if m > 0 && n > 0 {
		dp[1][1] = s[0] == p[0] || p[0] == '.'
	}
	for j:=2 ;j<=m; j++ {
		dp[0][j] = dp[0][j-2] && p[j-1] == '*'
	}
	for i:=1; i<=n; i+=1 {
		for j:=2; j<=m; j+=1 {
			if p[j-1] != '*' {
				dp[i][j] = dp[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1] == '.')
			} else {
				dp[i][j] = dp[i][j-2] || (dp[i-1][j] && (s[i-1]==p[j-2] || p[j-2] == '.'))
			}
		}
	}
	return dp[n][m]
}