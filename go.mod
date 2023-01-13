module example

go 1.19

// go modules do not support import <Relative Path>, you should declare it as below.
require gache v0.0.0

require (
	github.com/golang/protobuf v1.5.2 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

replace gache => ./gache
