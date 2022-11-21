
create table users (
  id         int primary key,
  name       varchar(255),
  email      varchar(255) not null unique,
  password   varchar(255) not null,
  created_at text not null 
);

create table sessions (
  id         int primary key,
  uuid       varchar(64) not null unique,
  email      varchar(255),
  user_id    int references users(id),
  created_at text not null 
);

create table threads (
  id         int primary key,
  topic      text,
  user_id    int references users(id),
  created_at text not null 
);

create table posts (
  id         int primary key,
  body       text,
  user_id    int references users(id),
  thread_id  int references threads(id),
  created_at text not null
);
