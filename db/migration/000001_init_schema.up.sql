CREATE TABLE "user_watches" (
    "id" bigserial PRIMARY KEY,
    "user_id" UUID NOT NULL UNIQUE, -- Add UNIQUE constraint here
    "watch_reason" TEXT NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "fraud_alerts" (
    "id" bigserial PRIMARY KEY,
    "alert_id" SERIAL,
    "user_id" UUID NOT NULL,
    "alert_reason" TEXT NOT NULL,
    "alert_level" INT CHECK (alert_level BETWEEN 1 AND 5),
    "slug" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user_watches (user_id)
);

CREATE TABLE "transactions" (
    "id" bigserial PRIMARY KEY,
    "from_user" bigint NOT NULL,
    "to_user" bigint NOT NULL,
    "amount" bigint NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE "transactions" ADD FOREIGN KEY ("from_user") REFERENCES "user_watches" ("id");
ALTER TABLE "transactions" ADD FOREIGN KEY ("to_user") REFERENCES "user_watches" ("id");

CREATE INDEX ON "user_watches" ("user_id");

CREATE INDEX ON "transactions" ("from_user");
CREATE INDEX ON "transactions" ("to_user");
CREATE INDEX ON "transactions" ("from_user", "to_user");

COMMENT ON COLUMN "transactions"."amount" IS 'must be positive';
