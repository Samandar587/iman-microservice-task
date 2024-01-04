CREATE TABLE "posts" (
    "id" SERIAL PRIMARY KEY NOT NULL,
    "user_id" INTEGER,
    "title" VARCHAR(255),
    "body" TEXT,
    "page" INTEGER
)
