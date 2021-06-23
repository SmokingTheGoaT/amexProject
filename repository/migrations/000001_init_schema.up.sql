CREATE TABLE "product" (
                           "id" bigserial PRIMARY KEY,
                           "title" varchar NOT NULL,
                           "description" text NOT NULL,
                           "vendor" varchar NOT NULL,
                           "product_type" varchar NOT NULL,
                           "created_at" timestamptz NOT NULL DEFAULT 'now()',
                           "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);
