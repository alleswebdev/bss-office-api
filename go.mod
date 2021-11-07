module github.com/ozonmp/bss-office-api

go 1.16

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/gammazero/workerpool v1.1.2
	github.com/golang/mock v1.6.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.6.0
	github.com/jackc/pgx/v4 v4.13.0
	github.com/jmoiron/sqlx v1.3.4
	github.com/lib/pq v1.10.3
	github.com/mattn/go-sqlite3 v1.14.9 // indirect
	github.com/opentracing/opentracing-go v1.2.0
	github.com/ozonmp/bss-office-api/pkg/bss-office-api v0.0.0-00010101000000-000000000000
	github.com/pressly/goose/v3 v3.3.1
	github.com/prometheus/client_golang v1.11.0
	github.com/rs/zerolog v1.24.0
	github.com/stretchr/testify v1.7.0
	github.com/uber/jaeger-client-go v2.29.1+incompatible
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	google.golang.org/grpc v1.42.0
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

replace github.com/ozonmp/bss-office-api/pkg/bss-office-api => ./pkg/bss-office-api
