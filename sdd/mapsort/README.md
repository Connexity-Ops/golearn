Map Sort KxV
======

Sorts (ascending) `map[string]int` keys (`string`) by unqiue values (`int`) into a mutable array slice of key/value pairs.

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
