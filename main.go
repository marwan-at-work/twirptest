package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"marwan.io/twirptest/twirp9"
)

func main() {
	log.Fatal(http.ListenAndServe(":5959", twirp9.NewHaberdasherServer(&server{})))
}

type server struct {
}

// MakeHat implements twirp9.Haberdasher
func (*server) MakeHat(ctx context.Context, req *twirp9.Size) (*twirp9.Hat, error) {
	fmt.Println("GOT", req.Inches)
	return &twirp9.Hat{
		Size:  123,
		Color: "abc",
		Name:  "qwer",
	}, nil
}
