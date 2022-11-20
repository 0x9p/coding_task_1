CREATE TABLE port
(
    id          varchar(255) not null primary key,
    name        varchar(255) not null,
    city        varchar(255) not null,
    country     varchar(255) not null,
    alias       text[] null,
    regions     text[] null,
    coordinates double precision[] null,
    province    varchar(255) null,
    timezone    varchar(255) null,
    unlocs      text[] not null,
    code        varchar(255) not null
);
