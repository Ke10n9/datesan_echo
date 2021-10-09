-- +migrate Up
create table if not exists menus_dishes (
  menu_id integer not null,
  dish_id integer not null,
  primary key (menu_id, dish_id),
  created_at datetime not null default current_timestamp,
  updated_at datetime not null default current_timestamp on update current_timestamp
);

-- +migrate Down
drop table if exists menus_dishes;