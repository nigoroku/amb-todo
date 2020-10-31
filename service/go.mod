module github.com/nigoroku/amb-todo/service

go 1.14

require (
	github.com/friendsofgo/errors v0.9.2 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gofrs/uuid v3.3.0+incompatible // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/stretchr/testify v1.2.2
	github.com/volatiletech/inflect v0.0.1 // indirect
	github.com/volatiletech/null v8.0.0+incompatible // indirect
	github.com/volatiletech/sqlboiler v3.7.1+incompatible
	local.packages/models v0.0.0-00010101000000-000000000000

)

replace local.packages/models => ./../models
