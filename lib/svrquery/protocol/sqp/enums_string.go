// Code generated by "stringer -type=DataType -output=enums_string.go"; DO NOT EDIT.

package sqp

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Byte-0]
	_ = x[Uint16-1]
	_ = x[Uint32-2]
	_ = x[Uint64-3]
	_ = x[String-4]
	_ = x[Float32-5]
}

const _DataType_name = "ByteUint16Uint32Uint64StringFloat32"

var _DataType_index = [...]uint8{0, 4, 10, 16, 22, 28, 35}

func (i DataType) String() string {
	if i >= DataType(len(_DataType_index)-1) {
		return "DataType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _DataType_name[_DataType_index[i]:_DataType_index[i+1]]
}
