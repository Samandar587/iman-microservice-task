CREATE TABLE "posts" (
    "id" SERIAL PRIMARY KEY NOT NULL,
    "original_post_id" INTEGER,
    "user_id" INTEGER,
    "title" VARCHAR(255),
    "body" TEXT,
    "page" INTEGER
)
 