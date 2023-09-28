
CREATE TABLE "users" (
     "id" serial not null,
     "number" varchar unique not null,
     "name" varchar not null,
     "nik" varchar unique not null,
     "phone" varchar unique not null,
     "balance" numeric not null,
     "created_at" timestamptz not null default current_timestamp,
     "updated_at" timestamptz not null default current_timestamp
);
