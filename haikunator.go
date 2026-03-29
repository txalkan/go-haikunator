package haikunator

import (
	"fmt"
	"math/rand/v2"
)

var (
	adjectives = [...]string{
		"aged", "ancient", "autumn", "billowing", "bitter", "black", "blue", "bold",
		"broad", "broken", "calm", "cold", "cool", "crimson", "curly", "damp",
		"dark", "dawn", "delicate", "divine", "dry", "empty", "falling", "fancy",
		"flat", "floral", "fragrant", "frosty", "gentle", "green", "hidden", "holy",
		"icy", "jolly", "late", "lingering", "little", "lively", "long", "lucky",
		"misty", "morning", "muddy", "mute", "nameless", "noisy", "odd", "old",
		"orange", "patient", "plain", "polished", "proud", "purple", "quiet", "rapid",
		"raspy", "red", "restless", "rough", "round", "royal", "shiny", "shrill",
		"shy", "silent", "small", "snowy", "soft", "solitary", "sparkling", "spring",
		"square", "steep", "still", "summer", "super", "sweet", "throbbing", "tight",
		"tiny", "twilight", "wandering", "weathered", "white", "wild", "winter", "wispy",
		"withered", "yellow", "young",
	}
	nouns = [...]string{
		"art", "band", "bar", "base", "bird", "block", "boat", "bonus",
		"bread", "breeze", "brook", "bush", "butterfly", "cake", "cell", "cherry",
		"cloud", "credit", "darkness", "dawn", "dew", "disk", "dream", "dust",
		"feather", "field", "fire", "firefly", "flower", "fog", "forest", "frog",
		"frost", "glade", "glitter", "grass", "hall", "hat", "haze", "heart",
		"hill", "king", "lab", "lake", "leaf", "limit", "math", "meadow",
		"mode", "moon", "morning", "mountain", "mouse", "mud", "night", "paper",
		"pine", "poetry", "pond", "queen", "rain", "recipe", "resonance", "rice",
		"river", "salad", "scene", "sea", "shadow", "shape", "silence", "sky",
		"smoke", "snow", "snowflake", "sound", "star", "sun", "sunset",
		"surf", "term", "thunder", "tooth", "tree", "truth", "union", "unit",
		"violet", "voice", "water", "waterfall", "wave", "wildflower", "wind", "wood",
	}
)

// Haikunator generates Heroku-like memorable random names.
type Haikunator struct {
	r *rand.Rand
}

// New creates a Haikunator that uses the global random source (auto-seeded).
func New() *Haikunator {
	return &Haikunator{}
}

// NewWithSeed creates a Haikunator with a deterministic seed for reproducible output.
func NewWithSeed(seed uint64) *Haikunator {
	return &Haikunator{r: rand.New(rand.NewPCG(seed, seed))} //nolint:gosec // crypto rand unnecessary for name generation
}

// Haikunate returns a random name in the form "adjective-noun".
func (h *Haikunator) Haikunate() string {
	adj, noun := h.pick()
	return fmt.Sprintf("%s-%s", adj, noun)
}

// Size returns the total number of unique name combinations.
func (h *Haikunator) Size() int {
	return len(adjectives) * len(nouns)
}

func (h *Haikunator) pick() (string, string) {
	if h.r != nil {
		return adjectives[h.r.IntN(len(adjectives))], nouns[h.r.IntN(len(nouns))]
	}
	return adjectives[rand.IntN(len(adjectives))], nouns[rand.IntN(len(nouns))] //nolint:gosec // crypto rand unnecessary for name generation
}
