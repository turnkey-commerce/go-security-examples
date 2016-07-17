package utility

import (
	"bufio"
	"bytes"
	"os"
)

// GetDatastoreString returns a string from the simple file-based datastore.
func GetDatastoreString(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var buffer bytes.Buffer
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		buffer.WriteString(scanner.Text())
	}
	return buffer.String(), scanner.Err()
}

// SaveDatastoreString saves a string to the simple file-based datastore.
func SaveDatastoreString(file string, input string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(input)
	if err != nil {
		return err
	}
	return nil
}

// Check allows to return short error lines in the main functions.
func Check(e error) {
	if e != nil {
		panic(e)
	}
}
