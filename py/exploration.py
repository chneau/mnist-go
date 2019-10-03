from __future__ import print_function

import json

import keras
import keras2onnx
import onnx
from keras import backend as K
from keras.datasets import mnist
from keras.layers import Conv2D, Dense, Dropout, Flatten, MaxPooling2D
from keras.models import Sequential, load_model
from keras.optimizers import RMSprop

batch_size = 128
num_classes = 10
epochs = 1

# the data, split between train and test sets
(x_train, y_train), (x_test, y_test) = mnist.load_data()

x_train = x_train.reshape(60000, 784)
x_test = x_test.reshape(10000, 784)

model = load_model('model.h5')

res = model.predict(x_test)
print(x_test[0].tolist())
print(res[0])
