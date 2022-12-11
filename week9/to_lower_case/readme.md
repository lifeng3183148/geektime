### 解题思路

一重循环判断字符的ASCII码在'A'到'Z'之间的字符通过（+'a'-'A'）操作变成小写

由于golang的[]byte无法直接修改其值，所以需要申请一个新的[]byte

### 代码实现

```go
func toLowerCase(s string) string {
	var newS = make([]byte, 0, len(s))
	for i := range s {
		if s[i] >= 'A' && s[i] <= 'Z' {
			newS = append(newS, s[i]+'a'-'A')
		} else {
			newS = append(newS, s[i])
		}
	}

	return string(newS)
}
```

### 时间空间复杂度

时间复杂度为O(n)

空间复杂度为O(n)