package sdk

// download dependencies
//go:generate go-getter -progress https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto ./google/api
//go:generate go-getter -progress https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/resource.proto ./google/api
//go:generate go-getter -progress https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/field_behavior.proto ./google/api
//go:generate go-getter -progress https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto ./google/api
//go:generate go-getter -progress https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/httpbody.proto ./google/api
//go:generate go-getter -progress https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/client.proto ./google/api
//go:generate go-getter -progress https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/field_behavior.proto ./google/api

// genarate boilerplate code
//go:generate protoc -I . --go_out=plugins=grpc:$GOPATH/src/ form_example.proto
//go:generate protoc -I . --grpc-gateway_out=logtostderr=true,allow_patch_feature=true:$GOPATH/src/ form_example.proto
//go:generate protoc-go-inject-tag -input=form_example.pb.go
//go:generate protoc -I . --swagger_out=logtostderr=true,fqn_for_swagger_name=true,json_names_for_fields=false,simple_operation_ids=true:. ./form_example.proto
