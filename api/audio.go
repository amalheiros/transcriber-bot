package bot

import (
	"path/filepath"
	"strings"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func convertToMP3(file string) (string, error) {
	path := file
	dir, file := filepath.Split(path)

	fileWithoutSuffix := strings.TrimSuffix(file, filepath.Ext(file))

	newPath := dir + fileWithoutSuffix + ".mp3"

	err := ffmpeg.Input(dir+file).
		Output(newPath, ffmpeg.KwArgs{"c:v": "libx265"}).
		OverWriteOutput().ErrorToStdOut().Run()

	return newPath, err

}
