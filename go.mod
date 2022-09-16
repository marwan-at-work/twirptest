module marwan.io/twirptest

go 1.19

require (
	github.com/twitchtv/twirp v8.1.2+incompatible
	github.com/twitchtv/twirp/v9 v9.0.0-00010101000000-000000000000
	google.golang.org/protobuf v1.28.1
)

require github.com/pkg/errors v0.9.1 // indirect

replace github.com/twitchtv/twirp/v9 => /Users/marwansulaiman/marwan/twirp
