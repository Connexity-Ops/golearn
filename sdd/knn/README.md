KNN
======

Classify samples using the k-nearest neighbors algorithm


## Example
```go
knn := NewKNNClassifier(4)
trainingSet := []*Sample{
    NewSample([]float64{0.0, 0.0}, "red"),
    NewSample([]float64{1.0, 1.0}, "red"),
    NewSample([]float64{10.0, 10.0}, "blue"),
    NewSample([]float64{11.0, 11.0}, "blue"),
}
knn.Fit(trainingSet)
unclassifiedSample := NewSample([]float64{10.5, 10.5}, "blue")
prediction := knn.Predict(unclassifiedSample)
fmt.Println(prediction)
```
Returns:
```go
blue
```

## Testing
```
 go test
 go test -bench=.
```

## Help

Try these guides to understand KNN and how to implement it:

- [KNN overview video](https://www.youtube.com/watch?v=MDniRwXizWo)
- [KNN article with pseudocode](http://dataaspirant.com/2016/12/23/k-nearest-neighbor-classifier-intro/)

Here's more info on the training dataset used:

[Iris Data Set](http://archive.ics.uci.edu/ml/datasets/Iris)
