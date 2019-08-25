//验证给定的字符串是否可以解释为十进制数字。
//
//例如:
//
//"0" => true
//" 0.1 " => true
//"abc" => false
//"1 a" => false
//"2e10" => true
//" -90e3   " => true
//" 1e" => false
//"e3" => false
//" 6e-1" => true
//" 99e2.5 " => false
//"53.5e93" => true
//" --6 " => false
//"-+3" => false
//"95a54e53" => false
//
//说明: 我们有意将问题陈述地比较模糊。在实现代码之前，你应当事先思考所有可能的情况。这里给出一份可能存在于有效十进制数字中的字符列表：
//
//数字 0-9
//指数 - "e"
//正/负号 - "+"/"-"
//小数点 - "."
//当然，在输入中，这些字符的上下文也很重要。

package number

import (
	"strings"
)

func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	if len(s) < 1 {
		return false
	}
	if len(s) == 1 {
		return s[0] >= '0' && s[0] <= '9'
	}
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
	}
	idxE := strings.Index(s, "e")
	if idxE != -1 {
		if idxE == len(s) - 1 || idxE == 0 {
			return false
		}
		return isFloat(s[:idxE]) && isSignedInt(s[idxE+1:])
	}
	return isFloat(s)
}

func isFloat(s string) bool {
	idxDot := strings.Index(s, ".")
	if idxDot == -1 {
		return isInt(s)
	} else {
		if idxDot == 0 {
			return isInt(s[1:])
		} else if idxDot == len(s) - 1 {
			return isInt(s[:idxDot])
		} else {
			return isInt(s[:idxDot]) && isInt(s[idxDot+1:])
		}
	}
}

func isSignedInt(s string) bool {
	if len(s) == 0 {
		return false
	}
	if s[0] == '-' || s[0] == '+' {
		return len(s) > 1 && isInt(s[1:])
	}
	return isInt(s)
}

func isInt(s string) bool {
	if len(s) == 0 {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}