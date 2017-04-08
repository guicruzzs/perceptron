import static org.junit.Assert.assertEquals;
import org.junit.Test;

public class PerceptronTest {
  public Perceptron getTrainedNeuron(){
    int trainData[][] = {{1,1,1,0,1,0,0,1,0},
                         {1,0,1,1,1,1,1,0,1}};
    int target[] = {0,1};
    Perceptron neuron = new Perceptron();
    neuron.toTrain(trainData, target);
    return neuron;
  }

  @Test
  public void testTraining() {
    Perceptron neuron = getTrainedNeuron();
    double weights[] = neuron.weights;
    double precision = 0.00000000000000001;

    assertEquals(-0.09999999999999987, weights[0], precision);
    assertEquals(-0.09999999999999987, weights[1], precision);
    assertEquals(-0.09999999999999987, weights[2], precision);
    assertEquals(-0.09999999999999987, weights[3], precision);
    assertEquals( 1.0,                 weights[4], precision);
    assertEquals(-0.09999999999999987, weights[5], precision);
    assertEquals( 1.0,                 weights[6], precision);
    assertEquals( 1.0,                 weights[7], precision);
    assertEquals(-0.09999999999999987, weights[8], precision);
  }

  @Test
  public void testCalculateOutput() {
    int sampleClass1[] = {1,1,1,0,1,0,0,1,0};
    int sampleClass2[] = {1,0,1,1,1,1,0,0,1};

    Perceptron neuron = getTrainedNeuron();

    assertEquals(0, neuron.calculateOutput(sampleClass1));
    assertEquals(1, neuron.calculateOutput(sampleClass2));
  }

  @Test
  public void testCalculateError(){
    Perceptron neuron = getTrainedNeuron();

    assertEquals(0, neuron.calculateError(0));
    assertEquals(0, neuron.calculateError(1));

    int anotherTrainData[][] = {{1,1,1,0,1,0,0,0,1},
                                {1,1,1,0,0,0,0,1,0}};
    neuron.trainData = anotherTrainData;

    assertEquals(-1, neuron.calculateError(0));
    assertEquals( 1, neuron.calculateError(1));
  }
}
