# mr-roboto

[![Build Status](https://travis-ci.org/gugahoi/mr-roboto.svg?branch=master)](https://travis-ci.org/gugahoi/mr-roboto)
[![Coveralls github](https://img.shields.io/coveralls/github/gugahoi/mr-roboto.svg)](https://coveralls.io/github/gugahoi/mr-roboto?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/gugahoi/mr-roboto)](https://goreportcard.com/report/github.com/gugahoi/mr-roboto)
[![Maintainability](https://api.codeclimate.com/v1/badges/5b997906d998fa9d104e/maintainability)](https://codeclimate.com/github/gugahoi/mr-roboto/maintainability)

## Challenge

The challenge is to simulate a toy robot moving on a square board of 6 x 6 units. The robot can roam
around the surface of the board but shouldn’t be able to fall off the edge. The twist is you are able to place
any number of robots on the board and they move independently of each other. Any movement that would
result in the robot falling from the board or colliding with another robot must be prevented, however further
valid movement commands must still be allowed.

* Input can be from a file or from standard input
* PLACE will put a toy robot on the board in position X,Y and facing NORTH, SOUTH, EAST or WEST.
* The origin (0,0) can be considered to be the SOUTH WEST most corner.
* The first valid command for a robot is a PLACE command, after that, any sequence of commands may be issued, in any order, including another PLACE command. The application should discard all commands in the sequence for a given robot until a valid PLACE command has been executed
* MOVE will move a toy robot one unit forward in the direction it is currently facing.
* LEFT and RIGHT will rotate a robot 90 degrees in the specified direction without changing the position of the robot.
* REPORT will announce the X,Y and F of a robot
* A toy robot must not hit the edges of the board during movement, including when the robot is initially placed on the board.
* Any move that would cause the robot to fall off the edge of the board or collide with another robot must be ignored.

## Downloading

To run the latest release, simply download the executable from the [GiHub Releases page](https://github.com/gugahoi/mr-roboto/releases/latest), make it executable and run the binary:

```bash
# with an input file
mr-roboto testdata/example1.txt
# from stdin
mr-roboto
```

## Development

```bash
# Install go: https://golang.org/doc/install

# clone the repo into your go path src directory
git clone git@github.com:gugahoi/mr-roboto $GOPATH/src/github.com/gugahoi/mr-roboto

# build the binary
# go build -o build/mr-roboto ./...

# or simply use the make task
make build

# the binary is now under the build directory, run it to get started:
# when given an argument, it will try to read it as the input file
build/mr-roboto testdata/example1.txt

# when given no arguments it will wait for input in stdin for each command
build/mr-roboto
> ANA: PLACE 1,1,NORTH
> ANA: REPORT
ANA: 1,1,NORTH
```

Alternatively you can build a scratch docker image with the binary and run it:

```bash
make docker-run
> ANA: PLACE 1,1,NORTH
> ANA: REPORT
ANA: 1,1,NORTH

# or with an example file
docker run -v $(pwd)/testdata/:/data/ gugahoi/mr-roboto:latest /data/example1.txt
ALICE: 0,1,NORTH
```

## Testing

Unit tests are found under the `src/*_test.go` files as per the standard golang convention.
You can run them with the standard commands

```bash
go test ./...
ok      github.com/gugahoi/mr-roboto/src        0.002s
```
