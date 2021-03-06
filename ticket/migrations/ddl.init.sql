-- v0.1.0
BEGIN;

DROP TABLE if exists tickets CASCADE;
DROP TABLE if exists tags CASCADE;
DROP TABLE if exists ticket_tags CASCADE;

drop TYPE if exists enum_tickets_priority;

-- EXTENSIONS ------------------------------------------------------------------
CREATE EXTENSION if not exists "pgcrypto";
CREATE EXTENSION if not exists CITEXT;

--- COMMON FUNCTIONS -----------------------------------------------------------

CREATE or replace FUNCTION touch_updated_at()
    returns trigger as
$$
begin
    NEW.updated_at = (now() AT TIME ZONE 'utc');
    return NEW;
end;
$$ language 'plpgsql';


CREATE or replace FUNCTION get_tickets_by_tag(tag_uuid uuid)
    returns TABLE
            (
                ticket     uuid,
                owner_id   uuid,
                updated_at timestamptz,
                active     bool
            )
as
$$
begin
    return QUERY
        SELECT t.id, t.owner_id, t.updated_at, t.active
        FROM ticket_tags
                 left join tickets t on t.id = ticket_tags.ticket_id
        WHERE ticket_tags.tag_id = tag_uuid;
end;
$$ language 'plpgsql';

--- TABLES ---------------------------------------------------------------------

--- tickets --------------------------------------------------------------------
CREATE TYPE enum_tickets_priority AS ENUM
    ('Draft','Regular','Premium', 'Promoted');

CREATE table if not exists public.tickets
(
    id          uuid                  not null default (gen_random_uuid()),
    owner_id    uuid                  not null,
    name_short  citext                not null CHECK ( name_short <> '' ),
    name_ext    citext,
    description text,
    amount      integer               not null CHECK ( amount > 0 ),
    price       decimal(10, 2)        not null,
    currency    integer               not null,
    priority    enum_tickets_priority not null default 'Draft'::enum_tickets_priority,
    published   bool                  not null default false,
    active      bool                  not null default true,
    created_at  timestamp           not null default (now() AT TIME ZONE 'utc'),
    updated_at  timestamp          not null default (now() AT TIME ZONE 'utc'),
    deleted     bool                  not null default false,
    PRIMARY KEY (id)

);
DROP INDEX if exists idx_tickets_pagination;

CREATE INDEX idx_tickets_pagination
    ON tickets
        USING btree
        (created_at, id);


CREATE trigger tickets_touch_updated_at_trigger
    before update
    on public.tickets
    FOR EACH ROW
EXECUTE procedure touch_updated_at();


--- tags -----------------------------------------------------------------------
CREATE table if not exists tags
(
    id          uuid        not null default gen_random_uuid(),
    name        citext      not null,
    description citext      null,
    created_at  timestamp not null default (now() AT TIME ZONE 'utc'),
    updated_at  timestamp not null default (now() AT TIME ZONE 'utc'),
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
    created_at timestamp not null default (now() AT TIME ZONE 'utc'),
    updated_at timestamp not null default (now() AT TIME ZONE 'utc'),

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
COMMIT;
