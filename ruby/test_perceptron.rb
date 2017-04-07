require 'minitest/autorun'
require './perceptron'

class TestPerceptron < Minitest::Test
  def setup
    @perceptron = Perceptron.new
    train_data = [[1,1,1,0,1,0,0,1,0],
                  [1,0,1,1,1,1,1,0,1]]
    target = [0, 1]

    @perceptron.to_train(train_data, target)
  end

  def test_to_train
    weights = @perceptron.send(:weights)

    assert_equal -0.09999999999999987, weights[0]
    assert_equal -0.09999999999999987, weights[1]
    assert_equal -0.09999999999999987, weights[2]
    assert_equal  1.0,                 weights[3]
    assert_equal -0.09999999999999987, weights[4]
    assert_equal  1.0,                 weights[5]
    assert_equal  1.0,                 weights[6]
    assert_equal -0.09999999999999987, weights[7]
    assert_equal  1.0,                 weights[8]
  end

  def test_calculate_output
    sample = [1,0,1,0,1,0,0,1,0]
    assert_equal 0, @perceptron.calculate_output(sample)

    other_class_sample = [1,0,1,1,1,1,0,0,1]
    assert_equal 1, @perceptron.calculate_output(other_class_sample)
  end

  def test_calculate_error
    assert_equal 0, @perceptron.send(:calculate_error, 0)
    assert_equal 0, @perceptron.send(:calculate_error, 1)

    another_train_data = [[1,1,1,0,1,0,0,0,1], [1,1,1,0,0,0,0,1,0]]
    @perceptron.send(:train_data=, another_train_data)

    assert_equal -1, @perceptron.send(:calculate_error, 0)
    assert_equal  1, @perceptron.send(:calculate_error, 1)
  end
end