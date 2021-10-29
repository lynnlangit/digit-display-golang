package recognize

import (
	"fmt"

	"github.com/jeremybytes/digit-display-golang/shared"
)

type ManhattanClassifier struct {
	TrainingData []shared.Record
}

func (c ManhattanClassifier) String() string {
	return "Manhattan Classifier"
}

func (c *ManhattanClassifier) Train(traingingData []shared.Record) error {
	c.TrainingData = traingingData
	return nil
}

func Abs(input int) int {
	if input < 0 {
		return -input
	}
	return input
}

func (c ManhattanClassifier) Predict(input shared.Record) (prediction Prediction, err error) {
	best := shared.Record{Actual: -1, Image: nil}
	var bestTotal int = 100000000
	for _, train := range c.TrainingData {
		total := 0
		for i := range input.Image {
			total = total + Abs(input.Image[i]-train.Image[i])
		}
		if total < bestTotal {
			bestTotal = total
			best = train
		}
	}

	if best.Actual == -1 {
		return Prediction{}, fmt.Errorf("unable to get a valid prediction")
	}

	return Prediction{Actual: input, Predicted: best}, err
}
