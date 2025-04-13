package main

import (
	"testing"
)

func TestIterateOverUrlSync(t *testing.T) {
	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.example.com",
		"https://www.twitter.com",
		"https://blog.gypsydave5.com",
	}
	expected := map[string]bool{
		"https://www.google.com":      true,
		"https://www.facebook.com":    true,
		"https://www.example.com":     true,
		"https://blog.gypsydave5.com": true,
		"https://www.twitter.com":     false,
	}

	got := iterateOverUrlSync(urls)

	for key, value := range expected {
		if got[key] != value {
			t.Errorf("got %+v, expected %+v", got[key], value)
		}
	}
}

func TestIterateOverUrlAsync(t *testing.T) {
	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.example.com",
		"https://www.twitter.com",
		"https://blog.gypsydave5.com",
	}
	expected := map[string]bool{
		"https://www.google.com":      true,
		"https://www.facebook.com":    true,
		"https://www.example.com":     true,
		"https://blog.gypsydave5.com": true,
		"https://www.twitter.com":     false,
	}

	got := iterateOverUrlAsync(urls)

	for key, value := range expected {
		if got[key] != value {
			t.Errorf("got %+v, expected %+v", got[key], value)
		}
	}
}
