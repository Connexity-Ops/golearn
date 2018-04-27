# Objectivy

Convert a slice of strings into a slice of objects.

## Example
```go
objects = Objectify([]string{"Orange, fruit, 7.5, orange", "Apple, fruit, 6.5, red")
for obj := range objects {
	fmt.Printf("Name: %s, Type: %s, Size: %f, Color: %s", obj.Name, obj.Type, obj.Size, obj.Color)
}
```
Returns:
```go
Name: Orange, Type: fruit, Size: 7.5, Color: orange
Name: Apple, Type: fruit, Size: 6.5, Color: red
```

## Testing
```bash
go test
```
