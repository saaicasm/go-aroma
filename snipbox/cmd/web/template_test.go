package main

import (
	"github/saaicasm/snipbox/internal/assert"
	"testing"
	"time"
)

func TestReadableDate(t *testing.T) {
	t.Run("UTC time", func(t *testing.T) {
		tm := time.Date(2024, 3, 17, 10, 15, 0, 0, time.UTC)
		want := "17 Mar 2024 at 10:15"
		got := readableDate(tm)
		assert.Equal(t, want, got)
	})

	t.Run("Empty time", func(t *testing.T) {
		tm := time.Time{}
		want := ""
		got := readableDate(tm)
		assert.Equal(t, want, got)
	})

	t.Run("CET time", func(t *testing.T) {
		tm := time.Date(2024, 3, 17, 10, 15, 0, 0, time.FixedZone("CET", 1*60*60))
		want := "17 Mar 2024 at 09:15"
		got := readableDate(tm)
		assert.Equal(t, want, got)
	})
}
