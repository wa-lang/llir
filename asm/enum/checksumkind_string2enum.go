// Code generated by "string2enum -linecomment -type ChecksumKind ../../ir/enum"; DO NOT EDIT.

package enum

import (
	"fmt"

	"github.com/llir/llvm/ir/enum"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the string2enum command to generate them again.
	var x [1]struct{}
	_ = x[enum.ChecksumKindMD5-1]
	_ = x[enum.ChecksumKindSHA1-2]
}

const _ChecksumKind_name = "CSK_MD5CSK_SHA1"

var _ChecksumKind_index = [...]uint8{0, 7, 15}

// ChecksumKindFromString returns the ChecksumKind enum corresponding to s.
func ChecksumKindFromString(s string) enum.ChecksumKind {
	if len(s) == 0 {
		return 0
	}
	for i := range _ChecksumKind_index[:len(_ChecksumKind_index)-1] {
		if s == _ChecksumKind_name[_ChecksumKind_index[i]:_ChecksumKind_index[i+1]] {
			return enum.ChecksumKind(i + 1)
		}
	}
	panic(fmt.Errorf("unable to locate ChecksumKind enum corresponding to %q", s))
}