# Valid Sudoku

### Prompt:


```
Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.


A Sudoku board (partially filled) could be valid but is not necessarily solvable.
Only the filled cells need to be validated according to the mentioned rules.
```


### Thinking:

First it seems trivial to determine what row and column the algorithim would be in, but square seems more difficult. One solution could just be a look up table, but thinking more `3*(row/3) + (col / 3)` (Came back to this had it wrong originally) should actually yield the right box assuming we go left to right than descend, so:

```
 0 1 2
 3 4 5
 6 7 8
```

Next we can see if we iterate each row it would be simple enough to store values in a map, and so long as there are no collisions that row is good.
To do this for the 9 columns and 9 boxes does mean we would need an extra 18 maps, to make sure we can make a single pass.



### Post Solution Analysis:

Unsuprisingly the silution I have is more memory intense, when compared to others even at the same low runtime percentiles - It seems other strategies opt to due multiple passes focusing on rebuilding a 9 long array of all the fields per row, col or box.

