package gomedia_length

// #cgo LDFLAGS: -L. -llibrary
// #include "media_library/library.h"
import "C"

func GetMediaDuration(path string) float64 {
	return float64(C.getMediaDuration(C.CString(path)))
}

func IsValidMediaFile(path string) bool {
	return C.isValidMediaFile(C.CString(path)) == 1
}
