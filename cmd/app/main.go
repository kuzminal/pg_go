package main

import (
	"context"
	"fmt"
	"os"
	"pg_go/internal/godb"
	"pg_go/pkg/helpers/pg"
)

func main() {
	cfg := &pg.Config{}
	cfg.Host = "localhost"
	cfg.Username = "db_user"
	cfg.Password = "pwd123"
	cfg.Port = "54320"
	cfg.DbName = "db_test"
	cfg.Timeout = 5

	//Создаем конфиг для пула
	poolConfig, err := pg.NewPoolConfig(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Pool config error: %v\n", err)
		os.Exit(1)
	}
	poolConfig.MaxConns = 5
	//poolConfig.MinConns = 20

	//Создаем пул подключений
	c, err := pg.NewConnection(poolConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Connect to database failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Connection OK!")

	//Проверяем подключение
	_, err = c.Exec(context.Background(), ";")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ping failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Ping OK!")

	ins := &godb.Instance{Db: c}
	ins.Start()
}
