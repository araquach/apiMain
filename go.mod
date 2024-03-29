module github.com/araquach/apiMain

go 1.18

require (
	github.com/araquach/apiAuth v0.0.1
	github.com/araquach/apiFinance23 v0.0.1
	github.com/araquach/apiHelpers v0.0.1
	github.com/araquach/apiTeam v0.0.1
	github.com/araquach/apiTime v0.0.2
	github.com/araquach/dbService v0.0.1
	github.com/gorilla/mux v1.8.0
	github.com/joho/godotenv v1.5.1
	github.com/rs/cors v1.9.0
)

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.3.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/crypto v0.9.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	gorm.io/datatypes v1.2.0 // indirect
	gorm.io/driver/mysql v1.4.7 // indirect
	gorm.io/driver/postgres v1.5.2 // indirect
	gorm.io/gorm v1.25.1 // indirect
)

replace (
	github.com/araquach/apiAuth => ../apiAuth
	github.com/araquach/apiFinance23 => ../apiFinance23
	github.com/araquach/apiHelpers => ../apiHelpers
	github.com/araquach/apiTeam => ../apiTeam
	github.com/araquach/apiTime => ../apiTime
	github.com/araquach/dbService => ../dbService
)
