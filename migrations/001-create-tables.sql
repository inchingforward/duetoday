create table todo_user (
    id bigserial primary key,
    email text not null,
    password text not null,
    created_at timestamp with time zone not null,
    last_login_at timestamp with time zone
);

create table todo (
    id bigserial primary key,
    title text not null,
    details text,
    user_id bigint references todo_user(id),
    created_at timestamp with time zone not null default now(),
    completed_flag boolean not null default false,
    due_at date not null,
    rank int
);

commit;