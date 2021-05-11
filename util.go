package go_demo

import (
	"fmt"
	"os"
)

/**
	Write response into a file
**/
func writeData(response string) error {
	f, err := os.Create(OutputFile)
	if err != nil {
		return err
	}

	result, err := f.Write([]byte(response))
	if err != nil {
		return err
	}

	fmt.Printf("wrote %d bytes\n", result)
	if err := f.Close(); err != nil {
		return err
	}

	return nil
}
