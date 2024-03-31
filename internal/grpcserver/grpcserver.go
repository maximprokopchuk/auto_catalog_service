package grpcserver

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/maximprokopchuk/auto_catalog_service/internal/api"
	"github.com/maximprokopchuk/auto_catalog_service/internal/sqlc"
	"github.com/maximprokopchuk/auto_catalog_service/internal/store"
)

type GRPCServer struct {
	Store *store.Store
}

func New(st *store.Store) *GRPCServer {
	return &GRPCServer{Store: st}
}

func (server *GRPCServer) CreateCarModel(ctx context.Context, req *api.CreateCarModelRequest) (*api.CarModel, error) {
	rec, err := server.Store.Queries.CreateCarModel(ctx, req.Name)

	if err != nil {
		return nil, err
	}

	return &api.CarModel{
		Id:   int32(rec.ID),
		Name: rec.Name,
	}, nil
}

func (server *GRPCServer) GetCarModelById(ctx context.Context, req *api.GetCarModelByIdRequest) (*api.CarModel, error) {
	rec, err := server.Store.Queries.GetCarModel(ctx, int64(req.Id))

	if err != nil {
		return nil, err
	}

	return &api.CarModel{
		Id:   int32(rec.ID),
		Name: rec.Name,
	}, nil
}

func (server *GRPCServer) ListCarModels(ctx context.Context, req *api.Empty) (*api.ListCarModelsResponse, error) {
	rec, err := server.Store.Queries.ListCarModels(ctx)

	if err != nil {
		return nil, err
	}

	var result []*api.CarModel

	for _, carModel := range rec {
		result = append(result, &api.CarModel{
			Id:   int32(carModel.ID),
			Name: carModel.Name,
		})
	}

	return &api.ListCarModelsResponse{Items: result}, nil
}

func (server *GRPCServer) DeleteCarModel(ctx context.Context, req *api.DeleteCarModelRequest) (*api.Empty, error) {
	err := server.Store.Queries.DeleteCarModel(ctx, int64(req.Id))
	if err != nil {
		return nil, err
	}

	return &api.Empty{}, nil
}

func (server *GRPCServer) GetTopLevelComponentsForCarModel(ctx context.Context, req *api.GetTopLevelComponentsForCarModelRequest) (*api.ListComponentResponse, error) {

	rec, err := server.Store.Queries.GetTopLevelComponentsForCarModel(ctx, pgtype.Int4{Int32: req.CarModelId, Valid: true})
	if err != nil {
		return nil, err
	}

	var result []*api.Component

	for _, component := range rec {
		result = append(result, &api.Component{
			Id:         int32(component.ID),
			Name:       component.Name,
			CarModelId: component.CarModelID.Int32,
			ParentId:   component.ParentID.Int32,
		})
	}

	return &api.ListComponentResponse{Items: result}, nil
}

func (server *GRPCServer) CreateComponent(ctx context.Context, req *api.CreateComponentRequest) (*api.Component, error) {
	params := sqlc.CreateComponentParams{
		Name: req.Name,
	}

	if req.CarModelId != 0 {
		params.CarModelID = pgtype.Int4{Int32: req.CarModelId, Valid: true}
	} else if req.ParentId != 0 {
		params.ParentID = pgtype.Int4{Int32: req.ParentId, Valid: true}
	}
	rec, err := server.Store.Queries.CreateComponent(ctx, params)

	if err != nil {
		return nil, err
	}

	return &api.Component{
		Id:         int32(rec.ID),
		Name:       rec.Name,
		ParentId:   rec.ParentID.Int32,
		CarModelId: rec.CarModelID.Int32,
	}, nil
}

func (server *GRPCServer) GetChildComponentsForComponent(ctx context.Context, req *api.GetChildComponentsForComponentRequest) (*api.ListComponentResponse, error) {

	rec, err := server.Store.Queries.GetChildComponentsForComponent(ctx, pgtype.Int4{Int32: req.ParentId, Valid: true})
	if err != nil {
		return nil, err
	}

	var result []*api.Component

	for _, component := range rec {
		result = append(result, &api.Component{
			Id:         int32(component.ID),
			Name:       component.Name,
			CarModelId: component.CarModelID.Int32,
			ParentId:   component.ParentID.Int32,
		})
	}

	return &api.ListComponentResponse{Items: result}, nil
}

func (server *GRPCServer) DeleteComponent(ctx context.Context, req *api.DeleteComponentRequest) (*api.Empty, error) {
	err := server.Store.Queries.DeleteComponent(ctx, int64(req.Id))
	if err != nil {
		return nil, err
	}

	return &api.Empty{}, nil
}

func (server *GRPCServer) UpdateComponent(ctx context.Context, req *api.UpdateComponentRequest) (*api.Component, error) {
	rec, err := server.Store.Queries.UpdateComponent(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	return &api.Component{
		Id:         int32(rec.ID),
		ParentId:   rec.ParentID.Int32,
		CarModelId: rec.CarModelID.Int32,
		Name:       rec.Name,
	}, nil
}
