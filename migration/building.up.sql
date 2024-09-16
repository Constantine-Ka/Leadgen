create table if not exists public.buildings
(
    id     SERIAL PRIMARY KEY,
    name   varchar(500),
    city   varchar(100),
    year   integer,
    level integer
    );

alter table public.buildings
    owner to docker;

create index if not exists city
    on public.buildings (city);

create index if not exists name
    on public.buildings (name);

create index if not exists year
    on public.buildings (year);

