protoc greet/greetPb/greet.proto --go-grpc_out=.

protoc greet/greetPb/UnaryGreet.proto --go-grpc_out=.

protoc greet/greetPb/UnaryGreet.proto --gogo_out=plugins=grpc:.

protoc calculator/calculator_proto_ex/calculator_proto_experimental.proto --gogo_out=plugins=grpc:.

protoc greet/greetPb/streamming.proto --gogo_out=plugins=grpc:greet/

protoc greet/greetPb/ClientStreaming.proto --gogo_out=plugins=grpc:greet/

protoc greet/greetPb/BiDirectionalStreaming.proto --gogo_out=plugins=grpc:greet/

protoc greet/greetPb/greetWithDeadline.proto --gogo_out=plugins=grpc:greet/

protoc blog/blogpb/blog.proto --gogo_out=plugins=grpc:blog/