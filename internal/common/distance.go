package common

import (
	"fmt"
	"math"
)

func Euclidean(vectorA []float64, vectorB []float64) (float64, error) {
	if len(vectorA) != len(vectorB) {
		return 0, fmt.Errorf("input vector must have the same length")
	}

	sum := 0.0
	vectorLength := len(vectorA)

	for i := 0; i < vectorLength; i++ {
		diff := vectorA[i] - vectorB[i]
		sum += diff * diff
	}

	return math.Sqrt(sum), nil
}

func EuclideanL2(vectorA []float64, vectorB []float64) (float64, error) {
	result, err := Euclidean(L2Normalization(vectorA), L2Normalization(vectorB))
	if err != nil {
		return 0, err
	}
	return result, nil
}

func Cosine(vectorA []float64, vectorB []float64) (float64, error) {
	if len(vectorA) != len(vectorB) {
		return 0, fmt.Errorf("input vector must have the same length")
	}

	dotProduct := 0.0
	magnitudeA := 0.0
	magnitudeB := 0.0

	for i := 0; i < len(vectorA); i++ {
		dotProduct += vectorA[i] * vectorB[i]
		magnitudeA += vectorA[i] * vectorA[i]
		magnitudeB += vectorB[i] * vectorB[i]
	}

	magnitudeA = math.Sqrt(magnitudeA)
	magnitudeB = math.Sqrt(magnitudeB)

	return (dotProduct / (magnitudeA * magnitudeB)) * -1, nil
}

func L2Normalization(vector []float64) []float64 {
	sum := 0.0
	vectorLength := len(vector)

	for i := 0; i < vectorLength; i++ {
		sum += vector[i] * vector[i]
	}

	sqrt := math.Sqrt(sum)

	for i, val := range vector {
		vector[i] = val / sqrt
	}

	return vector
}
