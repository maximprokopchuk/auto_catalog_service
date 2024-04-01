-- +goose Up
-- +goose StatementBegin
CREATE TABLE car_model (
  id   BIGSERIAL PRIMARY KEY,
  name text      NOT NULL
);

CREATE TABLE component (
  id   BIGSERIAL PRIMARY KEY,
  name text      NOT NULL,
  car_model_id INTEGER NULL,
  parent_id INTEGER NULL,

  FOREIGN KEY (car_model_id) REFERENCES car_model(id) ON DELETE CASCADE,
  FOREIGN KEY (parent_id) REFERENCES component(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE component; 
DROP TABLE car_model;
-- +goose StatementEnd
