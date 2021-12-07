create table course
(
    id          bigserial primary key,
    -- user_id     bigint not null,
    title       varchar,
    price       int,
    description varchar,
    -- foreign key (user_id) references person (id)
    --     on delete cascade
);