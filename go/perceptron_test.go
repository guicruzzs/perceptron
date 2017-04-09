package perceptron

import "testing"

func trainNeuron() {
	trainData := [][]int{{1,1,1,0,1,0,0,1,0},
                         {1,0,1,1,1,1,1,0,1}}
    target := []int{0, 1}
	ToTrain(trainData, target)
}

func TestToTrain(t *testing.T) {
	trainNeuron()

	if Weights[0] != -0.09999999999999987 {
		t.Error("Wrong neuron weight")
	}
	if Weights[1] != -0.09999999999999987 {
		t.Error("Wrong neuron weight")
	}
	if Weights[2] != -0.09999999999999987 {
		t.Error("Wrong neuron weight")
	}
	if Weights[3] != -0.09999999999999987 {
		t.Error("Wrong neuron weight")
	}
	if Weights[4] !=  1.0                 {
		t.Error("Wrong neuron weight")
	}
	if Weights[5] != -0.09999999999999987 {
		t.Error("Wrong neuron weight")
	}
	if Weights[6] !=  1.0                 {
		t.Error("Wrong neuron weight")
	}
	if Weights[7] !=  1.0                 {
		t.Error("Wrong neuron weight")
	}
	if Weights[8] != -0.09999999999999987 {
		t.Error("Wrong neuron weight")
	}
}

func TestCalculateOutput(t *testing.T) {
	sampleClass1 := []int{1,0,1,0,1,0,0,1,0}

	if CalculateOutput(sampleClass1) != 0 {
		t.Error("Classification error")
	}

	sampleClass2 := []int{1,0,1,1,1,1,0,0,1}
	if CalculateOutput(sampleClass2) != 1 {
		t.Error("Classification error")
	}
}

func TestCalculateError(t *testing.T) {
	if CalculateError(0) != 0 {
		t.Error("Fail on error calculation")
	}

	if CalculateError(1) != 0 {
		t.Error("Fail on error calculation")
	}

    anotherTrainData := [][]int{{1,1,1,0,1,0,0,0,1},
                                {1,1,1,0,0,0,0,1,0}}
    TrainData = anotherTrainData

    if CalculateError(0) != -1 {
		t.Error("Fail on error calculation")
	}

	if CalculateError(1) !=  1 {
		t.Error("Fail on error calculation")
	}
}
