package perceptron

var Weights []float64
var TrainData [][]int
var Target []int
var ErrorThreshold float64

func ToTrain(trainData [][]int, target []int) {
	ToTrainAccurracy(trainData, target, 0.1)
}

func ToTrainAccurracy(trainData [][]int, target []int, errorThreshold float64) {
	TrainData = trainData
	Target = target
	ErrorThreshold = errorThreshold
	Weights = []float64{1}
	for i := 0; i < len(TrainData[0]); i++ {
		Weights = append(Weights, 1)
	}

	for isItOver := false; !isItOver; {
		isItOver = true
		for line, _ := range Target {
			error := CalculateError(line)
			if error != 0 {
				CalculateWeights(error, line)
				isItOver = false
			}
		}
	}
}

func CalculateOutput(data []int) int {
	var total float64 = 0
	for index, value := range append([]int{1},data...) {
		total += float64(value) * Weights[index]
	}
	if total >= 0 {return 1} else { return 0}
}

func CalculateError(line int) int {
	return Target[line] - CalculateOutput(TrainData[line])
}

func CalculateWeights(error int, line int) {
	for index, value := range append([]int{1},TrainData[line]...) {
		Weights[index] = Weights[index] + ErrorThreshold * float64(error) * float64(value)
	}
}