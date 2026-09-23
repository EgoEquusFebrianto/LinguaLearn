INSERT INTO roles (name)
VALUES
    ('user'),
    ('admin')
ON DUPLICATE KEY UPDATE
    name = VALUES(name);