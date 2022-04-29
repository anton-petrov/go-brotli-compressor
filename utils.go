package main

import (
	"bytes"
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
	defer timeTrack(time.Now(), "Compress")
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
	defer timeTrack(time.Now(), "Deompress")
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
