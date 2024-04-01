// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createCarModel = `-- name: CreateCarModel :one
INSERT INTO car_model (
  name
) VALUES (
  $1
)
RETURNING id, name
`

func (q *Queries) CreateCarModel(ctx context.Context, name string) (CarModel, error) {
	row := q.db.QueryRow(ctx, createCarModel, name)
	var i CarModel
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const createComponent = `-- name: CreateComponent :one
INSERT INTO component (
  name, car_model_id, parent_id
) VALUES (
  $1, $2, $3
)
RETURNING id, name, car_model_id, parent_id
`

type CreateComponentParams struct {
	Name       string
	CarModelID pgtype.Int4
	ParentID   pgtype.Int4
}

func (q *Queries) CreateComponent(ctx context.Context, arg CreateComponentParams) (Component, error) {
	row := q.db.QueryRow(ctx, createComponent, arg.Name, arg.CarModelID, arg.ParentID)
	var i Component
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CarModelID,
		&i.ParentID,
	)
	return i, err
}

const deleteCarModel = `-- name: DeleteCarModel :exec
DELETE FROM car_model
WHERE id = $1
`

func (q *Queries) DeleteCarModel(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteCarModel, id)
	return err
}

const deleteComponent = `-- name: DeleteComponent :exec
DELETE FROM component
WHERE id = $1
`

func (q *Queries) DeleteComponent(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteComponent, id)
	return err
}

const getCarModel = `-- name: GetCarModel :one
SELECT id, name FROM car_model
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCarModel(ctx context.Context, id int64) (CarModel, error) {
	row := q.db.QueryRow(ctx, getCarModel, id)
	var i CarModel
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getChildComponentsByComponent = `-- name: GetChildComponentsByComponent :many
SELECT id, name, car_model_id, parent_id FROM component
WHERE parent_id = $1 LIMIT 1
`

func (q *Queries) GetChildComponentsByComponent(ctx context.Context, parentID pgtype.Int4) ([]Component, error) {
	rows, err := q.db.Query(ctx, getChildComponentsByComponent, parentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Component
	for rows.Next() {
		var i Component
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CarModelID,
			&i.ParentID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTopLevelComponentsByCarModel = `-- name: GetTopLevelComponentsByCarModel :many
SELECT id, name, car_model_id, parent_id FROM component
WHERE car_model_id = $1 LIMIT 1
`

func (q *Queries) GetTopLevelComponentsByCarModel(ctx context.Context, carModelID pgtype.Int4) ([]Component, error) {
	rows, err := q.db.Query(ctx, getTopLevelComponentsByCarModel, carModelID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Component
	for rows.Next() {
		var i Component
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CarModelID,
			&i.ParentID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listCarModels = `-- name: ListCarModels :many
SELECT id, name FROM car_model
ORDER BY name
`

func (q *Queries) ListCarModels(ctx context.Context) ([]CarModel, error) {
	rows, err := q.db.Query(ctx, listCarModels)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CarModel
	for rows.Next() {
		var i CarModel
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateComponent = `-- name: UpdateComponent :one
UPDATE component
SET name = $1
RETURNING id, name, car_model_id, parent_id
`

func (q *Queries) UpdateComponent(ctx context.Context, name string) (Component, error) {
	row := q.db.QueryRow(ctx, updateComponent, name)
	var i Component
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CarModelID,
		&i.ParentID,
	)
	return i, err
}
