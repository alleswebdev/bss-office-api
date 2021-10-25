package internal

//go:generate mockgen -destination=./mocks/repo_mock.go -package=mocks github.com/ozonmp/bss-office-api/internal/app/repo EventRepo
//go:generate mockgen -destination=./mocks/sender_mock.go -package=mocks github.com/ozonmp/bss-office-api/internal/app/sender EventSender
//go:generate mockgen -destination=./mocks/test_mock.go -package=mocks github.com/ozonmp/bss-office-api/internal/app/test Foo