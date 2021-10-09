-- +migrate Up
create table if not exists users (
  id integer auto_increment primary key,
  email varchar(255) unique not null,
  password varchar(255) not null,
  created_at datetime not null default current_timestamp,
  updated_at datetime not null default current_timestamp on update current_timestamp
);

-- +migrate Down
drop table if exists users;