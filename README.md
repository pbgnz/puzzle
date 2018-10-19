# puzzle [![Build Status](https://travis-ci.org/pbgnz/puzzle.svg?branch=master)](https://travis-ci.org/pbgnz/puzzle) [![Coverage Status](https://coveralls.io/repos/github/pbgnz/puzzle/badge.svg?branch=master)](https://coveralls.io/github/pbgnz/puzzle?branch=master)

## Requirements
1. Go 1.7 or later

## Details

Assume a variation of the 8-puzzle called 11d-puzzle. The 11d-puzzle is identical to the 8-puzzle, except
for 2 differences:
1. the board is a 3x4
2. diagonal moves into the empty tile are legal (assume it can be done on a physical board). So we have
at most 8 possible moves: UP , UP – RIGHT , RIGHT , DOWN – RIGHT , DOWN , DOWN – LEFT , LEFT , UP – LEFT .
The goal configuration of the 11d-puzzle is:

```
+---+---+---+---+
| 1 | 2 | 3 | 4 |
+---+---+---+---+
| 5 | 6 | 7 | 8 |
+---+---+---+---+
| 9 | 10| 11| 0 |
+---+---+---+---+
```

### Usage

``` bash
$ go get -u github.com/pbgnz/puzzle
```

``` bash
$ puzzle 1 0 3 7 5 2 6 4 9 10 11 8
```
