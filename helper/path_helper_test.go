package helper

import (
	"fmt"
	"testing"
)

func TestInitRootPath(t *testing.T) {
	InitRootPath()
	fmt.Println(GetRootPath())
}
