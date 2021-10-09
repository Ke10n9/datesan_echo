-- +migrate Up
create table if not exists dishes (
  id integer auto_increment primary key,
  name varchar(31) not null,
  created_at datetime not null default current_timestamp,
  updated_at datetime not null default current_timestamp on update current_timestamp
);

-- +migrate Down
drop table if exists dishes;