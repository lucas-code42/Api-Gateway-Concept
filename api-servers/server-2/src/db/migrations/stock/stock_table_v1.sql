-- create stock table
CREATE TABLE public.stock(
    id SERIAL NOT NULL,
    book_id INT NOT NULL,
    quantity INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY(book_id) REFERENCES books(id)
);