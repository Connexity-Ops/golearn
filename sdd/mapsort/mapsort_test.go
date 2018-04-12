package mapsort

import (
	"reflect"
	"testing"
)
func TestMapSort(t *testing.T) {
	tests := []struct {
		name string
		args map[string]int
		want PairList
	}{
		{
			name: "Presidential Numbers",
			args: map[string]int{
				"George Washington"		: 1,
				"Franklin Delano Roosevelt"	: 32,
				"John Fitzgerald Kennedy"	: 35,
				"Thomas Jefferson"		: 3,
				"Abraham Lincoln"		: 16,
				"Barack Hussein Obama"		: 44,
				},
			want: PairList{
				{"George Washington"		, 1}, 
				{"Thomas Jefferson"		, 3},
				{"Abraham Lincoln"		, 16},
				{"Franklin Delano Roosevelt"	, 32},
				{"John Fitzgerald Kennedy"	, 35},
				{"Barack Hussein Obama"		, 44},
			},
		},
		{
			name: "Roman Numerals",
			args: map[string]int{
				"M"	: 1000,
				"CM"	: 900,
				"D"	: 500,
				"CD"	: 400,
				"C"	: 100,
				"XC"	: 90,
				"L"	: 50,
				"XL"	: 40,
				"X"	: 10,
				"IX"	: 9,
				"V"	: 5,
				"IV"	: 4,
				"I"	: 1,
				},
			want: PairList{
				{"I" ,	1},
				{"IV",	4},
				{"V" ,	5},
				{"IX",	9},
				{"X" ,	10},
				{"XL",	40},
				{"L" ,	50},
				{"XC",	90},
				{"C" ,	100},
				{"CD",	400},
				{"D" ,	500},
				{"CM",	900},	
				{"M" ,	1000},		
			},
		},
		{
			name: "Mr Robot S1 Episode List",
			args: map[string]int{
				"eps1.0_hellofriend.mov"	: 1,
				"eps1.1_ones-and-zer0es.mpeg"	: 2,
				"eps1.2_d3bug.mkv"		: 3,
				"eps1.3_da3m0ns.mp4"		: 4,
				"eps1.4_3xpl0its.wmv"		: 5,
				"eps1.5_br4ve-trave1er.asf"	: 6,
				"eps1.6_v1ew-s0urce.flv"	: 7,
				"eps1.7_wh1ter0se.m4v"		: 8,
				"eps1.8_m1rr0r1ng.qt"		: 9,
				"eps1.9_zer0-day.avi"		: 10,
				},
			want: PairList{
				{"eps1.0_hellofriend.mov"	, 1},
				{"eps1.1_ones-and-zer0es.mpeg"	, 2},
				{"eps1.2_d3bug.mkv"		, 3},
				{"eps1.3_da3m0ns.mp4"		, 4},
				{"eps1.4_3xpl0its.wmv"		, 5},
				{"eps1.5_br4ve-trave1er.asf"	, 6},
				{"eps1.6_v1ew-s0urce.flv"	, 7},
				{"eps1.7_wh1ter0se.m4v"		, 8},
				{"eps1.8_m1rr0r1ng.qt"		, 9},
				{"eps1.9_zer0-day.avi"		, 10},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapSort(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapSort(m map[string]int) = \n%v\nWants \n%v", got, tt.want)
			}
		})
	}
}
