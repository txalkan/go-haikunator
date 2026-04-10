package haikunator

import (
	cryptorand "crypto/rand"
	"fmt"
	"math/big"
	mathrand "math/rand"
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
	rng *mathrand.Rand
}

// New creates a Haikunator.
func New() *Haikunator {
	return &Haikunator{}
}

// NewWithSeed creates a deterministic Haikunator seeded with the given value.
func NewWithSeed(seed int64) *Haikunator {
	return &Haikunator{
		rng: mathrand.New(mathrand.NewSource(seed)),
	}
}

// Haikunate returns a random name in the form "adjective-noun".
func (h *Haikunator) Haikunate() string {
	return fmt.Sprintf("%s-%s", adjectives[h.randIntn(len(adjectives))], nouns[h.randIntn(len(nouns))])
}

// Size returns the total number of unique name combinations.
func (h *Haikunator) Size() int {
	return len(adjectives) * len(nouns)
}

func (h *Haikunator) randIntn(n int) int {
	if h != nil && h.rng != nil {
		return h.rng.Intn(n)
	}
	return randIntn(n)
}

func randIntn(n int) int {
	v, err := cryptorand.Int(cryptorand.Reader, big.NewInt(int64(n)))
	if err != nil {
		panic(fmt.Sprintf("crypto/rand failed: %v", err))
	}
	return int(v.Int64())
}
