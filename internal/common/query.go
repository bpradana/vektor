package common

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"vektor/internal/constants"
)

type Query struct {
	Action    string
	Method    string
	Threshold float64
	Key       string
	Vectors   [][]float64
}

func ParseQuery(query string) (*Query, error) {
	parts := strings.Split(query, " ")
	if len(parts) < 2 {
		return nil, fmt.Errorf(constants.ERROR_INVALID_QUERY)
	}

	q := &Query{
		Action: strings.ToUpper(parts[0]),
		Key:    parts[1],
	}

	switch q.Action {
	case constants.ACTION_READ, constants.ACTION_DELETE:
		if len(parts) != 2 {
			return nil, fmt.Errorf(constants.ERROR_INVALID_QUERY)
		}
	case constants.ACTION_CREATE, constants.ACTION_UPDATE:
		if len(parts) != 3 {
			return nil, fmt.Errorf(constants.ERROR_INVALID_QUERY)
		}

		vectors := [][]float64{}
		err := json.Unmarshal([]byte(parts[2]), &vectors)
		if err != nil {
			return nil, fmt.Errorf(constants.ERROR_DECODING_DATA, err)
		}

		q.Vectors = vectors
	case constants.ACTION_SEARCH:
		if len(parts) != 4 {
			return nil, fmt.Errorf(constants.ERROR_INVALID_QUERY)
		}

		vector := []float64{}
		err := json.Unmarshal([]byte(parts[3]), &vector)
		if err != nil {
			return nil, fmt.Errorf(constants.ERROR_DECODING_DATA, err)
		}

		q.Vectors = [][]float64{vector}
		q.Method = parts[1]
		q.Threshold, err = strconv.ParseFloat(parts[2], 64)
		if err != nil {
			return nil, fmt.Errorf(constants.ERROR_DECODING_DATA, err)
		}
		q.Key = ""
	}

	return q, nil
}
