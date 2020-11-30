CREATE TABLE "medicines" (
  "medicine_id" SERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "company" varchar NOT NULL,
  "is_expired" boolean NOT NULL
);