package gomedia_length

import (
	"testing"
)

func TestGetVideoDurationWithValidPath(t *testing.T) {
	duration := getVideoDuration("/Users/crdzbird/Desktop/peruano.mp4")
	if duration <= 0 {
		t.Errorf("Expected duration to be greater than 0, got %f", duration)
	}
}

func TestGetVideoDurationWithInvalidPath(t *testing.T) {
	duration := getVideoDuration("/Users/crdzbird/Desktop/noFile.mp4")
	if duration != 0 {
		t.Errorf("Expected duration to be 0 for invalid path, got %f", duration)
	}
}

func TestGetAudioDurationWithValidPath(t *testing.T) {
	duration := getAudioDuration("/Users/crdzbird/Desktop/qwerty.mp3")
	if duration <= 0 {
		t.Errorf("Expected duration to be greater than 0, got %f", duration)
	}
}

func TestGetAudioDurationWithInvalidPath(t *testing.T) {
	duration := getAudioDuration("/Users/crdzbird/Desktop/noMusic.mp3")
	if duration != 0 {
		t.Errorf("Expected duration to be 0 for invalid path, got %f", duration)
	}
}

func TestIsValidMediaFileWithValidPath(t *testing.T) {
	isValid := isValidMediaFile("/Users/crdzbird/Desktop/peruano.mp4")
	if !isValid {
		t.Errorf("Expected file to be valid, got %v", isValid)
	}
}

func TestIsValidMediaFileWithInvalidPath(t *testing.T) {
	isValid := isValidMediaFile("/Users/crdzbird/Desktop/inexistent.mp4")
	if isValid {
		t.Errorf("Expected file to be invalid, got %v", isValid)
	}
}
