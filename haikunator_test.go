package haikunator

import (
	"strings"
	"testing"
)

func TestHaikunateFormat(t *testing.T) {
	h := New()
	name := h.Haikunate()

	parts := strings.SplitN(name, "-", 2)
	if len(parts) != 2 {
		t.Fatalf("expected adjective-noun format, got %q", name)
	}
	if parts[0] == "" || parts[1] == "" {
		t.Fatalf("empty component in %q", name)
	}
}

func TestUniqueness(t *testing.T) {
	h := NewWithSeed(12345)
	seen := make(map[string]bool)

	for range 20 {
		name := h.Haikunate()
		if seen[name] {
			t.Fatalf("duplicate name: %s", name)
		}
		seen[name] = true
	}
}

func TestDeterministicSeed(t *testing.T) {
	h1 := NewWithSeed(42)
	h2 := NewWithSeed(42)

	for range 50 {
		if h1.Haikunate() != h2.Haikunate() {
			t.Fatal("same seed should produce same sequence")
		}
	}
}

func TestSize(t *testing.T) {
	h := New()
	expected := len(adjectives) * len(nouns)
	if h.Size() != expected {
		t.Fatalf("expected %d combinations, got %d", expected, h.Size())
	}
}
