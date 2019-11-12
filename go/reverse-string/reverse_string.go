package reverse

func Reverse(s string) string {
	str := []rune(s)
	rev := make([]rune, len(str))

	for i := 0; i < len(str); i++ {
		rev[len(str)-1-i] = str[i]
	}
	return string(rev)
}
