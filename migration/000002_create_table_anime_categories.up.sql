-- Create animes table
CREATE TABLE animes (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    alt_titles VARCHAR(255),
    chapters VARCHAR(50),
    studio VARCHAR(255),
    year VARCHAR(20),
    rating DECIMAL(2,1),
    synopsis TEXT,
    image_source VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Create categories table
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Create pivot table for many-to-many relationship between animes and categories
CREATE TABLE anime_categories (
    anime_id INT NOT NULL,
    category_id INT NOT NULL,
    PRIMARY KEY (anime_id, category_id),
    FOREIGN KEY (anime_id) REFERENCES animes(id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
);
