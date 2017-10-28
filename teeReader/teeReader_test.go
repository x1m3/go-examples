package teeReader_test

import (
	"io"
	"os"
	"strings"
	"io/ioutil"
	"fmt"
)

func ExampleTeeReaderBasic() {
	// This reader could be any Reader, like a file or an http.request.body
	r :=  strings.NewReader("El perro de San Roque no tiene rabo\n")
	w := os.Stdout

	//TeeReader returns a new Reader that reads from r. We can read from tee in te same way we read from r. The
	// interesting here is that everything we read from tee is written again to w, so we are multiplexing the reader
	tee := io.TeeReader(r, w)

	// so, if we read from tee, the same we read will be write to stdout
	out, _ := ioutil.ReadAll(tee)

	fmt.Print(string(out))

	// Output:
	// El perro de San Roque no tiene rabo
	// El perro de San Roque no tiene rabo
}


