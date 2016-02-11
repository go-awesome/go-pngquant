package pngquant

import (
	"testing"
	"image/png"
	"os"
	"bytes"
)

func TestCompress(t *testing.T) {
	var file *os.File
	file, _ = os.Open("gopher.png")
	defer file.Close()
	info, _ := file.Stat()
	orgSize := info.Size()
	orgImg, _ := png.Decode(file)
	newImg, err := Compress(orgImg)
	if err != nil {
		t.Errorf("error has occurred: %v", err)
	}
	var w bytes.Buffer
	png.Encode(&w, newImg)
	if len(w.Bytes()) > int(orgSize) {
		t.Error("image is not comppressed")
	}
}
