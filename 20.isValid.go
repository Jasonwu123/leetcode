package leetcode

/*
20. 有效的括号
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true
*/
func isValid(s string) bool {
	if s == "" {
		return true
	}

	pre := make([]byte, len(s)/2)
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '[' || s[i] == '(' || s[i] == '{' {
			pre[count] = s[i]
			count++
			continue
		}
		if s[i]-pre[count-1] > 2 {
			return false
		}
		count--
	}
	return true
}
