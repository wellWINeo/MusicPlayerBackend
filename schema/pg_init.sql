-- init script to create SQL database

-- creating users table
create table Users
(
    id serial primary key,
    username varchar(30) not null,
    email varchar(100) ,
    passwd varchar(100) not null,
    is_premium boolean
);

create table Tracks
(
    id serial primary key,
    title varchar(50) not null,
    audio blob,
    hash varchar(256) not null,
    -- foreign key
    artist integer references Artist (id)
    year int,
    has_video boolean
);

create table Likes
(
    id serial primary key,
    track integer references Tracks (id),
    [user] integer references Users (id),
    [time] date
);

create table History
(
    id serial primary key,
    track integer references Tracks (id),
    [user] integer references Users (id),
    [time] date
);

create table Referals
(
    id serial primary key,
    old_user integer references Users (id),
    new_user integer references Users (id)
);

create table Artists
(
    id serial primary key,
    name varchar(100)
);

create table Playlist
(
    id serial primary key,
    [user] integer references Users (id),
    title varchar(100)
);

create table Owns
(
    id serial primary key,
    track integer references Tracks (id),
    [user] integer references Users (id)
);

create table PlaylistContent
(
    id serial primary key,
    track integer references Tracks (id),
    playlist integer references Playlist (id)
);
