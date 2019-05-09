package service

// Greeting does nothing actually
type Greeting struct {
	Derp string
}

func (g Greeting) String() string {
	return "Well... duh"
}
