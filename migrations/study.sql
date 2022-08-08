-- Table: public.study

-- DROP TABLE IF EXISTS public.study;

CREATE TABLE IF NOT EXISTS public.study
(
    study_id integer NOT NULL,
    start_date date,
    end_date date,
    status_id integer,
    study_type_id integer,
    student_id numeric,
    CONSTRAINT study_pkey PRIMARY KEY (study_id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.study
    OWNER to postgres;