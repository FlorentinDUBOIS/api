CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS public."users"
(
  uid UUID DEFAULT gen_random_uuid() PRIMARY KEY NOT NULL,
  first_name VARCHAR(64),
  last_name  VARCHAR(64),
  email     VARCHAR(256),
  password  VARCHAR(256),
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS user_uid_uindex ON public."users" (UID);
CREATE UNIQUE INDEX IF NOT EXISTS user_email_uindex ON public."users" (Email);