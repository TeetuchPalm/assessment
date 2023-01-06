-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS news_articles_id_seq;

-- Table Definition
CREATE TABLE "news_articles" (
    "id" int4 NOT NULL DEFAULT nextval('news_articles_id_seq'::regclass),
    "title" TEXT,
	"amount" FLOAT,
	"note" TEXT,
	"tags" TEXT[]
    PRIMARY KEY ("id")
);

INSERT INTO "expenses" ("id", "title", "amount", "note", "tags") values (1, 'nut', 27, 'ass', '{"zzzz","ssss"}');