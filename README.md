# perfectcuboid

perfectcuboid is a Go implementation of an algorithm to attempt to find a perfect cuboid.

## Installation

`go get github.com/kochman/perfectcuboid`

perfectcuboid should run anywhere Go runs, though it has only been tested on OS X.

## Usage

`perfectcuboid -max=40`

perfectcuboid will search for perfect cuboids with a minimum side length of 1 and a maximum side length of 40. If a perfect cuboid is found, its dimensions will be printed. The maximum side length can be changed, though a bigger number results in longer computation times than a smaller number.

`perfectcuboid -max=100 -min=41`

perfectcuboid will search for perfect cuboids with a minimum side length of 41 and a maximum side length of 100. The `-min` argument is good for starting a search from a previous maximum side length. In this case, the above two commands will have tried every cuboid with integer side lengths from 1 to 100 without trying the same cuboid twice.
