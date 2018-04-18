package knn

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strconv"
	"testing"
)

func readSamples(path string) ([]*Sample, error) {
	csvFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var samples []*Sample
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		attributes := make([]float64, len(line)-1)
		for i := 0; i < len(line)-1; i++ {
			attributes[i], err = strconv.ParseFloat(line[i], 64)
			if err != nil {
				return nil, err
			}
		}

		samples = append(samples, NewSample(attributes, line[4]))
	}
	return samples, nil
}

func TestKNNClassifier(t *testing.T) {
	trainingSet, err := readSamples("testdata/train.csv")
	if err != nil {
		t.Fatalf("Could not read training data: %s", err.Error())
	}

	testingSet, err := readSamples("testdata/test.csv")
	if err != nil {
		t.Fatalf("Could not read test data: %s", err.Error())
	}

	knn := NewKNNClassifier(4)
	knn.Fit(trainingSet)
	for _, sample := range testingSet {
		if got := knn.Predict(sample); !reflect.DeepEqual(got, sample.Class) {
			t.Errorf("knn.Predict(s Sample) = %v, want %v", got, sample.Class)
		}
	}
}

func ExampleKNNClassifier() {
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
	// Output: blue
}

func BenchmarkKNNClassifier(b *testing.B) {
	trainingSet, err := readSamples("testdata/train.csv")
	if err != nil {
		b.Fatalf("Could not read training data: %s", err.Error())
	}

	testingSet, err := readSamples("testdata/test.csv")
	if err != nil {
		b.Fatalf("Could not read test data: %s", err.Error())
	}

	for n := 0; n < b.N; n++ {
		knn := NewKNNClassifier(4)
		knn.Fit(trainingSet)
		for _, sample := range testingSet {
			knn.Predict(sample)
		}
	}
}
