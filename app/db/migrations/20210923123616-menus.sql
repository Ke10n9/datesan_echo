-- +migrate Up
create table if not exists menus (
  id integer auto_increment primary key,
  user_id integer not null,
  time_id integer not null,
  date datetime not null,
  created_at datetime not null default current_timestamp,
  updated_at datetime not null default current_timestamp on update current_timestamp
);

-- +migrate Down
drop table if exists menus;
