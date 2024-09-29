package funk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContains(t *testing.T) {

	// Test data
	data := []struct {
		name     string
		slice    []string
		contains string
		expected bool
	}{
		{"empty slice", []string{}, "a", false},
		{"single element", []string{"a"}, "a", true},
		{"single element", []string{"a"}, "b", false},
		{"multiple elements", []string{"a", "b", "c"}, "a", true},
		{"multiple elements", []string{"a", "b", "c"}, "b", true},
		{"multiple elements", []string{"a", "b", "c"}, "c", true},
		{"multiple elements", []string{"a", "b", "c"}, "d", false},
	}

	// Test loop
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			res := Contains(d.slice, d.contains)
			assert.Equal(t, d.expected, res)
		})
	}

}

func TestMax(t *testing.T) {

	// test data
	data := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"a > b", 1, 0, 1},
		{"a < b", 0, 1, 1},
		{"a = b", 1, 1, 1},
	}

	// test all
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			value := Max(d.a, d.b)
			assert.Equal(t, d.expected, value)
		})
	}
}

func TestMin(t *testing.T) {

	// test data
	data := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"a > b", 1, 0, 0},
		{"a < b", 0, 1, 0},
		{"a = b", 1, 1, 1},
	}

	// test all
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			value := Min(d.a, d.b)
			assert.Equal(t, d.expected, value)
		})
	}
}

func TestFuzzySearch(t *testing.T) {

	// test data
	data := []struct {
		name     string
		slice    []string
		search   string
		expected []string
	}{
		{"empty slice", []string{}, "a", []string{}},
		{"single element", []string{"a"}, "a", []string{"a"}},
		{"single element", []string{"a"}, "b", []string{}},
		{"multiple elements", []string{"a", "b", "c"}, "a", []string{"a"}},
		{"multiple elements", []string{"a", "b", "c"}, "b", []string{"b"}},
		{"multiple elements", []string{"a", "b", "c"}, "c", []string{"c"}},
		{"multiple elements", []string{"abc", "bcd", "cde"}, "d", []string{}},
		{"multiple elements", []string{"abc", "bcd", "cde"}, "b", []string{"bcd"}},
		{"multiple elements", []string{"abc", "bcd", "cde"}, "c", []string{"cde"}},
		{"multiple elements", []string{"abc", "bcd", "cde"}, "a", []string{"abc"}},
		{"multiple elements", []string{"abc", "acd", "cde"}, "a", []string{"abc", "acd"}},
		{"multiple elements", []string{"abc", "acd", "cde"}, "c", []string{"cde"}},
		{"multiple elements", []string{"abc", "acd", "cde"}, "d", []string{}},
	}

	// test all
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			res := FuzzySearch(d.search, d.slice...)
			assert.Equal(t, d.expected, res)
		})
	}
}
