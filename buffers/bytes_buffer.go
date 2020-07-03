package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := bytes.NewBuffer(make([]byte, 26))

	texts := []string{"As he came into the window",
		"It was the sound of a cresendo He came into her apartment",
		"He left the bloodstains on the carpet",
		"She ran underneath the table"}

	for i := range texts {
		b.Reset()
		b.WriteString(texts[i])
		fmt.Println("Length:", b.Len(), "\tCapacity:", b.Cap(), "\tText:", b.String())
	}
}
