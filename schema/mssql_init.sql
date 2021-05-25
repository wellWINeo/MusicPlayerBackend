-- init script to create SQL database

-- some configuration staff
set ansi_padding on
go

set quoted_identifier on
go

set ansi_nulls on
go

-- create database
create database MusicPlayer
go

use MusicPlayer
go


-- creating tables
create table Users
(
    id int not null identity(1,1),
    constraint pk_user_id primary key clustered (id)
    username varchar(30) not null,
    email varchar(100) ,
    passwd varchar(100) not null,
    is_premium boolean not null
);

create table Tracks
(
    id int not null identity(1,1),
    constraint pk_track_id primary key clustered (id)
    title varchar(50) not null,
    audio blob,
    hash varchar(256) not null,
    artist int,
    constraint fk_tracks_artist foreign key (artist) references Artist (id),
    year int,
    has_video boolean
);

create table Likes
(
    id int not null identity(1,1),
    constraint pk_likes_id primary key clustered (id),
    track int,
    constraint fk_likes_track foreign key (track) references Tracks (id),
    [user] int,
    constraint fk_likes_user foreign key ([user]) references Users (id),
    [time] date
);

create table History
(
    id int not null identity(1,1),
    constraint pk_history_id primary key clustered (id),
    track int,
    constraint fk_history_track foreign key (track) references Tracks (id),
    [user] int,
    constraint fk_history_user foreign key ([user]) references Users (id),
    [time] date
);

create table Referals
(
    id int not null identity(1,1),
    constraint pk_referals_id primary key clustered (id),
    old_user int,
    constraint fk_referaks_old_user foreign key ([user]) references Users (id),
    new_user int,
    constraint fk_referals_new_user foreign key ([user]) references Users (id)
);

create table Artists
(
    id int not null identity(1,1),
    constraint pk_artists_id primary key clustered (id),
    name varchar(100)
);

create table Playlist
(
    id int not null identity(1,1),
    constraint pk_artists_id primary key clustered (id),
    [user] int,
    constraint fk_playlist_user foreign key ([user]) references Users (id),
    title varchar(100)
);

create table Owns
(
    id int not null identity(1,1),
    constraint fk_owns_id primary key (id)
    track int,
    constraint fk_owns_track foreign key (track) references Tracks (id),
    [user] int,
    constraint fk_owns_user foreign key ([user]) references Users (id),
);

create table PlaylistContent
(
    id int not null identity(1,1),
    constraint pk_playlistcontent_id primary key (id),
    track integer references Tracks (id),
    constraint fk_playlist_content_track foreign key (track) references Tracks (id),
    playlist integer references Playlist (id),
    constraint fk_content_playlist references key (playlist) references Playlist (id),
);
