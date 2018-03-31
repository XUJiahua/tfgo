# tfgo

generated golang gRPC pb by below commands:

```
git clone git@github.com:tensorflow/serving.git
cd serving
git clone --recursive https://github.com/tensorflow/tensorflow.git

cd ..

mkdir -p $GOPATH/src/github.com/XUJiahua//tfgo

protoc -I=serving -I serving/tensorflow --go_out=plugins=grpc:$GOPATH/src/github.com/XUJiahua/tfgo serving/tensorflow_serving/apis/*.proto
protoc -I=serving/tensorflow --go_out=plugins=grpc:$GOPATH/src/github.com/XUJiahua//tfgo serving/tensorflow/tensorflow/core/framework/*.proto
protoc -I=serving/tensorflow --go_out=plugins=grpc:$GOPATH/src/github.com/XUJiahua/tfgo serving/tensorflow/tensorflow/core/protobuf/{saver,meta_graph}.proto
protoc -I=serving/tensorflow --go_out=plugins=grpc:$GOPATH/src/github.com/XUJiahua/tfgo serving/tensorflow/tensorflow/core/example/*.proto

```

ref: https://gist.github.com/mauri870/1f953a183ee6c186e70a0a72e78b088c
