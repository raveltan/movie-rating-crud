create table user
(
    user_id    varchar(40) not null
        primary key,
    password   varchar(60) not null,
    first_name varchar(40) not null,
    last_name  varchar(50) not null
);

create table movies
(
    movie_id varchar(40)                   not null
        primary key,
    name     varchar(50)                   not null,
    added_on timestamp     default 'now()' not null,
    added_by varchar(40)                   not null,
    rating   decimal(5, 2) default 0.00    not null,
    voters   int           default 0       not null,
    constraint movies_user_user_id_fk
        foreign key (added_by) references user (user_id)
            on update cascade
);

create table review
(
    review_id varchar(40)               not null
        primary key,
    author    varchar(40)               not null,
    comment   varchar(200)              not null,
    rate      int                       not null,
    added_on  timestamp default 'now()' not null,
    movie_id  varchar(40)               not null,
    constraint review_movies_movie_id_fk
        foreign key (movie_id) references movies (movie_id)
            on update cascade,
    constraint review_user_user_id_fk
        foreign key (author) references user (user_id)
            on update cascade
);