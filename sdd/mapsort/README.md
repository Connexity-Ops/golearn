Map Sort KxV
======

Sorts a `map[string]int`'s keys [`string`] by unique values [`int`], in ascending order, into a mutable array slice of key/value pairs.

## Example
```go
fmt.Println(MapSort(potusNum))
```
Returns:
```go
[{George Washington 1} {Thomas Jefferson 3} {Abraham Lincoln 16} {Franklin Delano Roosevelt 32} {John Fitzgerald Kennedy 35} {Barack Hussein Obama 44}]
```

## Testing
```
 go test
```
