package bst

import (
	"fmt"
	"testing"
)

func TestDelete(t *testing.T) {
	root := NewNode(Float64(1))
	data := New(root)

	slice := []float64{3, 9, 13, 17, 20, 25, 39, 16, 15, 2, 2.5}
	for _, num := range slice {
		data.Insert(NewNode(Float64(num)))
	}

	if fmt.Sprintf("%s", data) != "[1 [[2 [2.5]] 3 [9 [13 [[[15] 16] 17 [20 [25 [39]]]]]]]]" {
		t.Fatalf("Not expected output: %s\n", data)
	}

	if fmt.Sprintf("%s", data.Search(Float64(20))) != "[20 [25 [39]]]" {
		t.Fatalf("Not expected output: %s\n", data)
	}

	if data.Max().Key != Float64(39.0) {
		t.Fatalf("Expected 39.0 but %f", data.Max().Key)
	}

	if data.Min().Key != Float64(1.0) {
		t.Fatalf("Expected 1.0 but %f", data.Min().Key)
	}

	if data.SearchParent(Float64(16)).Key != Float64(17.0) {
		t.Fatalf("Expected 17.0 but %f", data.SearchParent(Float64(16)).Key)
	}

	if data.Search(Float64(39)).Right != nil {
		t.Fatal("39's Right child must be nil")
	}
	if data.Search(Float64(25)).Left != nil {
		t.Fatal("25's Left child must be nil")
	}

	deletes := []float64{13, 17, 3, 15, 1, 2.5}
	for _, num := range deletes {
		data.Delete(data.Search(Float64(num)))

		t.Logf("Deleted: %f\n", num)

		if data.Search(Float64(num)) != nil {
			t.Fatal(num, "must be nil")
		}
		if data.Search(Float64(1000.12)) != nil {
			t.Fatal(1000.12, "must be nil")
		}

		switch num {
		case 13:
			if data.Search(Float64(9)).Right.Key != Float64(17) {
				t.Fatal("9's Right child must be 17")
			}
			if data.SearchParent(Float64(17)).Key != Float64(9) {
				t.Fatal("17's Parent must be 9")
			}
			if data.Search(Float64(17)).Left.Key != Float64(16) {
				t.Fatal("17's Left child must be 16")
			}

		case 17:
			if data.Search(Float64(9)).Right.Key != Float64(16) {
				t.Fatal("9's Right child must be 16")
			}
			if data.SearchParent(Float64(16)).Key != Float64(9) {
				t.Fatal("16's Parent must be 9")
			}
			if data.Search(Float64(16)).Right.Key != Float64(20) {
				t.Fatal("16's Right child must be 20")
			}
			if data.Search(Float64(16)).Left.Key != Float64(15) {
				t.Fatal("16's Left must be 15")
			}
			if data.SearchParent(Float64(20)).Key != Float64(16) {
				t.Fatal("20's Parent must be 16")
			}
			if data.SearchParent(Float64(15)).Key != Float64(16) {
				t.Fatal("15's Parent must be 16")
			}

		case 3:
			if data.Search(Float64(9)).Right.Key != Float64(16) {
				t.Fatal("9's Right child must be 16")
			}
			if data.SearchParent(Float64(16)).Key != Float64(9) {
				t.Fatal("16's Parent must be 9")
			}
			if data.Search(Float64(16)).Right.Key != Float64(20) {
				t.Fatal("16's Right child must be 20")
			}
			if data.Search(Float64(16)).Left.Key != Float64(15) {
				t.Fatal("16's Left must be 15")
			}
			if data.SearchParent(Float64(20)).Key != Float64(16) {
				t.Fatal("20's Parent must be 16")
			}
			if data.SearchParent(Float64(15)).Key != Float64(16) {
				t.Fatal("15's Parent must be 16")
			}
			if data.SearchParent(Float64(2)).Key != Float64(2.5) {
				t.Fatal("2's Parent must be 2.5")
			}
			if data.Search(Float64(2.5)).Left.Key != Float64(2) {
				t.Fatal("2.5's Left must be 2")
			}
			if data.Search(Float64(2.5)).Right.Key != Float64(9) {
				t.Fatal("2.5's Right must be 9")
			}
			if data.Search(Float64(1)).Right.Key != Float64(2.5) {
				t.Fatal("1's Right must be 2.5")
			}
			if data.Search(Float64(1)).Left != nil {
				t.Fatal("1's Left must be nil")
			}

		case 15:
			if data.Search(Float64(9)).Right.Key != Float64(16) {
				t.Fatal("9's Right child must be 16")
			}
			if data.SearchParent(Float64(16)).Key != Float64(9) {
				t.Fatal("16's Parent must be 9")
			}
			if data.Search(Float64(16)).Right.Key != Float64(20) {
				t.Fatal("16's Right child must be 20")
			}
			if data.Search(Float64(16)).Left != nil {
				t.Fatal("16's Left must be nil")
			}
			if data.SearchParent(Float64(20)).Key != Float64(16) {
				t.Fatal("20's Parent must be 16")
			}
			if data.SearchParent(Float64(2)).Key != Float64(2.5) {
				t.Fatal("2's Parent must be 2.5")
			}
			if data.Search(Float64(2.5)).Left.Key != Float64(2) {
				t.Fatal("2.5's Left must be 2")
			}
			if data.Search(Float64(2.5)).Right.Key != Float64(9) {
				t.Fatal("2.5's Right must be 9")
			}
			if data.Search(Float64(1)).Right.Key != Float64(2.5) {
				t.Fatal("1's Right must be 2.5")
			}
			if data.Search(Float64(1)).Left != nil {
				t.Fatal("1's Left must be nil")
			}
			if data.Search(Float64(16)).Left != nil {
				t.Fatal("16's Left must be nil")
			}

		case 1:
			if data.Search(Float64(9)).Right.Key != Float64(16) {
				t.Fatal("9's Right child must be 16")
			}
			if data.SearchParent(Float64(16)).Key != Float64(9) {
				t.Fatal("16's Parent must be 9")
			}
			if data.Search(Float64(16)).Right.Key != Float64(20) {
				t.Fatal("16's Right child must be 20")
			}
			if data.Search(Float64(16)).Left != nil {
				t.Fatal("16's Left must be nil")
			}
			if data.SearchParent(Float64(20)).Key != Float64(16) {
				t.Fatal("20's Parent must be 16")
			}
			if data.Search(Float64(2.5)).Right.Key != Float64(9) {
				t.Fatal("2.5's Right must be 9")
			}
			if data.Search(Float64(16)).Left != nil {
				t.Fatal("16's Left must be nil")
			}
			if data.Search(Float64(2.5)).Left.Key != Float64(2) {
				t.Fatalf("2.5's Left must be 2 but %s\n", data.Search(Float64(2.5)).Left)
			}
			if data.SearchParent(Float64(2)).Key != Float64(2.5) {
				t.Fatal("2's Parent must be 2.5")
			}
			if data.Search(Float64(2)).Right != nil {
				t.Fatal("2's Right must be nil")
			}
			if data.SearchParent(Float64(2.5)) != nil {
				t.Fatal("2.5's Parent must be nil")
			}
			if data.Root.Key != Float64(2.5) {
				t.Fatal("Root must be 2.5")
			}

		case 2.5:
			if data.Search(Float64(2)).Right.Key != Float64(9) {
				t.Fatal("2's Right child must be 9")
			}
			if data.Search(Float64(9)).Right.Key != Float64(16) {
				t.Fatal("9's Right child must be 16")
			}
			if data.SearchParent(Float64(2)) != nil {
				t.Fatal("2's Parent must be nil")
			}
			if data.SearchParent(Float64(16)).Key != Float64(9) {
				t.Fatal("16's Parent must be 9")
			}
			if data.Search(Float64(16)).Right.Key != Float64(20) {
				t.Fatal("16's Right child must be 20")
			}
			if data.Search(Float64(16)).Left != nil {
				t.Fatal("16's Left must be nil")
			}
			if data.SearchParent(Float64(20)).Key != Float64(16) {
				t.Fatal("20's Parent must be 16")
			}
			if data.Search(Float64(2.5)) != nil {
				t.Fatal("2.5's Right must be nil")
			}
			if data.Search(Float64(16)).Left != nil {
				t.Fatal("16's Left must be nil")
			}
			if data.Search(Float64(2)).Right.Key != Float64(9) {
				t.Fatal("2's Right must be 9")
			}
			if data.Root.Key != Float64(2) {
				t.Fatal("Root must be 2")
			}

		default:
			t.Fatal(num, "shouldn't be there...")
		}
	}
}
