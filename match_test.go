package filetype

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/haxii/filetype/matchers"
	"github.com/haxii/filetype/types"
)

func TestMatch(t *testing.T) {
	cases := []struct {
		buf []byte
		ext string
	}{
		{[]byte{0x38, 0x42, 0x50, 0x53}, "psd"},
		{[]byte{0xFF, 0xD8, 0x00}, "unknown"},
		{[]byte{0x89, 0x50, 0x4E, 0x47}, "png"},
	}

	for _, test := range cases {
		match, err := Match(test.buf)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}

		if match.Extension != test.ext {
			t.Fatalf("Invalid image type: %s != %s", match.Extension, test.ext)
		}
	}
}

func TestMatchDocumentFile(t *testing.T) {
	cases := []struct {
		ext string
	}{
		{"xls"},
		{"ppt"},
	}

	for _, test := range cases {
		kind, _ := MatchFile("./test_files/documents/testsrc." + test.ext)
		if test.ext == "ppt" {
			if kind.Extension != "xls" {
				t.Fatalf("Invalid image type: %s != %s", kind.Extension, "xls")
			}
			continue
		}
		if kind.Extension != test.ext {
			t.Fatalf("Invalid image type: %s != %s", kind.Extension, test.ext)
		}
	}
}

func TestMatchArchiveFile(t *testing.T) {
	cases := []struct {
		ext string
	}{
		{"xz"},
		{"zip"},
		{"tar"},
		{"gz"},
		{"bz2"},
		{"7z"},
		{"pdf"},
		{"bz2"},
		{"cab"},
		{"crx"},
		{"ar"},
		{"exe"},
		{"lz"},
		{"nes"},
		{"ps"},
		{"rar"},
		{"rpm"},
		{"rtf"},
		{"swf"},
	}

	for _, test := range cases {
		kind, _ := MatchFile("./test_files/archive/testsrc." + test.ext)
		if test.ext == "elf" {
			file, _ := os.Open("./test_files/archive/testsrc." + test.ext)
			buffer := make([]byte, 512)
			file.Read(buffer)
			fmt.Printf("%X", buffer[0:16])
			continue
		}
		if kind.Extension != test.ext {
			t.Fatalf("Invalid image type: %s != %s", kind.Extension, test.ext)
		}
	}
}

func TestMatchFontFile(t *testing.T) {
	cases := []struct {
		ext string
	}{
		{"otf"},
		{"ttf"},
		{"woff"},
		{"woff2"},
	}

	for _, test := range cases {
		kind, _ := MatchFile("./test_files/font/testsrc." + test.ext)
		if kind.Extension != test.ext {
			t.Fatalf("Invalid image type: %s != %s", kind.Extension, test.ext)
		}
	}
}

func TestMatchAudioFile(t *testing.T) {
	cases := []struct {
		ext string
	}{
		{"aac"},
		{"caf"},
		{"flac"},
		{"m4a"},
		{"mp3"},
		{"ogg"},
		{"ra"},
		{"rm"},
		{"voc"},
		{"wav"},
		{"wmv"},
	}

	for _, test := range cases {
		kind, _ := MatchFile("./test_files/audio/output." + test.ext)
		if test.ext == "ra" || test.ext == "rm" || test.ext == "rmvb" {
			if kind.Extension != "ra" && kind.Extension != "rm" && kind.Extension != "rmvb" {
				t.Fatalf("Invalid audio type: %s != %s", kind.Extension, test.ext)
			}
			continue
		}
		if test.ext == "asf" || test.ext == "wmv" {
			if kind.Extension != "asf" && kind.Extension != "wmv" {
				t.Fatalf("Invalid audio type: %s != %s", kind.Extension, test.ext)
			}
			continue
		}
		if kind.Extension != test.ext {
			t.Fatalf("Invalid audio type: %s != %s", kind.Extension, test.ext)
		}
	}
}

func TestMatchImageFile(t *testing.T) {
	cases := []struct {
		ext string
	}{
		{"bmp"},
		{"gif"},
		{"ico"},
		{"png"},
		{"psd"},
		{"tif"},
		{"tiff"},
		{"jpg"},
	}

	for _, test := range cases {
		kind, _ := MatchFile("./test_files/images/download." + test.ext)
		if test.ext == "tiff" || test.ext == "tif" {
			if kind.Extension != "tiff" && kind.Extension != "tif" {
				t.Fatalf("Invalid image type: %s != %s", kind.Extension, test.ext)
			}
			continue
		}
		if test.ext == "jpg" || test.ext == "jpe" || test.ext == "jpeg" || test.ext == "jfif" {
			if kind.Extension != "jpg" && kind.Extension != "jpe" && kind.Extension != "jpeg" && kind.Extension != "jfif" {
				t.Fatalf("Invalid image type: %s != %s", kind.Extension, test.ext)
			}
			continue
		}
		if kind.Extension != test.ext {
			t.Fatalf("Invalid image type: %s != %s", kind.Extension, test.ext)
		}
	}
}

