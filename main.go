package main

import (
	"bytes"
	"log"

	"github.com/philpearl/scratchbuild"
)

func main() {
	o := scratchbuild.Options{
		Dir:      "./testdata",
		Name:     "test",
		BaseURL:  "http://localhost:5000",
		Tag:      "latest",
	}

	b := &bytes.Buffer{}
	if err := scratchbuild.TarDirectory("./testdata", b); err != nil {
		log.Fatalf("failed to tar layer. %s", err)
	}

	c := scratchbuild.New(&o)

	if err := c.BuildImage(&scratchbuild.ImageConfig{
		Entrypoint: []string{"/app"},
	}, b.Bytes()); err != nil {
		log.Fatalf("failed to build and send image. %s", err)
	}
}