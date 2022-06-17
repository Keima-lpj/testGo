package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"strings"
)

func main() {
	id := uuid.NewV4()
	fmt.Println(strings.Replace(id.String(), "-", "", -1))
}
