package haikunator

import (
	"slices"
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
	if !slices.Contains(adjectives[:], parts[0]) {
		t.Fatalf("unknown adjective: %q", parts[0])
	}
	if !slices.Contains(nouns[:], parts[1]) {
		t.Fatalf("unknown noun: %q", parts[1])
	}
}

func TestHaikunateSeededDeterministic(t *testing.T) {
	h1 := NewWithSeed(42)
	h2 := NewWithSeed(42)

	if got1, got2 := h1.Haikunate(), h2.Haikunate(); got1 != got2 {
		t.Fatalf("expected seeded names to match, got %q and %q", got1, got2)
	}
}

func TestUniqueness(t *testing.T) {
	h := New()
	seen := make(map[string]bool)

	for range 10 {
		name := h.Haikunate()
		if seen[name] {
			t.Fatalf("duplicate name: %s", name)
		}
		seen[name] = true
	}
}

func TestSize(t *testing.T) {
	h := New()
	expected := len(adjectives) * len(nouns)
	if h.Size() != expected {
		t.Fatalf("expected %d combinations, got %d", expected, h.Size())
	}
}
