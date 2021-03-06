-- v0.1.0

DROP TABLE if exists customers CASCADE ;

-- EXTENSIONS ------------------------------------------------------------------
CREATE EXTENSION if not exists "pgcrypto";

--- COMMON FUNCTIONS -----------------------------------------------------------
CREATE or replace FUNCTION touch_updated_at()
    returns trigger as
$$
begin
    NEW.updated_at = now();
    return NEW;
end;
$$ language 'plpgsql';

--- TABLES ---------------------------------------------------------------------

--- customers ------------------------------------------------------------------
CREATE TYPE enum_customers_type AS ENUM
    ('User', 'Application');

CREATE TYPE enum_customers_status AS ENUM
    ('Active', 'Pending', 'Blocked');

CREATE table if not exists customers
(
    id         uuid                  default gen_random_uuid(),
    type       enum_customers_type   default 'User'::enum_customers_type,
    status     enum_customers_status default 'Pending'::enum_customers_status,
    login      varchar(60)  not null unique CHECK ( login <> '' ),
    password   varchar(100) not null CHECK ( octet_length(password) <> 0 ),
    identity   text, -- how to make this unique if presented??
    created_at timestamptz  not null default now(),
    updated_at timestamptz  not null default now(),
    deleted    bool         not null default false,
    PRIMARY KEY (id)
);

CREATE trigger customers_touch_updated_at
    before update
    on customers
    FOR EACH ROW
EXECUTE procedure touch_updated_at();

--------------------------------------------------------------------------------
