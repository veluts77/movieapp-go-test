package model

// Metadata defines the movie metadata.
type Metadata struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Disrector   string `json:"director"`
}
