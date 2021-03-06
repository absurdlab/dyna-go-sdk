// pb contains the protobuf definition and its generated code. The protocol is designed to resemble
// idp.Interface. By doing so, adapter implementations can be provided in order to free users from
// having to worry about the GRPC details. See the github.com/absurdlab/dyna-go-sdk/idp/grpc package
// for client and server implementations.
//
// As a note, the idp.pb.go file in this directory is generated by the command:
//
//	// Install protobuf package on the OS before proceeding.
//	// For Mac OS, run `brew install protobuf`
//	protoc -I=./ --go_out=plugins=grpc:. idp.proto
package pb
