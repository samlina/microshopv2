module user-service

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/gogo/protobuf v1.2.2-0.20190723190241-65acae22fc9d // indirect
	github.com/golang/protobuf v1.4.2
	github.com/jinzhu/gorm v1.9.16
	github.com/kr/pretty v0.2.0 // indirect
	github.com/micro/go-micro/v2 v2.9.1
	github.com/stretchr/testify v1.5.1 // indirect
	golang.org/x/crypto v0.0.0-20200709230013-948cd5f35899
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect

)
