-- v0.1.0

DROP TABLE if exists tickets CASCADE;

-- EXTENSIONS ------------------------------------------------------------------
CREATE EXTENSION if not exists "pgcrypto";
CREATE EXTENSION if not exists CITEXT;

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

--- tickets -------------------------------------------------------------------------
CREATE TYPE enum_tickets_priority AS ENUM
    ('Draft','Regular','Premium', 'Promoted');

CREATE table if not exists tickets
(
    id          uuid                  not null default gen_random_uuid(),
    owner_id    uuid                  not null,
    name        citext                not null CHECK ( name <> '' ),
    name_ext    citext,
    description text,
    amount      integer               not null default 1::integer,
    price       decimal(10, 2)        not null,
    currency    integer               not null,
    priority    enum_tickets_priority not null default 'Draft'::enum_tickets_priority,
    published   bool                  not null default false,
    created_at  timestamptz           not null default now(),
    updated_at  timestamptz           not null default now(),
    deleted     bool                  not null default false,

    PRIMARY KEY (id)
);

CREATE trigger tickets_touch_updated_at
    before update
    on tickets
    FOR EACH ROW
EXECUTE procedure touch_updated_at();
--------------------------------------------------------------------------------
