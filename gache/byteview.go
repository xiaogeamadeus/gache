package gache

// ByteView is a struct holds an immutable view of bytes (read only).
type ByteView struct {
	byteArray []byte
}

// Len returns the view's length
func (byteView ByteView) Len() int {
	return len(byteView.byteArray)
}

// ByteSlice returns a copy of the data as a byte slice
func (byteView ByteView) ByteSlice() []byte {
	return cloneBytes(byteView.byteArray)
}

// ToString returns the data as a string, make a copy if necessary
func (byteView ByteView) ToString() string {
	return string(byteView.byteArray)
}

func cloneBytes(oldByteArray []byte) []byte {
	newByteArray := make([]byte, len(oldByteArray))
	copy(newByteArray, oldByteArray)
	return newByteArray
}
