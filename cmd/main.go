package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/AMOZINGAS/go-api-ecommerce/internal/env"
	"github.com/jackc/pgx/v5"
)

func main() {

	ctx := context.Background()

	cfg := config{
		addr: ":8080",
		db: dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING", "host=localhost user=postgres password=postgres dbname=ecom sslmode=disable"),
		},
	}

	//logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	//database
	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	logger.Info("Connected to database", "dsn", cfg.db.dsn)

	api := application{
		config: cfg,
		db:     conn,
	}

	if err := api.run(api.mounth()); err != nil {

		//este es un logger normal pero podemos usar uno mas estructurado o bonito
		//log.Printf("Server has failed to start, err: %s", err)

		//este podemos identificarlo como error y nos da mas informacion sobre el mensaje que enviamos
		//ya sea un simple texto o como aqui que es un error
		slog.Error("Server failed to start", "error", err)
		os.Exit(1)

	}

}
