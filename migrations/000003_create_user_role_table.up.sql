CREATE TABLE IF NOT EXISTS user_roles(
    users_id BIGSERIAL REFERENCES users(id),
    role_id BIGSERIAL REFERENCES roles(id)
)