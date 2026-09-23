CREATE TABLE user_statistics (
    user_id BIGINT UNSIGNED PRIMARY KEY,
    total_games_played INT UNSIGNED NOT NULL DEFAULT 0,
    guess_games_played INT UNSIGNED NOT NULL DEFAULT 0,
    ordering_games_played INT UNSIGNED NOT NULL DEFAULT 0,
    matching_games_played INT UNSIGNED NOT NULL DEFAULT 0,
    guess_total_points BIGINT UNSIGNED NOT NULL DEFAULT 0,
    ordering_total_points BIGINT UNSIGNED NOT NULL DEFAULT 0,
    matching_total_points BIGINT UNSIGNED NOT NULL DEFAULT 0,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
        ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_user_statistics_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);