package usecase

import (
	"math"
	"vektor/internal/common"
	"vektor/internal/constants"
	"vektor/internal/entity"
)

func (u *Usecase) Search(vector []float64, method string) (entity.Data, error) {
	methodFunc := u.methodSwitcher(method)

	files, err := u.repository.Files()
	if err != nil {
		return entity.Data{}, err
	}

	numProcesses := len(files)
	resultChannel := make(chan entity.Data, numProcesses)

	for _, file := range files {
		go u.calculateDistance(vector, file.Name(), resultChannel, methodFunc)
	}

	minDistance := math.Inf(1)
	key := ""
	for i := 0; i < numProcesses; i++ {
		data := <-resultChannel
		if data.MinDistance < minDistance {
			minDistance = data.MinDistance
			key = data.Key
		}
	}

	return entity.Data{
		Key:         key,
		MinDistance: minDistance,
	}, nil
}

func (u *Usecase) calculateDistance(vector []float64, key string, resultChannel chan<- entity.Data, methodFunc func(vectorA []float64, vectorB []float64) (float64, error)) {
	data, _ := u.repository.Read(key)

	minDistance := math.Inf(1)
	for _, targetVector := range data.Vectors {
		distance, _ := methodFunc(vector, targetVector)
		if distance < minDistance {
			minDistance = distance
		}
	}

	resultChannel <- entity.Data{
		Key:         key,
		MinDistance: minDistance,
	}
}

func (u *Usecase) methodSwitcher(method string) func(vectorA []float64, vectorB []float64) (float64, error) {
	var methodFunc func(vectorA []float64, vectorB []float64) (float64, error)
	switch method {
	case constants.DISTANCE_EUCLIDEAN:
		methodFunc = common.Euclidean
	case constants.DISTANCE_EUCLIDEAN_L2:
		methodFunc = common.EuclideanL2
	case constants.DISTANCE_COSINE:
		methodFunc = common.Cosine
	}
	return methodFunc
}
