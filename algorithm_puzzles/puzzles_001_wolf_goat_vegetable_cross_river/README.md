# Wolf Goat Vegetable Cross River

## Definition

1、F: Farmer 2、W: Wolf 3、G: Goat 4、V: Vegetable, there are four variable can only be 0 or 1, that 0 define left of river and 1 define right of river. We use `S = (F, W, G, V)` define state of puzzle, and begin state is `Ss = (0, 0, 0, 0)`, and final state is `Sf = (1, 1, 1, 1)`.


## Restriction

1、Wolf and goat can't both be in side of river without farmer. we can get formula: `!(W == G && F != W)`
2、Goat and vegetable can't both be in side of river without farmer, we get formula: `!(G == V && F != G)`.

## Action

1、We donate `F(S)` as only pass farmer, that `F(S) => (~F, W, G, V)`
2、We donate `FW(S)` as pass farmer and wolf, that `FW(S) => (~F, ~W, G, V)`
3、We donate `FG(S)` as pass farmer and goat, that `FG(S) => (~F, W, ~G, V)`
4、We donate `FV(S)` as pass farmer and vegetable, that `FV(S) => (~F, W, G, ~V)`


## Pseudo-code
```c
S = (0,0,0,0)
Q = { S }
while !empty(Q) {
    Sp = Q.pop()
    if Sp == (1,1,1,1) {
        // found solve
        break
    }

    VS = {F(Sp), FW(Sp), FG(Sp), FV(Sp)}
    for v in VS {
        if (v.W != v.G && v.F != v.W) && (v.G != v.V && v.F != v.G) {
            Q.push(v)
        }
    }
}
```
