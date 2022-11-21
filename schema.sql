
create table users (
  id         integer primary key autoincrement,
  name       varchar(255),
  email      varchar(255) not null unique,
  password   varchar(255) not null,
  created_at text not null 
);

create table sessions (
  id         integer primary key autoincrement,
  uuid       varchar(64) not null unique,
  email      varchar(255),
  user_id    int references users(id),
  created_at text not null 
);

create table threads (
  id         integer primary key autoincrement,
  topic      text,
  user_id    int references users(id),
  created_at text not null 
);

create table posts (
  id         integer primary key autoincrement,
  body       text,
  user_id    int references users(id),
  thread_id  int references threads(id),
  created_at text not null
);
