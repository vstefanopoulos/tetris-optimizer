# Description

 Tetris-Optimizer is a program that receives only one argument, a path to a text file which will contain a list of tetrominoes and assemble them in order to create the smallest square possible.

## Example of a text File

```
#...
#...
#...
#...

....
....
..##
..##
```

- Result:
```
ABB.
ABB.
A...
A...


$ cat -e sample.txt
...#$
...#$
...#$
...#$
$
....$
....$
....$
####$
$
.###$
...#$
....$
....$
$
....$
..##$
.##.$
....$
$
....$
.##.$
.##.$
....$
$
....$
....$
##..$
.##.$
$
##..$
.#..$
.#..$
....$
$
....$
###.$
.#..$
....$
$ go run . sample.txt | cat -e
ABBBB.$
ACCCEE$
AFFCEE$
A.FFGG$
HHHDDG$
.HDD.G$
$
```

# Author
- Vagelis Stefanopoulos