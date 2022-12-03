create table urls (
  id serial primary key,
  url text not null,
  short_id text not null,
  inserted_at timestamp(0) default (now() at time zone 'utc'),
  updated_at timestamp(0) default (now() at time zone 'utc')
);

create unique index urls_short_id_idx on urls (short_id);
