package main

import (
	//	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"path/filepath"
	"strings"
	"time"
	// . "github.com/andybalholm/brotli"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	// Configure flags
	compress := ""
	decompress := ""
	output := ""
	test := 0
	quality := 6
	flag.StringVar(&compress, "c", "", "compress a file")
	flag.StringVar(&decompress, "d", "", "decompress a file")
	flag.StringVar(&output, "o", "", "output file")
	flag.IntVar(&test, "t", 0, "buffer size for benchmark test (megabytes)")
	flag.IntVar(&quality, "q", 6, "compression quality (BestSpeed - 0, DefaultCompression - 6, BestCompression - 11)")
	flag.Parse()

	var input_data []byte
	var output_data []byte
	var input string

	if compress != "" {
		input = compress
		input_data = ReadFile(input)
		println("Compressing data...")
		output_data, _ = Compress(input_data, quality)
		ioutil.WriteFile(input+".brotli", output_data, 0777)
	} else if decompress != "" {
		input = decompress
		input_data = ReadFile(input)
		println("Desompressing data...")
		output_data, _ = Decompress(input_data)
		ioutil.WriteFile(strings.TrimSuffix(input, filepath.Ext(input)), output_data, 0777)
	} else if test != 0 {
		println("Preparing benchmark...")
		input_data = gen_rand_text(test * 100000)
		println("Compressing data") // bytes.NewBuffer(input_data).String()
		output_data, _ = Compress(input_data, quality)
		println("Desompressing data")
		Decompress(output_data)
	} else {
		log.Fatal("You must specify either compress or decompress")
	}

	println("Input data length:", len(input_data), "Output data length:", len(output_data))
	fmt.Printf("Ratio: %f\n", float64(len(output_data))/float64(len(input_data)))
	println("Go!")
}
