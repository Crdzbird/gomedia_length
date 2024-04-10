package gomedia_length

// #cgo LDFLAGS: -L. -llibrary
// #include "media_library/library.h"
import "C"

func getVideoDuration(path string) float64 {
	return float64(C.getVideoDuration(C.CString(path)))
}

func getAudioDuration(path string) float64 {
	return float64(C.getAudioDuration(C.CString(path)))
}

func isValidMediaFile(path string) bool {
	return C.isValidMediaFile(C.CString(path)) == 1
}
