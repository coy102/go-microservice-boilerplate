module go-microservices.org/core

go 1.16

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gogo/protobuf v1.2.2-0.20190723190241-65acae22fc9d // indirect
	github.com/golang/protobuf v1.4.2
	github.com/jinzhu/copier v0.0.0-20190924061706-b57f9002281a
	github.com/jmoiron/sqlx v1.2.0
	github.com/kr/pretty v0.2.0 // indirect
	github.com/lib/pq v1.3.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/kubernetes/v2 v2.9.1
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/stretchr/testify v1.5.1 // indirect
	golang.org/x/crypto v0.0.0-20200709230013-948cd5f35899 // indirect
	google.golang.org/protobuf v1.25.0
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
