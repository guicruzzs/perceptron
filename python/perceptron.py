class Perceptron(object):

	def to_train(self, trainData, target, errorThreshold=0.1):
		self.errorThreshold = errorThreshold
		self.trainData = trainData
		self.target = target
		self.weights = [1]
		for i in range(len(trainData[0])):
			self.weights = self.weights + [1]

		it_is_over = False
		while not it_is_over:
			it_is_over = True
			for line in range(len(self.target)):
				error = self.calculateError(line)
				if (error != 0):
					self.calculateWeights(error, line)
					it_is_over = False

	def calculateOutput(self, data):
		total = self.weights[0]
		for index in range(len(data)):
			total += data[index] * self.weights[index+1]

		return 1 if(total >= 0) else 0

	def calculateError(self, line):
		return self.target[line] - self.calculateOutput(self.trainData[line])

	def calculateWeights(self, error, line):
		self.weights[0] = self.weights[0] + self.errorThreshold * error
		for index in range(len(self.trainData[line])):
			self.weights[index+1] = self.weights[index+1] + self.errorThreshold * error * self.trainData[line][index]