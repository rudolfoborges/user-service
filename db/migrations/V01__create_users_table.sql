create table if not exists users (
    id uuid primary key unique,
    name text not null,
    email text not null unique,
    password text not null,
    active boolean not null default true,
    avatar_url text,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now()
);
create index users_email_idx on users(email);