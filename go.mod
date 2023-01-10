module example

go 1.19

// go modules do not support import <Relative Path>, you should declare it as below.
require gache v0.0.0

replace gache => ./gache
