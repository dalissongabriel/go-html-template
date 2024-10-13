create table products (
    id serial primary key,
    name varchar,
    description varchar,
    price decimal,
    amount integer
);


INSERT INTO products(name, description, price, amount) VALUES
('Tenis', 'Confortável', 89, 3),
('Fone', 'Muito bom', 59, 2);