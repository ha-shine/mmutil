package mmutil

import (
	"regexp"
)

// Regular expression strnig to detect zawgyi
const regString = `[ဳဴၚၠ-႟]|ေ[ႏ႐]|ေ[ျၾ-ႄ]|[^က-အဩျ်ြွ]ေ|[^က-အဩေ]ျ|\s[ေျၾ-ႄ]|^[ေျၾ-ႄ]|္[^က-ဪ]|င္|ြ[ျၾ-ႄ]|ြ[်႐]|ွြ|ုု`
var reg = regexp.MustCompile(regString)

// IsZawgyi detects whether the given byte array is written in Zawgyi.
// This is a low level implementation and you should be using StringIsZawgyi
func IsZawgyi(b []byte) bool {
	return reg.Find(b) != nil
}

// StringIsZawgyi detects whether the given string is written in Zawgyi
func StringIsZawgyi(s string) bool {
	return IsZawgyi([]byte(s))
}