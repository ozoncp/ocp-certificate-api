CREATE TABLE certificate
(
    id      SERIAL PRIMARY KEY,
    user_id INT,
    created TIMESTAMP(0) WITH TIME ZONE,
    link    VARCHAR(255)
);