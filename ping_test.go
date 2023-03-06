package hypolashlckping

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

// TestHTTPApi get API et test result
func TestPing(t *testing.T) {
	result := Call()

	if !result.IsUP {
		log.Err.Println("Is False")
	} else {
		log.Err.Println("OK")
	}

	readFile, err := os.Open(log.LogFile.Name())

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}

	readFile.Close()
}
