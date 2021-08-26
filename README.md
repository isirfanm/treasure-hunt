# Task 2: Treasure Hunt

## Problem Description

Build a simple command-line program for helping the user hunt for a treasure that satisfies the following requirements: 

1. Build a simple grid with the following layout:

    ```
    ########
    #......#
    #.###..#
    #...#.##
    #X#....#
    ########
    ```

    - \# represents an obstacle.
    - . represents a clear path.
    - X represents the player’s starting position.

2. A treasure is hidden within one of the clear path points, and the user must find it.
3. From the starting position, the user must navigate in a specific order:
    a. Up/North A step(s), then
    b. Right/East B step(s), then
    c. Down/South C step(s).
4. The program must output a list of probable coordinate points where the treasure might be located.
5. Bonus points: display the grid with all the probable treasure locations marked with a $ symbol.

## Run Program

Run program by following command.

```
$ go run .
```

The Program will run and show output. The output will be (1) Initial map, (2) Possible treasure posisitons, and (3) Final marked map.

```
########
#......#
#.###..#
#...#.##
#X#....#
########
(3, 4)
(5, 2)
(5, 3)
(5, 4)
(6, 2)
########
#......#
#.###$$#
#...#$##
#X#$.$.#
########
```
