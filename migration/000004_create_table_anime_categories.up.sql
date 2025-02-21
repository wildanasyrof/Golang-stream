CREATE TABLE anime_categories (
    anime_id INT UNSIGNED NOT NULL,
    category_id INT UNSIGNED NOT NULL,
    PRIMARY KEY (anime_id, category_id),
    FOREIGN KEY (anime_id) REFERENCES animes(id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
);
