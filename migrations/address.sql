-- Table: public.address

-- DROP TABLE IF EXISTS public.address;

CREATE TABLE IF NOT EXISTS public.address
(
    address_id integer NOT NULL,
    country character varying(255) COLLATE pg_catalog."default",
    city character varying(255) COLLATE pg_catalog."default",
    zip_code character(12) COLLATE pg_catalog."default",
    home_address text COLLATE pg_catalog."default",
    CONSTRAINT address_pkey PRIMARY KEY (address_id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.address
    OWNER to postgres;