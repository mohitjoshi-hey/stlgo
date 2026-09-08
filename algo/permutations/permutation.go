package permutations

import (
	"cmp"
	"slices"
)

func NextPermutation[T cmp.Ordered](slice []T) ([]T, bool) {
	n := len(slice)
	if n < 2 {
		return slice, false;
	}

	pivot := -1;
	for i := n - 2; i >= 0; i-- {
		if slice[i] < slice[i + 1] {
			pivot = i;
			break;
		}
	}
	if pivot == -1 {
		slices.Reverse(slice)
		return slice, false;
	}

	successor := -1;
	for i := n - 1; i > pivot; i-- {
		if slice[i] > slice[pivot] {
			successor = i;
			break;
		}
	}
	val := slice[pivot];
	slice[pivot] = slice[successor];
	slice[successor] = val;
	slices.Reverse(slice[pivot + 1:]);

	return slice, true;
}

func PrevPermutation[T cmp.Ordered](slice []T) ([]T, bool) {
	n := len(slice)
	if n < 2 {
		return slice, false;
	}

	pivot := -1;
	for i := n - 2; i >= 0; i-- {
		if slice[i] > slice[i + 1] {
			pivot = i;
			break;
		}
	}
	if pivot == -1 {
		slices.Reverse(slice)
		return slice, false;
	}

	successor := -1;
	for i := n - 1; i > pivot; i-- {
		if slice[i] < slice[pivot] {
			successor = i;
			break;
		}
	}
	val := slice[pivot];
	slice[pivot] = slice[successor];
	slice[successor] = val;
	slices.Reverse(slice[pivot + 1:]);

	return slice, true;
}

// HAve ton add KthPermutations function later.