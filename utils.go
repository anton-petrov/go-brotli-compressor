package main

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	. "github.com/andybalholm/brotli"
	"github.com/tjarratt/babble"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func gen_rand_bytes(size int) []byte {
	random_data := make([]byte, size)
	rand.Read(random_data)
	return random_data
}

func gen_rand_text(size int) []byte {
	babbler := babble.NewBabbler()
	babbler.Separator = " "
	babbler.Count = size
	return []byte(babbler.Babble())
}

// Compress returns content encoded with Brotli.
func Compress(content []byte, quality int) ([]byte, error) {
	defer timeTrack(time.Now(), "Compress Brotli")
	var buf bytes.Buffer
	writer := NewWriterOptions(&buf, WriterOptions{Quality: quality})
	_, err := writer.Write(content)
	if closeErr := writer.Close(); err == nil {
		err = closeErr
	}
	return buf.Bytes(), err
}

func SimpleBenchmark(size_bytes int) {

}

// Decompressing Brotli encoded data.
func Decompress(encodedData []byte) ([]byte, error) {
	defer timeTrack(time.Now(), "Decompress Brotli")
	r := NewReader(bytes.NewReader(encodedData))
	return ioutil.ReadAll(r)
}

// Read file
func ReadFile(file string) []byte {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	return data
}


func CompressGzip(data []byte) ([]byte, error) {
	defer timeTrack(time.Now(), "Compress Gzip")
	var compressed_data bytes.Buffer
	zw := gzip.NewWriter(&compressed_data)

	// Setting the Header fields is optional.
	// zw.Name = ""
	// zw.Comment = ""
	// zw.ModTime = time.Date(1977, time.May, 25, 0, 0, 0, 0, time.UTC)

	_, err := zw.Write(data)

	if err != nil {
		log.Println(err)
		return data, err
	}
	if err := zw.Close(); err != nil {
		log.Println(err)
		return data, err
	}

	return compressed_data.Bytes(), err
}

func DecompressGzip(data []byte) ([]byte, error) {
	defer timeTrack(time.Now(), "Decompress Gzip")
	var compressed_data = bytes.NewBuffer(data)
	var decompressed_data bytes.Buffer

	zr, err:= gzip.NewReader(compressed_data)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	
	if _, err := io.Copy(&decompressed_data, zr); err != nil {
		log.Println(err)
		return data, err
	}

	if err := zr.Close(); err != nil {
		log.Println(err)
		return nil, err
	}

	return decompressed_data.Bytes(), err
}