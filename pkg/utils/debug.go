package utils

import (
	"github.com/davecgh/go-spew/spew"
	"os"
)

func Pre(x interface{}, y ...interface{}) {
	spew.Dump(x)
	os.Exit(1)
}
