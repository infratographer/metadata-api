-- Drop index "status_metadata_id_status_namespace_id" from table: "status"
DROP INDEX "status_metadata_id_status_namespace_id";
-- Drop index "status_metadata_id_status_namespace_id_source" from table: "status"
DROP INDEX "status_metadata_id_status_namespace_id_source";
-- Create index "status_metadata_id_status_namespace_id" to table: "status"
CREATE INDEX "status_metadata_id_status_namespace_id" ON "status" ("metadata_id", "status_namespace_id");
-- Create index "status_metadata_id_status_namespace_id_source" to table: "status"
CREATE UNIQUE INDEX "status_metadata_id_status_namespace_id_source" ON "status" ("metadata_id", "status_namespace_id", "source");
