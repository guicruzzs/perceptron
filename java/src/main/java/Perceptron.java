public class Perceptron {
    double errorThreshold;
    double weights[];
    int trainData[][];
    int target[];

    public void toTrain(int _trainData[][], int _target[]) {
        this.toTrain(_trainData, _target, 0.1);
    }

    public void toTrain(int _trainData[][], int _target[], double _errorThreshold) {
        errorThreshold = _errorThreshold;
        trainData = _trainData;
        target = _target;

        weights = new double[_trainData[0].length + 1];
        for(int i=0; i < weights.length; i++)
            weights[i] = 1;

        boolean isItOver = true;
        do{
            isItOver = true;
            for (int line=0; line< target.length ; line++) {
                int error = calculateError(line);
                if (error != 0) {
                    calculateWeights(error, line);
                    isItOver = false;
                }
            }
        }while(isItOver == false);
    }

    public void calculateWeights(int error, int line) {
        weights[0] = weights[0] + errorThreshold * error;
        for (int i=1; i < weights.length ; i++) {
            weights[i] = weights[i] + errorThreshold * error * trainData[line][i-1];
        }
    }

    public int calculateError(int line) {
        return target[line] - calculateOutput(trainData[line]);
    }

    public int calculateOutput(int data[]) {
        double total = weights[0];
        for (int i=1; i < weights.length ; i++) {
            total += weights[i] * data[i-1];
        }
        return (total >=0) ? 1 : 0;
    }
}