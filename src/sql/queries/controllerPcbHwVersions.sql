-- name: GetAllControllerPcbHwVersions :many
SELECT *
FROM controller_pcb_hw_versions;

-- name: CreateControllerPcbHwVersion :one
BEGIN;
    -- Insert revision if not exists
    INSERT INTO controller_pcb_hw_versions_rev (revision)
    VALUES ($2) 
    ON CONFLICT (revision) DO NOTHING;

    -- Insert the main record
    INSERT INTO controller_pcb_hw_versions (version_number, revision)
    VALUES ($1, $2); 
COMMIT;