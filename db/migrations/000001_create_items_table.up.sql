-- -----------------------------------------------------
-- Schema warning_tracker
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS "warning_tracker" ;

-- -----------------------------------------------------
-- Table "warning_tracker"."warning"
-- -----------------------------------------------------

CREATE TABLE IF NOT EXISTS "warning_tracker"."warning" (
  "id" UUID NOT NULL DEFAULT gen_random_uuid (),
  "branch" varchar(256) NOT NULL,
  "commit" varchar(41) NOT NULL,
  "count" integer NOT NULL,
  "created_by" varchar(128) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT "warning_tracker_pkey" PRIMARY KEY ("id")
);
