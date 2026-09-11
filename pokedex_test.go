package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/marcel-alter/pokedexcli/internal/pokecache"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hello  world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  hello  world  AGain",
			expected: []string{"hello", "world", "again"},
		},
		{
			input:    "THiS WoRks JuSt FinE !!!",
			expected: []string{"this", "works", "just", "fine", "!!!"},
		},
	}
	for count, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice
		// if they don't match, use t.Errorf and continue to the next case
		if len(actual) != len(c.expected) {
			t.Errorf("Error: Not the expected lenght in case #%d", count+1)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("%s is not %s !!! Test #%d failed", word, expectedWord, count+1)
			}
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
		}
	}

}

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := pokecache.NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := pokecache.NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
