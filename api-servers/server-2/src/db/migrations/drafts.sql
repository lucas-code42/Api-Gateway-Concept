-- Database: books
-- DROP DATABASE IF EXISTS books;
CREATE DATABASE books WITH OWNER = root ENCODING = 'UTF8' LC_COLLATE = 'en_US.utf8' LC_CTYPE = 'en_US.utf8' TABLESPACE = pg_default CONNECTION
LIMIT
    = -1 IS_TEMPLATE = False;

-- create books table
CREATE TABLE public.books (
    id serial NOT NULL,
    name character varying(255) NOT NULL,
    price double precision NOT NULL,
    author character varying(64) NOT NULL,
    PRIMARY KEY (id)
);

-- create stock table
CREATE TABLE public.stock (
    id SERIAL NOT NULL,
    book_id INT NOT NULL,
    quantity INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY(book_id) REFERENCES books(id)
);

-- select all
SELECT
    *
FROM
    books;

SELECT
    *
FROM
    stock;

-- insert... id = 1
INSERT INTO
    public.books(name, price, author)
VALUES
    (
        'Redes de Computadores',
        208.55,
        'Andrew Tanenbaum'
    );

-- insert... id = 2
INSERT INTO
    public.books(name, price, author)
VALUES
    (
        'Sistemas Operacionais Modernos',
        257.00,
        'Andrew Tanenbaum'
    );

-- insert into stock (books id = 1)
INSERT INTO
    public.stock(book_id, quantity)
VALUES
    (1, 10);

-- insert into stock (books id = 2)
INSERT INTO
    public.stock(book_id, quantity)
VALUES
    (2, 168);

-- inner join example
SELECT
    *
from
    books
    inner join stock ON stock.book_id = books.id
where
    books.id = 1;