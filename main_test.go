package gomedia_length

import (
	"fmt"
	"testing"
)

func TestGetVideoDurationWithValidPath(t *testing.T) {
	duration := GetMediaDuration("/Users/crdzbird/Desktop/peruano.mp4")
	if duration <= 0 {
		t.Errorf("Expected duration to be greater than 0, got %f", duration)
	}
}

func TestThumbnailWithValidPath(t *testing.T) {
	result := GenerateThumbnail("/Users/crdzbird/Desktop/peruano.mp4", "/Users/crdzbird/Desktop/result.jpeg", 1200, 1200)
	if result != 1 {
		t.Errorf("Expected thumbnail to be generated, got %d", result)
	}
}

func TestThumbnailWithInvalidPath(t *testing.T) {
	result := GenerateThumbnail("/Users/crdzbird/Desktop/noFile.mp4", "/Users/crdzbird/Desktop/result.jpeg", 1200, 1200)
	if result == 0 {
		t.Errorf("Expected thumbnail to fail with invalid path, got %d", result)
	}
}

func TestGetVideoDurationWithInvalidPath(t *testing.T) {
	duration := GetMediaDuration("/Users/crdzbird/Desktop/noFile.mp4")
	if duration != 0 {
		t.Errorf("Expected duration to be 0 for invalid path, got %f", duration)
	}
}

func TestGetAudioDurationWithValidPath(t *testing.T) {
	duration := GetMediaDuration("/Users/crdzbird/Desktop/qwerty.mp3")
	fmt.Println(duration)
	if duration <= 0 {
		t.Errorf("Expected duration to be greater than 0, got %f", duration)
	}
}

func TestGetAudioDurationWithInvalidPath(t *testing.T) {
	duration := GetMediaDuration("/Users/crdzbird/Desktop/noMusic.mp3")
	if duration != 0 {
		t.Errorf("Expected duration to be 0 for invalid path, got %f", duration)
	}
}

func TestIsValidMediaFileWithValidPath(t *testing.T) {
	isValid := IsValidMediaFile("/Users/crdzbird/Desktop/peruano.mp4")
	if !isValid {
		t.Errorf("Expected file to be valid, got %v", isValid)
	}
}

func TestIsValidMediaFileWithInvalidPath(t *testing.T) {
	isValid := IsValidMediaFile("/Users/crdzbird/Desktop/inexistent.mp4")
	if isValid {
		t.Errorf("Expected file to be invalid, got %v", isValid)
	}
}

func TestConversionAudioWithValidInputs(t *testing.T) {
	result := ConvertMediaFormat("/Users/crdzbird/Desktop/qwerty.mp3", "/Users/crdzbird/Desktop", "converted", "mp3")
	if result != 1 {
		t.Errorf("Expected conversion to be successful, got %d", result)
	}
}

func TestConversionAudioWithInvalidPath(t *testing.T) {
	result := ConvertMediaFormat("invalid_path.mp3", "/Users/crdzbird/Desktop", "converted", "mp3")
	if result == 0 {
		t.Errorf("Expected conversion to fail with invalid path, got %d", result)
	}
}
