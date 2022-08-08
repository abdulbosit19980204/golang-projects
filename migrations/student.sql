-- Table: public.student

-- DROP TABLE IF EXISTS public.student;

CREATE TABLE IF NOT EXISTS public.student
(
    student_id integer NOT NULL,
    first_name character varying(255) COLLATE pg_catalog."default",
    last_name character varying(255) COLLATE pg_catalog."default",
    father_name character varying(255) COLLATE pg_catalog."default",
    age integer,
    address_id integer,
    CONSTRAINT student_pkey PRIMARY KEY (student_id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.student
    OWNER to postgres;