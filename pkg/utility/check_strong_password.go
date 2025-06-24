package utility

import (
	"errors"
	"unicode"
)

func CheckStrongPassword(password string) error {
	if len(password) < 8 {
		return errors.New("รหัสผ่านต้องมีความยาวอย่างน้อย 8 ตัวอักษร")
	}

	var hasUpper, hasLower, hasNumber, hasSpecial bool

	for _, ch := range password {
		switch {
		case unicode.IsUpper(ch):
			hasUpper = true
		case unicode.IsLower(ch):
			hasLower = true
		case unicode.IsDigit(ch):
			hasNumber = true
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			hasSpecial = true
		}
	}

	if !hasUpper {
		return errors.New("รหัสผ่านต้องมีตัวอักษรภาษาอังกฤษตัวใหญ่ (A-Z) อย่างน้อย 1 ตัว")
	}
	if !hasLower {
		return errors.New("รหัสผ่านต้องมีตัวอักษรภาษาอังกฤษตัวเล็ก (a-z) อย่างน้อย 1 ตัว")
	}
	if !hasNumber {
		return errors.New("รหัสผ่านต้องมีตัวเลข (0-9) อย่างน้อย 1 ตัว")
	}
	if !hasSpecial {
		return errors.New("รหัสผ่านต้องมีอักขระพิเศษ (เช่น !@#$%^&*) อย่างน้อย 1 ตัว")
	}

	return nil
}
