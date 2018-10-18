package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"

	"bitbucket.org/kit-systems/dari/pkg/database"
	"bitbucket.org/kit-systems/dari/pkg/parser"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

var (
	debug      = flag.Bool("debug", false, "Debug mode")
	stdout     = flag.String("stdout", "", "Custom stderr path")
	connString = flag.String("conn", "", "Connection string")
	limit      = flag.Uint("limit", 0, "Limit page number for parser")
	cfg        = zap.NewProductionConfig()
)

func dbConn(ctx context.Context) (*database.DB, error) {
	if *connString == "" {
		return nil, fmt.Errorf("empty connection string")
	}
	conn, err := sql.Open("mysql", *connString)
	if err != nil {
		return nil, err
	}
	return database.New(ctx, conn), nil
}

func main() {
	flag.Parse()
	if *debug {
		cfg = zap.NewDevelopmentConfig()
	}
	if *stdout != "" {
		cfg.OutputPaths = []string{
			*stdout,
		}
	}
	logger, err := cfg.Build()
	if err != nil {
		fmt.Print(err)
	}
	db, err := dbConn(context.Background())
	if err != nil {
		logger.Fatal("try to connect", zap.Error(err))
	}
	p, err := parser.New(db, parser.Options{
		Logger: logger,
		Debug:  *debug,
		Limit:  *limit,
	})
	if err != nil {
		logger.Fatal("new parser", zap.Error(err))
	}
	err = p.Run()
	if err != nil {
		logger.Fatal("run parser", zap.Error(err))
	}
}
