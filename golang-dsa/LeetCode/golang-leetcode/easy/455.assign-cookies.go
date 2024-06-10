package easy

import (
	"sort"
)

/*
 * @lc app=leetcode id=455 lang=golang
 *
 * [455] Assign Cookies
 */

// @lc code=start
func findContentChildren(g []int, s []int) int {

	sort.Ints(g)
	sort.Ints(s)

	//	g = [1, 3, 3, 4]
	//	s = [1, 2, 3]
	gi, gLen, si, sLen := 0, len(g), 0, len(s)

	for gi < gLen && si < sLen {
		gVal := g[gi]
		sVal := s[si]

		if gVal <= sVal {
			gi++
			si++
		} else {
			si++
		}
	}

	return gi
}

// @lc code=end
