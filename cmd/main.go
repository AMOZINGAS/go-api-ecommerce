package main

import (
	"log/slog"
	"os"
)

func main() {
	cfg := config{
		addr: ":8080",
		db:   dbConfig{},
	}

	api := application{
		config: cfg,
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	slog.SetDefault(logger)

	if err := api.run(api.mounth()); err != nil {

		//este es un logger normal pero podemos usar uno mas estructurado o bonito
		//log.Printf("Server has failed to start, err: %s", err)

		//este podemos identificarlo como error y nos da mas informacion sobre el mensaje que enviamos
		//ya sea un simple texto o como aqui que es un error
		slog.Error("Server failed to start", "error", err)
		os.Exit(1)

	}

}
