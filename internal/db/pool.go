package db

import (
	"log"

	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/stdlib"

	"github.com/jmoiron/sqlx"
)

func NewPGXConnectionPool() (db *sqlx.DB) {

	var connString = "postgres://localhost/postgres?sslmode=disable&user=postgres&password=postgres&pool_max_conns=10&pool_min_conns=3"
	// db, err := sqlx.Connect("pgx", "postgres://pgx_md5:secret@localhost:5432/pgx_test")
	db, err := sqlx.Connect("pgx", connString)
	if err != nil {
		log.Fatalln(err)
	}

	if err != nil {
		log.Printf("Unable to connect database: %v\n", err)
		return
	}

	log.Printf("New PG datasource connected to: %v", connString)
	return
}

// func NewPGXConnectionPool(ctx context.Context, cfg *configs.Configs) (pool *pgxpool.Pool, err error) {
//
// 	var connString string
// 	connString = fmt.Sprintf(
// 		"%s://%s/%s?sslmode=%s&user=%s&password=%s&pool_max_conns=%v&pool_min_conns=%v",
// 		cfg.Database.Driver,
// 		cfg.Database.Host,
// 		cfg.Database.Name,
// 		cfg.Database.SslMode,
// 		cfg.Database.User,
// 		cfg.Database.Pass,
// 		cfg.Database.MaxPoolConnections,
// 		cfg.Database.MinPoolConnections,
// 	)
// 	pool, err = pgxpool.Connect(ctx, connString)
//
// 	if err != nil {
// 		log.Printf("Unable to connect database: %v\n", err)
// 		return nil, err
// 	}
// 	log.Printf("New PG datasource connected to: %v", connString)
// 	return pool, nil
// }
