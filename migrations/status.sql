-- Table: public.status

-- DROP TABLE IF EXISTS public.status;

CREATE TABLE IF NOT EXISTS public.status
(
    status_id integer NOT NULL,
    status_title character varying(255) COLLATE pg_catalog."default",
    CONSTRAINT status_pkey PRIMARY KEY (status_id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.status
    OWNER to postgres;