func TestMatchVedioFile(t *testing.T) {
	cases := []struct {
		ext string
	}{
		{"3gp"},
		{"asf"},
		{"avi"},
		{"flv"},
		{"m4v"},
		{"mkv"},
		{"mov"},
		{"mp4"},
		{"mpg"},
		{"vob"},
	}

	for _, test := range cases {
		kind, _ := MatchFile("./test_files/vedio/testsrc." + test.ext)
		if test.ext == "mkv" {
			if kind.Extension != "webm" && kind.Extension != "mkv" {
				t.Fatalf("Invalid video type: %s != %s", kind.Extension, test.ext)
			}
			continue
		}
		if test.ext == "3gp" {
			if kind.Extension != "3g2" && kind.Extension != "3gg" && kind.Extension != "3gp" {
				t.Fatalf("Invalid video type: %s != %s", kind.Extension, test.ext)
			}
			continue
		}
		if test.ext == "asf" {
			if kind.Extension != "asf" && kind.Extension != "wmv" {
				t.Fatalf("Invalid video type: %s != %s", kind.Extension, test.ext)
			}
			continue
		}
		if test.ext == "mpg" || test.ext == "mepg" || test.ext == "vob" {
			if kind.Extension != "mpg" && kind.Extension != "mpeg" && kind.Extension != "vob" {
				t.Fatalf("Invalid video type: %s != %s", kind.Extension, test.ext)
			}
			continue
		}
		if kind.Extension != test.ext {
			t.Fatalf("Invalid video type: %s != %s", kind.Extension, test.ext)
		}
	}
}

func TestMatchReader(t *testing.T) {
	cases := []struct {
		buf io.Reader
		ext string
	}{
		{bytes.NewBuffer([]byte{0xFF, 0xD8, 0x00}), "unknown"},
		{bytes.NewBuffer([]byte{0x89, 0x50, 0x4E, 0x47}), "png"},
	}

	for _, test := range cases {
		match, err := MatchReader(test.buf)
		if err != nil {
			t.Fatalf("Error: %s", err)
		}

		if match.Extension != test.ext {
			t.Fatalf("Invalid image type: %s", match.Extension)
		}
	}
}

func TestMatches(t *testing.T) {
	cases := []struct {
		buf   []byte
		match bool
	}{
		{[]byte{0xFF, 0xD8, 0xFF}, true},
		{[]byte{0xFF, 0x0, 0x0}, false},
		{[]byte{0x89, 0x50, 0x4E, 0x47}, true},
	}

	for _, test := range cases {
		if Matches(test.buf) != test.match {
			t.Fatalf("Do not matches: %#v", test.buf)
		}
	}
}

func TestAddMatcher(t *testing.T) {
	fileType := AddType("foo", "foo/foo")

	AddMatcher(fileType, func(buf []byte) bool {
		return len(buf) == 2 && buf[0] == 0x00 && buf[1] == 0x00
	})

	if !Is([]byte{0x00, 0x00}, "foo") {
		t.Fatalf("Type cannot match")
	}

	if !IsSupported("foo") {
		t.Fatalf("Not supported extension")
	}

	if !IsMIMESupported("foo/foo") {
		t.Fatalf("Not supported MIME type")
	}
}

func TestMatchMap(t *testing.T) {
	cases := []struct {
		buf  []byte
		kind types.Type
	}{
		{[]byte{0x89, 0x50, 0x4E, 0x47}, types.Get("png")},
		{[]byte{0xFF, 0x0, 0x0}, Unknown},
	}

	for _, test := range cases {
		kind := MatchMap(test.buf, matchers.Image)
		if kind != test.kind {
			t.Fatalf("Do not matches: %#v", test.buf)
		}
	}
}

func TestMatchesMap(t *testing.T) {
	cases := []struct {
		buf   []byte
		match bool
	}{
		{[]byte{0xFF, 0xD8, 0xFF}, true},
		{[]byte{0x89, 0x50, 0x4E, 0x47}, true},
		{[]byte{0xFF, 0x0, 0x0}, false},
	}

	for _, test := range cases {
		if match := MatchesMap(test.buf, matchers.Image); match != test.match {
			t.Fatalf("Do not matches: %#v", test.buf)
		}
	}
}

//
// Benchmarks
//

var tarBuffer, _ = ioutil.ReadFile("./fixtures/sample.tar")
var zipBuffer, _ = ioutil.ReadFile("./fixtures/sample.zip")
var jpgBuffer, _ = ioutil.ReadFile("./fixtures/sample.jpg")
var gifBuffer, _ = ioutil.ReadFile("./fixtures/sample.gif")
var pngBuffer, _ = ioutil.ReadFile("./fixtures/sample.png")
var m4vBuffer, _ = ioutil.ReadFile("./fixtures/sample.m4v")
var tifBuffer, _ = ioutil.ReadFile("./fixtures/sample.tif")

func BenchmarkMatchTar(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Match(tarBuffer)
	}
}

func BenchmarkMatchZip(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Match(zipBuffer)
	}
}

func BenchmarkMatchJpeg(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Match(jpgBuffer)
	}
}

func BenchmarkMatchGif(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Match(gifBuffer)
	}
}

func BenchmarkMatchPng(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Match(pngBuffer)
	}
}

func BenchmarkMatchTif(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Match(tifBuffer)
	}
}

func BenchmarkMatchM4v(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Match(m4vBuffer)
	}
}
