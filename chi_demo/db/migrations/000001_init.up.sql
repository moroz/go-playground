create extension if not exists "citext" with schema "public";

create table users (
  id serial primary key,
  first_name text not null,
  email citext not null unique
);
