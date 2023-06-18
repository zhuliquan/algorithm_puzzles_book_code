# Rectangle Split Right Triangles

## Description
how to find a way to split a rectangle to `N` right triangles (`N>1`).

## Solution

1、case1: `N == 2 * k, k is natural number`, split a rectangle to `k` a rectangle, and split each rectangle to two right triangle.

2、case2: `N == 2 * k + 1, k is natural number`, split a rectangle to `2 * k` right triangles (case1). and split a right triangle to two right triangle.

