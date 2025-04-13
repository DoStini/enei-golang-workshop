package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckUrls(t *testing.T) {
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

	got := CheckUrls(urls)
	assert.EqualValues(t, len(expected), len(got))

	for key, value := range expected {
		if got[key] != value {
			t.Errorf("key %s :got %+v, expected %+v", key, got[key], value)
		}
	}
}

func TestCheckUrlsMalformed(t *testing.T) {
	urls := []string{
		"malfored url ",
		"https://",
		"what://?asd",
	}
	expected := map[string]bool{
		"malfored url ": false,
		"https://":      false,
		"what://?asd":   false,
	}

	got := CheckUrls(urls)
	assert.EqualValues(t, len(expected), len(got))

	for key, value := range expected {
		if got[key] != value {
			t.Errorf("key %s :got %+v, expected %+v", key, got[key], value)
		}
	}
}
