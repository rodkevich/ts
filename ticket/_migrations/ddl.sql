-- v0.1.0

DROP SCHEMA public CASCADE;

CREATE SCHEMA if not exists public;

-- EXTENSIONS ------------------------------------------------------------------

CREATE EXTENSION if not exists "pgcrypto";

--- COMMON FUNCTIONS -----------------------------------------------------------

CREATE or replace function touch_updated_at()
    returns trigger as
$$
begin
    NEW.updated_at = now();
    return NEW;
end;
$$ language 'plpgsql';

--- COMMON FUNCTIONS -----------------------------------------------------------

CREATE or replace function touch_updated_at()
    returns trigger as
$$
begin
    NEW.updated_at = now();
    return NEW;
end;
$$ language 'plpgsql';

--- TABLES ---------------------------------------------------------------------

--- tickets --------------------------------------------------------------------

CREATE TYPE enum_tickets_advantages_type AS ENUM
    ('Draft','Regular','Premium', 'Promoted');

CREATE table if not exists tickets
(
    id           uuid                    default gen_random_uuid(),
    owner_id     uuid           not null,
    name_short   varchar        not null,
    name_ext     varchar,
    description  text,
    amount       integer        not null default 1,
    price        decimal(10, 2) not null,
    currency     integer        not null,
    active       bool           not null default true,
    advantage    enum_tickets_advantages_type,
    published_at timestamptz,
    created_at   timestamptz    not null default now(),
    updated_at   timestamptz    not null default now(),
    deleted_at   timestamptz,
    PRIMARY KEY (id)

--  TODO: create event type for delete from here

);

CREATE trigger tickets_update_updated_at
    before update
    on tickets
    FOR EACH ROW
EXECUTE procedure touch_updated_at();
