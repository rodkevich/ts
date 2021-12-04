-- v0.1.0

DROP TABLE if exists photos CASCADE;

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


--- photos -------------------------------------------------------------------------
CREATE TYPE enum_photos_type AS ENUM
    (
        'image/gif',
        'image/x-icon',
        'image/jpeg',
        'image/png',
        'image/tiff'
        );

CREATE table if not exists photos
(
    id          uuid             not null default gen_random_uuid(),
    type        enum_photos_type not null,
    size_kb     decimal(10, 2)   not null,
    upload_name citext           not null,
    image_url   varchar(250)     not null CHECK ( image_url <> '' ),
    description citext,
    owner_id    uuid             not null,
    created_at  timestamptz      not null default now(),
    updated_at  timestamptz      not null default now(),
    deleted     bool             not null default false,
    PRIMARY KEY (id)
);

CREATE trigger photos_touch_updated_at
    before update
    on photos
    FOR EACH ROW
EXECUTE procedure touch_updated_at();
--------------------------------------------------------------------------------

-- INSERT INTO public.photos (id, type, size_kb, upload_name, image_url, description,
--                            owner_id, created_at, updated_at, deleted)
-- VALUES (DEFAULT, 'image/jpeg'::enum_photos_type, 11.50, 'sam name'::citext,
--         'http://ssaas', null::citext, 'af0cca37-0703-495b-9872-42795f0f35f6', DEFAULT,
--         DEFAULT, DEFAULT);


-- UPDATE public.photos
-- SET upload_name = 'some name'::citext
-- WHERE id = 'b92d9697-7e82-4fc6-a8f3-875ddff33ad4';
--


-- UPDATE public.photos
-- SET type        = 'image/tiff'::enum_photos_type,
--     size_kb     = 22.50,
--     upload_name = 'some name new'::citext,
--     image_url   = 'http://ssaas/222',
--     description = 'new'::citext,
--     owner_id    = 'af0cca37-0703-495b-1872-42795f0f35f6'
-- WHERE id = 'b92d9697-7e82-4fc6-a8f3-875ddff33ad4';

