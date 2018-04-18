// Package knn implements a K nearest neighbors classifier.

package knn

type Sample struct {
	Attributes []float64
	Class      string
}

func (s *Sample) Distance(other *Sample) float64 {
	/**
	Distance should calculate the euclidean distance from one sample to another.
	Use this formula: d = (x1 - x2)^2 + (y1 - y2)^2 + (z1 - z2)^2
	**/
}

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

func (knn *KNNClassifier) Fit(samples []*Sample) {
	// This method should store the slice of samples in the struct
}

func (knn *KNNClassifier) Predict(sample *Sample) string {
	/**
	This is where the work happens

	1. Get the distance from sample to every other sample in knn.TrainingData
	2. Sort by closest samples based on calculated distances
	3. Get the list of classes from the top K closest samples
	4. Count the classes and return the one that occurs most

	**/
}
