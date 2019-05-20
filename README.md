# MarsRover

## Example input

 8 10
 
 1 2  E
 
 MMLMRMMRRMML
     
     
## This is how the rover will interpret the above commands:
  * The first line describes how big the current zone is. This zone’s boundary is at
    the Cartesian coordinate of 8,10 and the zone comprises 80 blocks.
  * The second line describes the rover’s starting location and orientation.
  * This rover would start at position 1 on the horizontal axis, position 2 on the
    vertical axis and would be facing East (E). The third line is the list of commands
    (movements and rotations) to be executed by the rover.
  * As a result of following these commands, a rover starting at 1 2 E in this zone
    would land up at 3 3 S.
