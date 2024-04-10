package main

import "github.com/crdzbird/gomedia_length"

func main() {

	gomedia_length.ConvertMediaFormat("/Users/crdzbird/Desktop/peruano.mp4", "/Users/crdzbird/Desktop/", "destination", "mp4")
	gomedia_length.ConvertMediaFormat("/Users/crdzbird/Desktop/peruano.mp4", "/Users/crdzbird/Desktop/", "destination", "mov")
	gomedia_length.ConvertMediaFormat("/Users/crdzbird/Desktop/peruano.mp4", "/Users/crdzbird/Desktop/", "destination", "avi")
	gomedia_length.ConvertMediaFormat("/Users/crdzbird/Desktop/peruano.mp4", "/Users/crdzbird/Desktop/", "destination", "flv")
	gomedia_length.ConvertMediaFormat("/Users/crdzbird/Desktop/peruano.mp4", "/Users/crdzbird/Desktop/", "destination", "wmv")
	gomedia_length.ConvertMediaFormat("/Users/crdzbird/Desktop/peruano.mp4", "/Users/crdzbird/Desktop/", "destination", "webm")
	gomedia_length.ConvertMediaFormat("/Users/crdzbird/Desktop/peruano.mp4", "/Users/crdzbird/Desktop/", "destination", "3gp")
	gomedia_length.ConvertMediaFormat("/Users/crdzbird/Desktop/peruano.mp4", "/Users/crdzbird/Desktop/", "destination", "asf")
	gomedia_length.ConvertMediaFormat("/Users/crdzbird/Desktop/peruano.mp4", "/Users/crdzbird/Desktop/", "destination", "rm")
	gomedia_length.ConvertMediaFormat("/Users/crdzbird/Desktop/peruano.mp4", "/Users/crdzbird/Desktop/", "destination", "swf")
	gomedia_length.ConvertMediaFormat("/Users/crdzbird/Desktop/peruano.mp4", "/Users/crdzbird/Desktop/", "destination", "vob")
	gomedia_length.ConvertMediaFormat("/Users/crdzbird/Desktop/peruano.mp4", "/Users/crdzbird/Desktop/", "destination", "mkv")

	gomedia_length.ConvertMediaFormat("/Users/crdzbird/Desktop/qwerty.mp3", "/Users/crdzbird/Desktop/", "destination", "aac")
	gomedia_length.ConvertMediaFormat("/Users/crdzbird/Desktop/qwerty.mp3", "/Users/crdzbird/Desktop/", "destination", "ac3")
	gomedia_length.ConvertMediaFormat("/Users/crdzbird/Desktop/qwerty.mp3", "/Users/crdzbird/Desktop/", "destination", "flac")
	gomedia_length.ConvertMediaFormat("/Users/crdzbird/Desktop/qwerty.mp3", "/Users/crdzbird/Desktop/", "destination", "mp3")
	gomedia_length.ConvertMediaFormat("/Users/crdzbird/Desktop/qwerty.mp3", "/Users/crdzbird/Desktop/", "destination", "opus")
	gomedia_length.ConvertMediaFormat("/Users/crdzbird/Desktop/qwerty.mp3", "/Users/crdzbird/Desktop/", "destination", "vorbis")
	gomedia_length.ConvertMediaFormat("/Users/crdzbird/Desktop/qwerty.mp3", "/Users/crdzbird/Desktop/", "destination", "wma")
	gomedia_length.ConvertMediaFormat("/Users/crdzbird/Desktop/qwerty.mp3", "/Users/crdzbird/Desktop/", "destination", "pcm")

	//result := gomedia_length.ConvertMediaFormat("/Users/crdzbird/Desktop/qwerty.mp3", "/Users/crdzbird/Desktop/", "destination", "mp3")

}

//Input #0, mov,mp4,m4a,3gp,3g2,mj2, from '/Users/crdzbird/Desktop/peruano.mp4':
// errors: avi, flv, asf
