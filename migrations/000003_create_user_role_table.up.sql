CREATE TABLE IF NOT EXISTS user_role(
    user_id BIGSERIAL REFERENCES users(id),
    role_id BIGSERIAL REFERENCES roles(id)
)