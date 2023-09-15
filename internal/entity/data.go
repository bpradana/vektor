package entity

type Data struct {
	Key         string      `json:"key,omitempty"`
	Vectors     [][]float64 `json:"vectors,omitempty"`
	Distances   []float64
	MinDistance float64
}
