package main

import (
	"context"
	"fmt"
	"net/http"

	"marwan.io/twirptest/twirp8"
)

func main() {
	c := twirp8.NewHaberdasherProtobufClient("http://localhost:5959", http.DefaultClient)
	resp, err := c.MakeHat(context.Background(), &twirp8.Size{
		Inches: 456,
	})
	must(err)
	fmt.Println("GOT", resp.Color, resp.Name, resp.Size)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
