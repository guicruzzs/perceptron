class Perceptron

  def to_train(train_data, target, error_threshold=0.1)
    @error_threshold = error_threshold
    @train_data = train_data
    @target = target
    @weights = Array.new(train_data.first.size + 1, 1)

    loop do
      it_is_over = true
      @target.size.times do |line|
        error = calculate_error(line)
        unless error == 0
          calculate_weights(error, line)
          it_is_over = false
        end
      end

      break if it_is_over
    end
  end

  def calculate_output(data)
    total = weights[0]
    data.each_with_index do |value, index|
      total += value * weights[index+1]
    end
    (total >= 0) ? 1 : 0
  end

  private
    attr_accessor :weights, :error_threshold, :train_data, :target

    def calculate_error(line)
      target[line] - calculate_output(train_data[line])
    end

    def calculate_weights(error, line)
      weights[0] = weights[0] + error_threshold * error
      train_data[line].each_with_index do |value, index|
        weights[index+1] = weights[index+1] + error_threshold * error * value
      end
    end
end