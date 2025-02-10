CREATE TABLE tokens (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT,
    token_hash BYTEA NOT NULL,
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    expired_at TIMESTAMP(0) WITH TIME ZONE NOT NULL
);

-- Should be like this
-- CREATE TABLE tokens (
--                         token_hash BYTEA PRIMARY KEY,
--                         user_id BIGINT,
--                         created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
--                         expired_at TIMESTAMP(0) WITH TIME ZONE NOT NULL
-- );

