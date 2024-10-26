package gomedia_length

// #cgo LDFLAGS: -L. -llibrary
// #include "media_library/library.h"
import "C"
import (
	"unsafe"
)

func GetMediaDuration(path string) float64 {
	return float64(C.getMediaDuration(C.CString(path)))
}

func IsValidMediaFile(path string) bool {
	return C.isValidMediaFile(C.CString(path)) == 1
}

func ConvertMediaFormat(path, destPath, outputFile, outputFormat string) int {
	return int(C.convertMediaFormat(C.CString(path), C.CString(destPath), C.CString(outputFile), C.CString(outputFormat)))
}

func GenerateThumbnail(path, destPath, outputFile, outputFormat string, width, height int) int {
	return int(C.generateThumbnail(C.CString(path), C.CString(destPath), C.CString(outputFile), C.CString(outputFormat), C.int(width), C.int(height)))
}

func GenerateThumbnails(path, destPath string, width, height, totalThumbnails int) []string {
	return convertCArrayToStringArray(C.generateThumbnails(C.CString(path), C.CString(destPath), C.int(width), C.int(height), C.int(totalThumbnails)))
}

func convertCArrayToStringArray(cArray **C.char) []string {
	var result []string
	for {
		// Dereference the pointer to get the C string
		cStr := *cArray
		if cStr == nil {
			break
		}
		// Convert C string to Go string
		goStr := C.GoString(cStr)
		result = append(result, goStr)
		// Move to the next element in the array
		cArray = (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(cArray)) + unsafe.Sizeof(cArray)))
	}
	return result
}
