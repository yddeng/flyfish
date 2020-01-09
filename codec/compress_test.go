package codec

import (
	"fmt"
	"strings"
	"testing"
)

func TestCompress(t *testing.T) {

	func() {

		zipCompressor := &ZipCompressor{}

		zipOut, err := zipCompressor.Compress([]byte("abcdefg"))
		if nil != err {
			t.Fatal(err)
		}

		fmt.Println("zip len", len(zipOut))

		zipUnCompressor := &ZipUnCompressor{}

		unzipOut, err := zipUnCompressor.UnCompress(zipOut)

		if nil != err {
			t.Fatal(err)
		}

		if string(unzipOut) != "abcdefg" {
			t.Fatal(unzipOut)
		}

		unzipOut, err = zipUnCompressor.UnCompress(zipOut)

		if nil != err {
			t.Fatal(err)
		}

		if string(unzipOut) != "abcdefg" {
			t.Fatal(unzipOut)
		}

		fmt.Println(zipUnCompressor.zipBuff.Len(), zipUnCompressor.zipBuff.Cap())
		zipUnCompressor.ResetBuffer()
		fmt.Println(zipUnCompressor.zipBuff.Len(), zipUnCompressor.zipBuff.Cap())

	}()

	func() {

		gzipCompressor := &GZipCompressor{}

		zipOut, err := gzipCompressor.Compress([]byte("abcdefg"))
		if nil != err {
			t.Fatal(err)
		}

		fmt.Println("gzip len", len(zipOut))

		gzipUnCompressor := &GZipUnCompressor{}

		unzipOut, err := gzipUnCompressor.UnCompress(zipOut)

		if nil != err {
			t.Fatal(err)
		}

		if string(unzipOut) != "abcdefg" {
			t.Fatal(unzipOut)
		}

		unzipOut, err = gzipUnCompressor.UnCompress(zipOut)

		if nil != err {
			t.Fatal(err)
		}

		if string(unzipOut) != "abcdefg" {
			t.Fatal(unzipOut)
		}

	}()

}

func BenchmarkZip(b *testing.B) {
	numLoops := b.N
	s := []byte(strings.Repeat("a", 4096))
	zipCompressor := &ZipCompressor{}
	zipUnCompressor := &ZipUnCompressor{}

	for i := 0; i < numLoops; i++ {
		zipOut, err := zipCompressor.Compress(s)
		if nil != err {
			b.Fatal(err)
		}
		_, err = zipUnCompressor.UnCompress(zipOut)
		if nil != err {
			b.Fatal(err)
		}
	}
}

func BenchmarkGZip(b *testing.B) {
	numLoops := b.N
	s := []byte(strings.Repeat("a", 4096))
	zipCompressor := &GZipCompressor{}
	zipUnCompressor := &GZipUnCompressor{}

	for i := 0; i < numLoops; i++ {
		zipOut, err := zipCompressor.Compress(s)
		if nil != err {
			b.Fatal(err)
		}
		_, err = zipUnCompressor.UnCompress(zipOut)
		if nil != err {
			b.Fatal(err)
		}
	}
}
