## To install ##

1. Install Go
2. `go get -u github.com/200sc/CBBA-Slideshow`
3. From `$GOPATH/src/gihub.com/200sc/CBBA-Slideshow`, run `go run core.go`

This might not work because it relies on a version of oak that isn't master.
Try `git checkout release/2.0.0` in `$GOPATH/src/github.com/oakmound/oak` to fix that.

Alternatively just look at the slides in the folder named slides.

## To use ##

Move between slides with left and right arrows. ESC closes. 

`c shot` takes a screenshot of the current slide.

`c allShots` takes screenshots of all slides following the current slide.
