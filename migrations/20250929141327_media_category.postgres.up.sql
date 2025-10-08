CREATE TABLE media_categories (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

INSERT INTO media_categories (name) VALUES
('Random'),
('Ramadhan Brawijaya'),
('UAKI School 1'),
('Flourish'),
('Sama Kahfi'),
('Mentoring'),
('Ngabarkuy');
