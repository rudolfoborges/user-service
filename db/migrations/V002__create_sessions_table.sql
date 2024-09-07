create table if not exists sessions (
    id bigserial primary key,
    user_id uuid not null references users(id),
    created_at timestamptz not null default now()
);
