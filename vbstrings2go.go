package vbstrings2go

//InStr search inside a string
func InStr(str1 string, str2 string) int {
	r1 := []rune(str1)
	r2 := []rune(str2)
	l1 := len(r1)
	l2 := len(r2)
	switch {
	case l1 == 0:
		return -1
	case l2 == 0:
		return -1
	case l1 == l2:
		if str1 == str2 {
			return 1
		}
		return -1
	case l2 > l1:
		return -1
	default:
		var found bool
		for i := 0; i <= l1-l2; i++ {
			if r1[i] == r2[0] {
				found = true
				for j := 0; j < l2; j++ {
					if r1[i+j] != r2[j] {
						found = false
						break
					}
				}
				if found {
					return i + 1
				}
			}
		}
		return -1
	}
}

//Len length of the string
func Len(str string) int {
	return len([]rune(str))
}

//Left selects characters from a string from the beginning
func Left(str string, length int) string {
	if length == 0 {
		return ""
	}
	r := []rune(str)
	if len(r) <= length {
		return str
	}
	var strReturn string = ""
	for i := 0; i < length; i++ {
		strReturn += string(r[i])
	}
	return strReturn
}

//Right selects characters from a string from the end
func Right(str string, length int) string {
	if length == 0 {
		return ""
	}
	r := []rune(str)
	l := len(r)
	if l <= length {
		return str
	}
	var strReturn string = ""
	for i := l - length; i <= l-1; i++ {
		strReturn += string(r[i])
	}
	return strReturn
}

//Mid selects characters from the middle starting from position to the given length
func Mid(str string, pos int, length int) string {
	if length == 0 {
		return ""
	}
	r := []rune(str)
	l := len(r)
	if pos > l {
		return ""
	}
	var strReturn string = ""
	if pos+length-2 > l {
		for i := pos - 1; i <= l-1; i++ {
			strReturn += string(r[i])
		}
	} else {
		for i := pos - 1; i <= pos+length-2; i++ {
			strReturn += string(r[i])
		}
	}
	return strReturn
}

//MidSimple selects characters from the middle from position to end of string
func MidSimple(str string, pos int) string {
	r := []rune(str)
	l := len(r)
	if pos > l {
		return ""
	}
	var strReturn string = ""
	for i := pos - 1; i <= l-1; i++ {
		strReturn += string(r[i])
	}
	return strReturn
}
