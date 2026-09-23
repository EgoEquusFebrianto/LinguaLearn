ALTER TABLE users
ADD COLUMN role_id BIGINT UNSIGNED NOT NULL AFTER id,
ADD CONSTRAINT fk_users_role
    FOREIGN KEY (role_id)
    REFERENCES roles(id)
    ON DELETE RESTRICT
    ON UPDATE CASCADE;

CREATE INDEX idx_users_role_id
ON users(role_id);