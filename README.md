<h1 align="center">TETRIS-OPTIMIZER</h1>

## About The Project
Tetris-optimizer is a program that receives only one argument, a path to a text file which contains a list of tetrominoes to assemble them in order to create the smallest square possible.

## Installation
```
git clone https://github.com/SpectreH/tetris-optimizer.git
cd tetris-optimizer
```

## Usage
```
go run . [FILE]
```

## Examples
```
cat tests/sample.txt
...#
...#
...#
...#

....
....
....
####

.###
...#
....
....

....
..##
.##.
....

....
.##.
.##.
....

....
....
##..
.##.

##..
.#..
.#..
....

....
###.
.#..
....

go run . tests/sample.txt
ABBBB.
ACCCEE
AFFCEE
A.FFGG
HHHDDG
.HDD.G
```

## Additional information

Only standard go packages were in use. In <code>tests</code> folder you can find several presets to generate squares and in <code>error-tests</code> folder some bad presets examples.

To go through all tests, you can use:
```
bash test.sh
```

## Author

* SpectreH (https://github.com/SpectreH)
