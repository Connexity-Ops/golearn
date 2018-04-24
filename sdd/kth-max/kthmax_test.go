package kthstats

import (
	"errors"
	"reflect"
	"testing"
)

func TestKthMaxInt(t *testing.T) {
	type args struct {
		k       int
		numbers []int
	}
	type wants struct {
		kmax int
		err  error
	}

	normalNumbers := []int{515, 818, -409, 327, 542, -607, 387, 811, 529, 768, 235, 625, 146, 749, 774, 800, 329, 125, 616, 596, 811}

	tests := []struct {
		name  string
		args  args
		wants wants
	}{
		{
			name:  "Normal maximum",
			args:  args{k: 1, numbers: normalNumbers},
			wants: wants{kmax: 818},
		},
		{
			name:  "k=2 normal maximum",
			args:  args{k: 2, numbers: normalNumbers},
			wants: wants{kmax: 811},
		},
		{
			name:  "k=7 normal maximum",
			args:  args{k: 7, numbers: normalNumbers},
			wants: wants{kmax: 625},
		},
		{
			name:  "k=20 normal maximum",
			args:  args{k: 20, numbers: normalNumbers},
			wants: wants{kmax: -607},
		},
		{
			name:  "k=20 short 5 element maximum",
			args:  args{k: 20, numbers: normalNumbers[:5]},
			wants: wants{err: errors.New("Invalid k value 20; list does not have enough unique elements")},
		},
		{
			name:  "k=1 exclusively same element maximum",
			args:  args{k: 1, numbers: []int{1, 1, 1, 1, 1, 1, 1, 1}},
			wants: wants{kmax: 1},
		},
		{
			name:  "k=2 exclusively same element maximum",
			args:  args{k: 2, numbers: []int{1, 1, 1, 1, 1, 1, 1, 1}},
			wants: wants{err: errors.New("Invalid k value 2; list does not have enough unique elements")},
		},
		{
			name:  "k=2 repeated element is k maximum",
			args:  args{k: 2, numbers: []int{1, 1, 1, 1, 1, 1, 1, 1, 2}},
			wants: wants{kmax: 1},
		},
		{
			name:  "k=2 repeated element is k maximum surrounded",
			args:  args{k: 2, numbers: []int{1, 1, 1, 1, 1, 1, 1, 1, 2, 0}},
			wants: wants{kmax: 1},
		},
		{
			name:  "k=1 one element maximum",
			args:  args{k: 1, numbers: normalNumbers[:1]},
			wants: wants{kmax: 515},
		},
		{
			name:  "k=3 maximum at beginning of the list",
			args:  args{k: 3, numbers: []int{1, 2, 3}},
			wants: wants{kmax: 1},
		},
		{
			name:  "k=3 maximum at end of the list",
			args:  args{k: 3, numbers: []int{2, 3, 1}},
			wants: wants{kmax: 1},
		},
		{
			name:  "k=0 one element maximum",
			args:  args{k: 0, numbers: normalNumbers},
			wants: wants{err: errors.New("Invalid k value 0")},
		},
		{
			name:  "k=-1 one element maximum",
			args:  args{k: -1, numbers: normalNumbers},
			wants: wants{err: errors.New("Invalid k value -1")},
		},
		{
			name:  "k=4 empty list",
			args:  args{k: 4, numbers: []int{}},
			wants: wants{err: errors.New("The list does not have enough unique elements for any value k")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kmax, err := KMaxInt(tt.args.k, tt.args.numbers)
			if !reflect.DeepEqual(tt.wants.err, err) {
				t.Errorf("%s expected err (%s) but got (%v)", tt.name, tt.wants.err, err)
			}
			if err == nil && kmax != tt.wants.kmax {
				t.Errorf("KMaxInt(%d, %v) = (%v, %v), wants (%v, %v)", tt.args.k, tt.args.numbers, kmax, err, tt.wants.kmax, tt.wants.err)
			}
		})
	}
}
