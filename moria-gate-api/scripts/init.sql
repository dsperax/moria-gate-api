CREATE TABLE clients (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    document VARCHAR(14) UNIQUE NOT NULL
);

CREATE TABLE addresses (
    id SERIAL PRIMARY KEY,
    document VARCHAR(14) NOT NULL,
    street VARCHAR(100),
    city VARCHAR(50),
    state VARCHAR(50),
    district VARCHAR(50),
    zip VARCHAR(10),
    country VARCHAR(50),
    FOREIGN KEY (document) REFERENCES clients(document)
);

CREATE TABLE credit_status (
    id SERIAL PRIMARY KEY,
    document VARCHAR(14) NOT NULL,
    situation VARCHAR(20),
    risk VARCHAR(10),
    type VARCHAR(2),
    FOREIGN KEY (document) REFERENCES clients(document)
);

INSERT INTO clients (name, document) VALUES
('Frodo Baggins', '12345678901'),
('Aragorn Elessar', '23456789012'),
('Gandalf the Grey', '34567890123'),
('Banana SA', '87654321'),
('Mamao LTDA', '76543210'),
('Morango ME', '65432109');

INSERT INTO addresses (document, zip, street, city, state, district, country) VALUES
('12345678901', '01001-000', 'Rua dos Hobbits', 'Shire', 'Eriador', 'Bolseiro', 'Middle Earth'),
('23456789012', '02002-000', 'Rua da Coroa', 'Gondor', 'Gondor', 'Minas Tirith', 'Middle Earth'),
('34567890123', '03003-000', 'Rua dos Magos', 'Valfenda', 'Lindon', 'Elfica', 'Middle Earth'),
('87654321', '04004-000', 'Av. das Frutas', 'São Paulo', 'SP', 'Centro', 'Brasil'),
('76543210', '05005-000', 'Rua Tropicais', 'Rio de Janeiro', 'RJ', 'Zona Sul', 'Brasil'),
('65432109', '06006-000', 'Travessa da Serra', 'Curitiba', 'PR', 'Batel', 'Brasil');

INSERT INTO credit_status (document, situation, risk, type) VALUES
('12345678901', 'limpo', 'baixo', 'PF'),
('23456789012', 'ja deveu', 'medio', 'PF'),
('34567890123', 'devedor', 'alto', 'PF'),
('87654321', 'limpo', 'baixo', 'PJ'),
('76543210', 'ja deveu', 'medio', 'PJ'),
('65432109', 'devedor', 'alto', 'PJ');
