// Package knn implements a K nearest neighbors classifier.

package knn

import (
	"math"
	"sort"
)

type Sample struct {
	Attributes []float64
	Class      string
}

// Distance calculates the euclidean distance from one sample to another.
// SDD work
func (s *Sample) Distance(other *Sample) float64 {
	//Use this formula: d = (x1 - x2)^2 + (y1 - y2)^2 + (z1 - z2)^2
	eucdist := 0.0
	for i := 0; i < len(other.Attributes); i++ {
		eucdist += math.Pow(s.Attributes[i]-other.Attributes[i], 2)
	}
	return eucdist
}

// NewSample creates the object with the data points.
func NewSample(attributes []float64, class string) *Sample {
	return &Sample{Attributes: attributes, Class: class}
}

type KNNClassifier struct {
	K            int
	TrainingData []*Sample
}

func NewKNNClassifier(k int) *KNNClassifier {
	return &KNNClassifier{K: k}
}

// Fit method stores the slice of samples in the struct
// SDD work
func (knn *KNNClassifier) Fit(samples []*Sample) {

	knn.TrainingData = samples
}

// SampleWithDist appends a distance field onto the Sample struct
// SDD work
type SampleWithDist struct {
	*Sample
	Distance float64
}

// Following sections used for struct value sort.
// SDD work
type sampleSorter struct {
	distances []SampleWithDist
	by        func(p1, p2 *SampleWithDist) bool // Closure used in the Less method.
}

// By is the type of a "less" function that defines the ordering of its SampleWithDist arguments.
type By func(p1, p2 *SampleWithDist) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(distances []SampleWithDist) {
	ps := &sampleSorter{
		distances: distances,
		by:        by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}

// Len is part of sort.Interface.
func (s *sampleSorter) Len() int {
	return len(s.distances)
}

// Swap is part of sort.Interface.
func (s *sampleSorter) Swap(i, j int) {
	s.distances[i], s.distances[j] = s.distances[j], s.distances[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *sampleSorter) Less(i, j int) bool {
	return s.by(&s.distances[i], &s.distances[j])
}

// End required code section for sort.

// Predict is where the work happens
func (knn *KNNClassifier) Predict(sample *Sample) string {

	//1. Get the distance from sample to every other sample in knn.TrainingData using Distance func.
	distances := []SampleWithDist{}
	for _, sampleGroup := range knn.TrainingData {
		distances = append(distances, SampleWithDist{Sample: sampleGroup,
			Distance: sample.Distance(sampleGroup)})
	}

	//2. Sort by closest samples based on calculated distances
	distance := func(p1, p2 *SampleWithDist) bool {
		return p1.Distance < p2.Distance
	}

	By(distance).Sort(distances)
	//fmt.Println("By distance:", distances)

	//3. Get the list of classes from the top K closest samples
	topK := distances[:knn.K]

	//4. Count the classes and return the one that occurs most.
	results := make(map[string]int64)
	var highest int64 = 1
	var prediction string

	for _, cls := range topK {
		if _, ok := results[cls.Class]; ok {
			results[cls.Class]++
		} else {
			results[cls.Class] = 1
		}
	}
	for key, value := range results {
		if value > highest {
			highest = value
			prediction = key
		}
	}
	return prediction
}
