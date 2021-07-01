# Tournament
TDD with go, building a small football competition app.

## Features:

Based on an input file containing which team played against which and what the outcome was, the program produces a file with a table like this:

Team                           | MP |  W |  D |  L |  P
Devastating Donkeys            |  3 |  2 |  1 |  0 |  7
Allegoric Alaskans             |  3 |  2 |  0 |  1 |  6
Blithering Badgers             |  3 |  1 |  0 |  2 |  3
Courageous Californians        |  3 |  0 |  1 |  2 |  1
What do those abbreviations mean?

MP: Matches Played
W: Matches Won
D: Matches Drawn (Tied)
L: Matches Lost
P: Points
A win earns a team 3 points. A draw earns 1. A loss earns 0.

The outcome is ordered by points, descending. In case of a tie, teams are ordered alphabetically.

### Input looks like:

Allegoric Alaskans;Blithering Badgers;win
Devastating Donkeys;Courageous Californians;draw
Devastating Donkeys;Allegoric Alaskans;win
Courageous Californians;Blithering Badgers;loss
Blithering Badgers;Devastating Donkeys;loss
Allegoric Alaskans;Courageous Californians;win
The result of the match refers to the first team listed.

So this line:

Allegoric Alaskans;Blithering Badgers;win
Means that the Allegoric Alaskans beat the Blithering Badgers.

This line:

Courageous Californians;Blithering Badgers;loss
Means that the Blithering Badgers beat the Courageous Californians.

And this line:

Devastating Donkeys;Courageous Californians;draw
Means that the Devastating Donkeys and Courageous Californians tied.

## Running the tests
To run the tests run the command `go test`.

The test suite contains benchmarks, you can run these with the --bench and --benchmem flags:

go test -v --bench . --benchmem
