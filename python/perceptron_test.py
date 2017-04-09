import unittest
import perceptron

class PerceptronTest(unittest.TestCase):
	def getTrainedNeuron(self):
		neuron = perceptron.Perceptron()
		trainData = [[1,1,1,0,1,0,0,1,0], [1,0,1,1,1,1,1,0,1]]
		target = [0, 1]

		neuron.to_train(trainData, target)
		return neuron

	def test_to_train(self):
		neuron = self.getTrainedNeuron()
		self.assertEqual(-0.09999999999999987, neuron.weights[0])
		self.assertEqual(-0.09999999999999987, neuron.weights[1])
		self.assertEqual(-0.09999999999999987, neuron.weights[2])
		self.assertEqual(-0.09999999999999987, neuron.weights[3])
		self.assertEqual( 1.0,                 neuron.weights[4])
		self.assertEqual(-0.09999999999999987, neuron.weights[5])
		self.assertEqual( 1.0,                 neuron.weights[6])
		self.assertEqual( 1.0,                 neuron.weights[7])
		self.assertEqual(-0.09999999999999987, neuron.weights[8])

	def testCalculateOutput(self):
		neuron = self.getTrainedNeuron()
		sample_class_1 = [1,0,1,0,1,0,0,1,0]
		self.assertEqual(0, neuron.calculateOutput(sample_class_1))

		sample_class_2 = [1,0,1,1,1,1,0,0,1]
		self.assertEqual(1, neuron.calculateOutput(sample_class_2))

	def testCalculateError(self):
		neuron = self.getTrainedNeuron()
		self.assertEqual(0, neuron.calculateError(0))
		self.assertEqual(0, neuron.calculateError(1))

		another_trainData = [[1,1,1,0,1,0,0,0,1],[1,1,1,0,0,0,0,1,0]]
		neuron.trainData = another_trainData
		self.assertEqual(-1, neuron.calculateError(0))
		self.assertEqual( 1, neuron.calculateError(1))


if __name__ == '__main__':
	unittest.main()