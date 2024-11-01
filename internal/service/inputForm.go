package service

import "regexp"

const EmailPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
const PasswordPattern = `^[a-zA-Z0-9!@#$%^&*()_+={}|[\]\\:";'<>?,./]{6,}$` //至少6位数字或大小写字母

// 判断格式是否合法
func IsValidForm(pattern string, s string) (bool, error) {

	matched, err := regexp.MatchString(pattern, s)
	if err != nil {
		return false, err
	}
	return matched, nil
}
