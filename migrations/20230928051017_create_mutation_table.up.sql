
CREATE TABLE "mutations" (
     "id" serial not null,
     "number" varchar not null,
     "code" varchar not null,
     "amount" numeric not null,
     "created_at" timestamptz not null default current_timestamp,
     "updated_at" timestamptz not null default current_timestamp
);
