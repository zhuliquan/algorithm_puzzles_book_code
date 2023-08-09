# Rectangle Split Right Triangles

## Description

There are `N` soldiers in a side of the river, and two children on the boat. Boat is very small, can load only two child or one soldier. How many ferries will it take to get all the soldiers from one side of the river to the other?

## Solution

We can reduce the problem to a small problem. Let's think how to take a soldier from on side of the river to the other. We can use four steps to solve this small problem: 1) two children crosses the river first; 2) then a child turned back; 3) and a soldier crossed the river; 4)Finally, the kid on the other side of the river turned back. There will be a total of four ferry times in the above process. Therefore, we can get a generic formula `4 * (N - 1) + 1` (the last soldier can cross the river alone).
