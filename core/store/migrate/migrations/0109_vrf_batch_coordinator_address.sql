-- +goose Up
ALTER TABLE vrf_specs ADD COLUMN batch_coordinator_address bytea;

-- +goose Down
ALTER TABLE vrf_specs DROP COLUMN batch_coordinator_address;
