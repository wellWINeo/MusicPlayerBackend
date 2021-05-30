
-- init script to create SQL database

-- some configuration staff
set ansi_padding on
go

set quoted_identifier on
go

set ansi_nulls on
go


--
-- DATABASE
--
drop database if exists MusicPlayer;
create database MusicPlayer;
use MusicPlayer
go

--
-- TABLES
--
create table Artists
(
    id_artist int not null identity(1,1),
    [name] varchar(100) unique,
    constraint pk_artists_id primary key clustered (id_artist)
);
go

create table Genre
(
    id_genre int not null identity(1,1),
    title varchar(100) unique,
    constraint fk_genre_id primary key clustered (id_genre)
);
go

create table TrackData
(
    id_track_data int not null identity(1,1),
    hash varchar(256) unique,
    [data] varbinary(max)
    constraint pk_track_data_id primary key clustered (id_track_data)
);

create table Users
(
    id_user int not null identity(1,1),
    username varchar(30) not null,
    email varchar(100),
    passwd varchar(100) not null,
    is_premium bit default 0,
    is_verified bit default 0,
    constraint pk_user_id primary key clustered (id_user)
);
go

-- creating tracks table
create table Tracks
(
    id_track int not null identity(1,1),
    title varchar(50) not null,
    [data] int,
    artist_id int,
    [year] int,
    genre_id int,
    has_video bit,
    owner_id int,
    ---
    constraint pk_id_track_ primary key clustered (id_track),
    constraint fk_track_id foreign key ([data]) references TrackData (id_track_data),
    constraint fk_tracks_artist foreign key (artist_id) references Artists (id_artist),
    constraint fk_tracks_genre foreign key (genre_id) references Genre (id_genre),
    constraint fk_tracks_owner foreign key (owner_id) references Users (id_user)
);
go

create table Likes
(
    id_likes int not null identity(1,1),
    track_id int,
    [user_id] int,
    [time] date,
    ---
    constraint pk_likes_id primary key clustered (id_likes),
    constraint fk_likes_track foreign key (track_id) references Tracks (id_track),
    constraint fk_likes_user foreign key ([user_id]) references Users (id_user) on delete cascade
);
go

create table History
(
    id_history int not null identity(1,1),
    track_id int,
    [user_id] int,
    [time] date,
    ---
    constraint pk_history_id primary key clustered (id_history),
    constraint fk_history_track foreign key (track_id) references Tracks (id_track),
    constraint fk_history_user foreign key ([user_id]) references Users (id_user) on delete cascade
);
go

create table Referals
(
    id_referal int not null identity(1,1),
    old_user_id int,
    new_user_id int,
    ---
    constraint pk_referals_id primary key clustered (id_referal),
    constraint fk_referals_old_user foreign key (old_user_id) references Users (id_user),
    constraint fk_referals_new_user foreign key (new_user_Id) references Users (id_user)
);
go

create table Playlist
(
    id_playlist int not null identity(1,1),
    [user_id] int,
    title varchar(100),
    artist_id integer,
    ---
    constraint pk_playlist_id primary key clustered (id_playlist),
    constraint fk_playlist_user foreign key ([user_id]) references Users (id_user) on delete cascade,
    constraint fk_playlist_artist foreign key (artist_id) references Artists (id_artist) on delete cascade,
);
go

create table PlaylistContent
(
    id_content int not null identity(1,1),
    track_id integer,
    playlist_id integer,
    ---
    constraint pk_playlistcontent_id primary key (id_content),
    constraint fk_playlist_content_track foreign key (track_id) references Tracks (id_track),
    constraint fk_content_playlist foreign key (playlist_id) references Playlist (id_playlist),
);
go

--
-- PROCEDURES
--

-- procedure to deal with genres
create procedure UpdateGenre
       @genre_name varchar(100),
       @genre_id int out
as
begin
    declare @table_id table (id int)

    begin try
          insert into Genre
          output INSERTED.id_genre into @table_id
          values(@genre_name)
          select @genre_id = id from @table_id
    end try
    begin catch
          select @genre_id = id_genre
          from Genre
          where title=@genre_name
    end catch
end;

create procedure UpdateArtist
       @artist_name varchar(100),
       @artist_id int out
as
begin
    declare @table_id table (id int)

    begin try
          insert into Artists
          output INSERTED.id_artist into @table_id
          values(@artist_name)
          select @artist_id = id from @table_id
    end try
    begin catch
          select @artist_id = id_artist
          from Artists
          where [name]=@artist_name
    end catch
end;

-- procedure to add new track
create procedure AddTrack
    @track_title varchar(50),
    @artist_name varchar(100),
    @genre_name varchar(100),
    @track_year int,
    @track_has_video bit,
    @owner_id int
as
begin
    declare @genre_id int, @artist_id int;

    exec UpdateArtist @artist_name, @artist_id out
    exec UpdateGenre @genre_name, @genre_id out


    insert into Tracks(title, artist_id, [year],
                       genre_id, has_video, owner_id)
    output INSERTED.id_track
    values(@track_title, @artist_id, @track_year,
           @genre_id, @track_has_video, @owner_id)
end;

-- procedure to update track and linked data
create procedure UpdateTrack
    @track_id int,
    @track_title varchar(50),
    @artist_name varchar(100),
    @genre_name varchar(100),
    @track_year int,
    @track_has_video bit
as
begin
    declare @artist_id int, @genre_id int

    -- exec @artist_id = UpdateArtist @artist_name
    -- exec @genre_id = UpdateGenre @genre_name
    exec UpdateArtist @artist_name, @artist_id out
    exec UpdateGenre @genre_name, @genre_id out

    update Tracks
    set title=@track_title, [year]=@track_year, has_video=@track_has_video
    where id_track=@track_id

end;

--
-- TRIGGERS
--

create trigger TrackDeleteTrigger
on Tracks
after update, delete
as
begin
    -- purging artists
    delete from Artists
    where not exists (select * from Tracks where artist_id=id_artist)
          and not exists (select * from Playlist where artist_id=id_artist)

    -- purging genres
    delete from Genre
    where not exists (select * from Tracks where genre_id=id_genre)
end;
