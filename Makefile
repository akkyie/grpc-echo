echo.pb.go: echo.proto
	protoc --go_out=plugins=grpc:echo $^
