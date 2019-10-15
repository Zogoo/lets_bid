-- drop sequence if exists users_sequence;
drop database if exists lets_bid;
create database lets_bid;
\c lets_bid;

drop table if exists tokens;
drop table if exists users;

create table public.users (
  id serial primary key,
  name character varying,
  email character varying,
  password character varying
);

-- create sequence public.users_sequence 
--   start with 1
--   increment by 1
--   no minvalue
--   no maxvalue
--   cache 1;

create table public.tokens (
  user_id int not null,
  jwt_token character varying,
  CONSTRAINT user_tokens_id FOREIGN KEY (user_id) REFERENCES users(id)
);