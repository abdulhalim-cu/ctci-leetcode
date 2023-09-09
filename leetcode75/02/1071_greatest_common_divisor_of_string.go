package _2

func gcdOfStrings(str1 string, str2 string) string {
	if (str1 + str2) == (str2 + str1) {
		return str1[0:GCD(len(str1), len(str2))]
	}
	return ""
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
