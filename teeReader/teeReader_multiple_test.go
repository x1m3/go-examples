package teeReader

import (
	"bytes"
	"os"
	"log"
	"io"
	"io/ioutil"
	"fmt"
)

// This is not the best performance example.
// What we are showing here is the way of creating multiple buffers to perform different operations with the content
// in this buffers.
// We will launch 2 goroutines that perform a math operation over the content of a file(better said, over a reader),
// opening a file and multiplexing to 2 buffers. We will see how each goroutine is blocked until we start reading
// the original file. At this moment, the content of the file will start flowing to each buffer and each goroutine
// will perform its operation
func ExampleTeeReaderMultiple() {
	var buf1, buf2 bytes.Buffer

	// Opening a file, nothing interesting here
	file, err := os.Open("testdata/numbers.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Chaining buffers. When reading from tee, we will read from file, writing to buf1 and buf2
	tee:= io.TeeReader(io.TeeReader(file, &buf1), &buf2)


	csum := make(chan int)
	cmean := make(chan float64)


	// Will start reading from buf1, but there is no content there right now, so IT WILL BLOCK
	go fileSum(&buf1, csum)

	// Will start reading from buf2, but there is no content there right now, so IT WILL BLOCK
	go fileMean(&buf2, cmean)


	// We read all the tee. All data start flowing to buf1 and buf2
	ioutil.ReadAll(tee)


	// We got the response via a channel
	sum := <-csum
	mean := <-cmean

	fmt.Println(sum)
	fmt.Println(mean)

	//Output:
	// 121
	// 15.125
}

func fileSum(r io.Reader, out chan int)  {
	var acum int = 0
	var new int = 0

	_,err := fmt.Fscanf(r, "%d\n", &new)
	for err==nil  {
		acum += new
		_,err = fmt.Fscanf(r, "%d\n", &new)
	}
	out <- acum
}

func fileMean(r io.Reader, out chan float64)  {
	var acum int = 0
	var new int = 0
	var total int = 0

	_,err := fmt.Fscanf(r, "%d\n", &new)
	for err==nil  {
		acum += new
		total++
		_,err = fmt.Fscanf(r, "%d\n", &new)

	}
	out <- float64(acum) / float64(total)
}

