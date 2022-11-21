create table cities
(
    id   varchar default gen_random_uuid() not null
        constraint cities_pkey
            primary key,
    name varchar                           not null
        constraint names
            unique,
    lat  numeric                           not null,
    lon  numeric                           not null
);

alter table cities
    owner to postgres;

create table temperature
(
    id        varchar default gen_random_uuid() not null
        constraint temperature_pkey
            primary key,
    city_name varchar                           not null,
    temp      numeric                           not null,
    dt        integer                           not null
);

alter table temperature
    owner to postgres;

create table avg_temperature
(
    id         varchar default gen_random_uuid() not null
        constraint avg_temperature_pkey
            primary key,
    temp       numeric                           not null,
    start_date integer                           not null,
    end_date   integer                           not null,
    city_name  varchar                           not null
);

alter table avg_temperature
    owner to postgres;