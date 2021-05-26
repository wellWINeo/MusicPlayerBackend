
-- init script to create SQL database

-- some configuration staff
set ansi_padding on
go

set quoted_identifier on
go

set ansi_nulls on
go


-- create database
drop database if exists MusicPlayer;
create database MusicPlayer;
use MusicPlayer
go


create table Artists
(
    id_artist int not null identity(1,1),
    name varchar(100),
    constraint pk_artists_id primary key clustered (id_artist)
);
--go

create table Genre
(
    id_genre int not null identity(1,1),
    title varchar(100),
    constraint fk_genre_id primary key clustered (id_genre)
);
go

-- creating user table
create table Users
(
    id_user int not null identity(1,1),
    username varchar(30) not null,
    email varchar(100),
    passwd varchar(100) not null,
    is_premium bit not null,
    constraint pk_user_id primary key clustered (id_user)

);
--go

-- creating tracks table
create table Tracks
(
    id_track int not null identity(1,1),
    title varchar(50) not null,
    audio varbinary(max),
    hash varchar(256) not null,
    artist_id int,
    [year] int,
    genre_id int,
    has_video bit,
    ---
    constraint pk_track_id primary key clustered (id_track),
    constraint fk_tracks_artist foreign key (artist_id) references Artists (id_artist),
    constraint fk_tracks_genre foreign key (genre_id) references Genre (id_genre)
);
--go

create table Likes
(
    id_likes int not null identity(1,1),
    track_id int,
    [user_id] int,
    [time] date,
    ---
    constraint pk_likes_id primary key clustered (id_likes),
    constraint fk_likes_track foreign key (track_id) references Tracks (id_track),
    constraint fk_likes_user foreign key ([user_id]) references Users (id_user)
);
--go

create table History
(
    id_history int not null identity(1,1),
    track_id int,
    [user_id] int,
    [time] date,
    ---
    constraint pk_history_id primary key clustered (id_history),
    constraint fk_history_track foreign key (track_id) references Tracks (id_track),
    constraint fk_history_user foreign key ([user_id]) references Users (id_user)
);
--go

create table Referals
(
    id_referal int not null identity(1,1),
    old_user_id int,
    new_user_id int,
    ---
    constraint pk_referals_id primary key clustered (id_referal),
    constraint fk_referaks_old_user foreign key (old_user_id) references Users (id_user),
    constraint fk_referals_new_user foreign key (new_user_Id) references Users (id_user)
);
--go

create table Playlist
(
    id_playlist int not null identity(1,1),
    [user_id] int,
    title varchar(100),
    ---
    constraint pk_playlist_id primary key clustered (id_playlist),
    constraint fk_playlist_user foreign key ([user_id]) references Users (id_user)
);
--go

create table Owns
(
    id_own int not null identity(1,1),
    track_id int,
    [user_id] int,
    ---
    constraint fk_owns_id primary key (id_own),
    constraint fk_owns_track foreign key (track_id) references Tracks (id_track),
    constraint fk_owns_user foreign key ([user_id]) references Users (id_user)
);
--go

create table PlaylistContent
(
    id_content int not null identity(1,1),
    track_id integer,
    playlist_id integer,
    ---
    constraint pk_playlistcontent_id primary key (id_content),
    constraint fk_playlist_content_track foreign key (track_id) references Tracks (id_track),
    constraint fk_content_playlist foreign key (playlist_id) references Playlist (id_playlist)
);
--go
