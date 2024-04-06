package grpcserver_test

import (
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/maximprokopchuk/auto_reference_catalog_service/internal/config"
	"github.com/maximprokopchuk/auto_reference_catalog_service/internal/grpcserver"
	"github.com/maximprokopchuk/auto_reference_catalog_service/internal/store"
	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	cfg := config.NewConfig()
	_, err := toml.DecodeFile("../../configs/config.test.toml", cfg)
	assert.Nil(t, err)
	st := store.New(cfg.Store)
	server := grpcserver.New(st)
	assert.NotNil(t, server)
	assert.Equal(t, st, server.Store, "should include store")
	assert.NotNil(t, server.CreateCarModel)
	assert.NotNil(t, server.GetCarModelById)
	assert.NotNil(t, server.ListCarModels)
	assert.NotNil(t, server.DeleteCarModel)
	assert.NotNil(t, server.CreateComponent)
	assert.NotNil(t, server.GetTopLevelComponentsByCarModel)
	assert.NotNil(t, server.GetChildComponentsByComponent)
	assert.NotNil(t, server.DeleteComponent)
	assert.NotNil(t, server.UpdateComponent)
}
