
-- sql script to create database

create table users
(
    id integer primary key autoincrement,
    login text unique,
    email text unique,
    passwd text,
    is_premium boolean
);

create table tracks
(
    id integer primary key autoincrement,
    [name] text not null,
    artist text not null,
    [path] text
);

create table history
(
    id integer primary key autoincrement,
    stamp date not null,
    track_id integer not null,
    foreign key (track_id) references tracks (id)
);

create table liked
(
    id integer primary key autoincrement,
    track_id integer not null,
    foreign key (track_id) references tracks (id)
);

create table tokens
(
    id integer primary key autoincrement,
    token text not null,
    user_id int not null,
    foreign key (user_id) references users (id)
);
