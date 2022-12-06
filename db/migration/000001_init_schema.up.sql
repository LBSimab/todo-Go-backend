CREATE TABLE "tasks" (
  "id" serial UNIQUE PRIMARY KEY,
  "name" varchar,
  "finished" boolean,
  "supervisor" int,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "category" varchar
); 

CREATE TABLE "users" (
  "id" serial UNIQUE PRIMARY KEY,
  "full_name" varchar,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "supervisor" boolean,
  "commits" int DEFAULT 0
);

CREATE TABLE "commits" (
  "title" varchar,
  "commit_id" serial UNIQUE PRIMARY KEY,
  "user_id" int,
  "task_id" int,
  "comment" varchar,
  "category" varchar,
  "supervisor_id" int
);

CREATE INDEX ON "tasks" ("name", "finished", "category");

CREATE INDEX ON "users" ("full_name", "supervisor");

CREATE INDEX ON "commits" ("title", "category", "supervisor_id", "user_id");

ALTER TABLE "tasks" ADD FOREIGN KEY ("supervisor") REFERENCES "users" ("id");

ALTER TABLE "commits" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "commits" ADD FOREIGN KEY ("task_id") REFERENCES "Tasks" ("id");

ALTER TABLE "commits" ADD FOREIGN KEY ("supervisor_id") REFERENCES "users" ("id");
