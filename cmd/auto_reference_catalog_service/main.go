package main

import (
	"context"
	"flag"
	"log"
	"net"

	"github.com/BurntSushi/toml"
	"github.com/maximprokopchuk/auto_reference_catalog_service/internal/config"
	"github.com/maximprokopchuk/auto_reference_catalog_service/internal/grpcserver"
	"github.com/maximprokopchuk/auto_reference_catalog_service/internal/store"
	"github.com/maximprokopchuk/auto_reference_catalog_service/pkg/api"
	"google.golang.org/grpc"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config_path", "configs/config.toml", "path to config TOML file")
}

func run() error {
	ctx := context.Background()
	cfg := config.NewConfig()
	_, err := toml.DecodeFile(configPath, cfg)

	if err != nil {
		return err
	}

	st := store.New(cfg.Store)

	if err := st.Open(ctx); err != nil {
		return err
	}

	defer st.Close(ctx)

	s := grpc.NewServer()
	srv := grpcserver.New(st)
	api.RegisterAutoCatalogServiceServer(s, srv)

	l, err := net.Listen("tcp", ":"+cfg.BindUrl)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("LISTEN " + cfg.BindUrl)

	if err = s.Serve(l); err != nil {
		log.Fatal(err)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
