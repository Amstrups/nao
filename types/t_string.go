// Code generated by "stringer -type=T"; DO NOT EDIT.

package types

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[T_ILLEGAL-0]
	_ = x[T_INT-1]
	_ = x[T_BINARY-2]
	_ = x[T_FLOAT-3]
	_ = x[T_STRING-4]
	_ = x[T_BOOL-5]
	_ = x[T_VOID-6]
	_ = x[T_IDENT-7]
}

const _T_name = "T_ILLEGALT_INTT_BINARYT_FLOATT_STRINGT_BOOLT_VOIDT_IDENT"

var _T_index = [...]uint8{0, 9, 14, 22, 29, 37, 43, 49, 56}

func (i T) String() string {
	if i >= T(len(_T_index)-1) {
		return "T(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _T_name[_T_index[i]:_T_index[i+1]]
}
