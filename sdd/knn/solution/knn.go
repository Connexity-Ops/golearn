package knn

import "github.com/hongshibao/go-kdtree"

type Sample struct {
	kdtree.Point
	Attributes []float64
	Class      string
}

func (s *Sample) Dim() int {
	return len(s.Attributes)
}

func (s *Sample) GetValue(dim int) float64 {
	return s.Attributes[dim]
}

func (s *Sample) GetClass() string {
	return s.Class
}

func (s *Sample) Distance(other kdtree.Point) float64 {
	var ret float64
	for i := 0; i < s.Dim(); i++ {
		tmp := s.GetValue(i) - other.GetValue(i)
		ret += tmp * tmp
	}
	return ret
}

func (s *Sample) PlaneDistance(val float64, dim int) float64 {
	tmp := s.GetValue(dim) - val
	return tmp * tmp
}

func NewSample(attributes []float64, class string) *Sample {
	return &Sample{Attributes: attributes, Class: class}
}

type KNNClassifier struct {
	K            int
	TrainingData []kdtree.Point
}

func NewKNNClassifier(k int) *KNNClassifier {
	return &KNNClassifier{K: k}
}

func (knn *KNNClassifier) Fit(samples []*Sample) {
	knn.TrainingData = make([]kdtree.Point, len(samples))
	for i, s := range samples {
		knn.TrainingData[i] = kdtree.Point(s)
	}
}

func (knn *KNNClassifier) Predict(sample *Sample) string {
	tree := kdtree.NewKDTree(knn.TrainingData)
	neighbors := tree.KNN(sample, knn.K)
	votes := make(map[string]int)
	var max int
	var class, winner string
	for _, p := range neighbors {
		class = p.(*Sample).Class
		votes[class]++
		if votes[class] > max {
			winner = class
			max = votes[class]
		}
	}
	return winner
}
