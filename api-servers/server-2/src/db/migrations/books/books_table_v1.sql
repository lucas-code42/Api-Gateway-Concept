CREATE TABLE public.books(
    id serial NOT NULL,
    name character varying(255) NOT NULL,
    price double precision NOT NULL,
    author character varying(64) NOT NULL,
    PRIMARY KEY (id)
);