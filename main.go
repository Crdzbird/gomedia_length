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

func ConvertMediaFormat(path, destPath, outputFile, outputFormat string) int {
	return int(C.convertMediaFormat(C.CString(path), C.CString(destPath), C.CString(outputFile), C.CString(outputFormat)))
}
