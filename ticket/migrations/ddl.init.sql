-- v0.1.0

DROP TABLE if exists tickets CASCADE;
DROP TABLE if exists tags CASCADE;
DROP TABLE if exists ticket_tags CASCADE;

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

--- tickets --------------------------------------------------------------------
CREATE TYPE enum_tickets_priority AS ENUM
    ('Draft','Regular','Premium', 'Promoted');

CREATE table if not exists tickets
(
    id          uuid                  not null default gen_random_uuid(),
    owner_id    uuid                  not null,
    name_short  citext                not null CHECK ( name_short <> '' ),
    name_ext    citext,
    description text,
    amount      integer               not null default 1::integer,
    price       decimal(10, 2)        not null,
    currency    integer               not null,
    priority    enum_tickets_priority not null default 'Draft'::enum_tickets_priority,
    published   bool                  not null default false,
    active      bool                  not null default true,
    created_at  timestamptz           not null default (now() AT TIME ZONE 'utc'),
    updated_at  timestamptz           not null default (now() AT TIME ZONE 'utc'),
    deleted     bool                  not null default false,
    PRIMARY KEY (id)
);

CREATE trigger tickets_touch_updated_at_trigger
    before update
    on tickets
    FOR EACH ROW
EXECUTE procedure touch_updated_at();


--- tags -----------------------------------------------------------------------
CREATE table if not exists tags
(
    id          uuid        not null default gen_random_uuid(),
    name        citext      not null,
    description citext      null,
    created_at  timestamptz not null default now(),
    updated_at  timestamptz not null default now(),
    deleted     bool        not null default false,
    PRIMARY KEY (id)
);

CREATE trigger tags_touch_updated_at
    before update
    on tags
    FOR EACH ROW
EXECUTE procedure touch_updated_at();


--- ticket_tags ----------------------------------------------------------------
CREATE table if not exists ticket_tags
(
    ticket_id  uuid        not null,
    tag_id     uuid        not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),

    PRIMARY KEY (tag_id, ticket_id),

    CONSTRAINT fk_ticket_tags_tag_id
        foreign key (tag_id)
            references tags (id)
            ON DELETE CASCADE,

    CONSTRAINT fk_ticket_tags_ticket_id
        foreign key (ticket_id)
            references tickets (id)
            ON DELETE CASCADE
);

CREATE trigger ticket_tags_touch_updated_at
    before update
    on ticket_tags
    FOR EACH ROW
EXECUTE procedure touch_updated_at();

--------------------------------------------------------------------------------
