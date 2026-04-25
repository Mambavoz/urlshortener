package main

import (
	"fmt"
	"urlshort/internal/config"
)

func main() {
	// TODO: init config: cleanenv
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// TODO: init logger: log/slog

	// TODO: init storage: sqlite

	// TODO: init router: chi совместим с net/http, "chi render"

	// TODO: run server
}
