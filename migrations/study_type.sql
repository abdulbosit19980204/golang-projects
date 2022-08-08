-- Table: public.study_type

-- DROP TABLE IF EXISTS public.study_type;

CREATE TABLE IF NOT EXISTS public.study_type
(
    stusy_type_id integer NOT NULL,
    status_title character varying(255) COLLATE pg_catalog."default",
    CONSTRAINT study_type_pkey PRIMARY KEY (stusy_type_id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.study_type
    OWNER to postgres;