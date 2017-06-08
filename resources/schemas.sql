CREATE TABLE IF NOT EXISTS public."users"
(
  id UUID    PRIMARY KEY NOT NULL,
  first_name VARCHAR(64),
  last_name  VARCHAR(64),
  email      VARCHAR(256) NOT NULL,
  password   VARCHAR(256),
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now(),
  deleted_at TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS user_id_uindex ON public."users" (id);
CREATE UNIQUE INDEX IF NOT EXISTS user_email_uindex ON public."users" (email);
