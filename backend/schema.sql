--
-- PostgreSQL database dump
--

-- Dumped from database version 16.4 (Debian 16.4-1.pgdg120+2)
-- Dumped by pg_dump version 17.2

-- Started on 2025-02-27 19:38:36

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 4 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: pg_database_owner
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO pg_database_owner;

--
-- TOC entry 3409 (class 0 OID 0)
-- Dependencies: 4
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: pg_database_owner
--

COMMENT ON SCHEMA public IS 'standard public schema';


--
-- TOC entry 856 (class 1247 OID 24706)
-- Name: request_status; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.request_status AS ENUM (
    'pending',
    'approved',
    'rejected'
);


ALTER TYPE public.request_status OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 216 (class 1259 OID 24662)
-- Name: employees; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.employees (
    id integer NOT NULL,
    nip character varying NOT NULL,
    name character varying NOT NULL,
    password character varying NOT NULL,
    role_id smallint NOT NULL,
    leave_balance smallint DEFAULT 12 NOT NULL,
    is_pns boolean DEFAULT true NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    created_by character varying DEFAULT CURRENT_USER NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_by character varying DEFAULT CURRENT_USER NOT NULL
);


ALTER TABLE public.employees OWNER TO postgres;

--
-- TOC entry 215 (class 1259 OID 24661)
-- Name: employees_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.employees_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.employees_id_seq OWNER TO postgres;

--
-- TOC entry 3410 (class 0 OID 0)
-- Dependencies: 215
-- Name: employees_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.employees_id_seq OWNED BY public.employees.id;


--
-- TOC entry 222 (class 1259 OID 24714)
-- Name: leave_requests; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.leave_requests (
    id integer NOT NULL,
    employee_id integer NOT NULL,
    leave_type smallint NOT NULL,
    start_date date NOT NULL,
    end_date date NOT NULL,
    reason character varying,
    status public.request_status NOT NULL,
    rejection_note character varying,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    created_by character varying DEFAULT CURRENT_USER NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_by character varying DEFAULT CURRENT_USER NOT NULL,
    total_days smallint DEFAULT 0 NOT NULL
);


ALTER TABLE public.leave_requests OWNER TO postgres;

--
-- TOC entry 221 (class 1259 OID 24713)
-- Name: leave_requests_column1_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.leave_requests_column1_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.leave_requests_column1_seq OWNER TO postgres;

--
-- TOC entry 3411 (class 0 OID 0)
-- Dependencies: 221
-- Name: leave_requests_column1_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.leave_requests_column1_seq OWNED BY public.leave_requests.id;


--
-- TOC entry 220 (class 1259 OID 24693)
-- Name: leave_types; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.leave_types (
    id smallint NOT NULL,
    type_name character varying NOT NULL,
    is_pns boolean DEFAULT true NOT NULL,
    is_pppk boolean DEFAULT false NOT NULL
);


ALTER TABLE public.leave_types OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 24692)
-- Name: leave_types_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.leave_types_id_seq
    AS smallint
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.leave_types_id_seq OWNER TO postgres;

--
-- TOC entry 3412 (class 0 OID 0)
-- Dependencies: 219
-- Name: leave_types_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.leave_types_id_seq OWNED BY public.leave_types.id;


--
-- TOC entry 218 (class 1259 OID 24677)
-- Name: roles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.roles (
    id smallint NOT NULL,
    role_name character varying NOT NULL
);


ALTER TABLE public.roles OWNER TO postgres;

--
-- TOC entry 217 (class 1259 OID 24676)
-- Name: roles_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.roles_id_seq
    AS smallint
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.roles_id_seq OWNER TO postgres;

--
-- TOC entry 3413 (class 0 OID 0)
-- Dependencies: 217
-- Name: roles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.roles_id_seq OWNED BY public.roles.id;


--
-- TOC entry 3221 (class 2604 OID 24665)
-- Name: employees id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.employees ALTER COLUMN id SET DEFAULT nextval('public.employees_id_seq'::regclass);


--
-- TOC entry 3232 (class 2604 OID 24717)
-- Name: leave_requests id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.leave_requests ALTER COLUMN id SET DEFAULT nextval('public.leave_requests_column1_seq'::regclass);


--
-- TOC entry 3229 (class 2604 OID 24696)
-- Name: leave_types id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.leave_types ALTER COLUMN id SET DEFAULT nextval('public.leave_types_id_seq'::regclass);


--
-- TOC entry 3228 (class 2604 OID 24680)
-- Name: roles id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roles ALTER COLUMN id SET DEFAULT nextval('public.roles_id_seq'::regclass);


--
-- TOC entry 3397 (class 0 OID 24662)
-- Dependencies: 216
-- Data for Name: employees; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.employees (id, nip, name, password, role_id, leave_balance, is_pns, created_at, created_by, updated_at, updated_by) FROM stdin;
1	12345678	deigo	$2a$12$KLsGVtpw/Fxmsz2BXpz3uewpLzCu1ifzagFTA.AzyD4eysmfq2SOi	1	6	t	2025-02-27 03:21:05.274068+00	postgres	2025-02-27 03:21:05.274068+00	postgres
2	1234567	siahaan	$2a$12$KLsGVtpw/Fxmsz2BXpz3uewpLzCu1ifzagFTA.AzyD4eysmfq2SOi	2	12	t	2025-02-27 09:27:54.401362+00	postgres	2025-02-27 09:27:54.401362+00	postgres
\.


--
-- TOC entry 3403 (class 0 OID 24714)
-- Dependencies: 222
-- Data for Name: leave_requests; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.leave_requests (id, employee_id, leave_type, start_date, end_date, reason, status, rejection_note, created_at, created_by, updated_at, updated_by, total_days) FROM stdin;
2	1	1	2025-03-01	2025-02-28	sakit	pending		2025-02-27 05:42:00.470537+00		2025-02-27 05:42:00.470537+00		0
3	1	1	2025-03-01	2025-02-28	sakit	pending		2025-02-27 05:46:30.063159+00		2025-02-27 05:46:30.063159+00		0
4	1	1	2025-03-01	2025-05-28	sakit	pending		2025-02-27 05:50:51.831589+00		2025-02-27 05:50:51.831589+00		0
5	1	1	2025-03-01	2025-05-28	sakit	pending		2025-02-27 05:51:26.033503+00		2025-02-27 05:51:26.033503+00		0
6	1	1	2025-02-28	2025-03-01	sakit	pending		2025-02-27 06:15:04.119596+00		2025-02-27 06:15:04.119596+00		0
7	1	1	2025-02-28	2025-03-01	sakit	pending		2025-02-27 06:15:45.247119+00		2025-02-27 06:15:45.247119+00		0
8	1	1	2025-02-28	2025-03-01	sakit	pending		2025-02-27 06:17:20.778237+00		2025-02-27 06:17:20.778237+00		0
9	1	1	2025-02-28	2025-03-01	sakit	pending		2025-02-27 06:18:29.320166+00		2025-02-27 06:18:29.320166+00		0
11	1	1	2025-02-28	2025-03-01	sakit	rejected		2025-02-27 06:22:33.756708+00		2025-02-27 06:22:33.756708+00	12345678	0
12	2	1	2025-02-28	2025-03-01	sakit	rejected		2025-02-27 09:28:37.362878+00		2025-02-27 09:28:37.362878+00	12345678	2
10	1	1	2025-02-28	2025-03-01	sakit	approved		2025-02-27 06:19:48.80702+00		2025-02-27 06:19:48.80702+00	12345678	0
\.


--
-- TOC entry 3401 (class 0 OID 24693)
-- Dependencies: 220
-- Data for Name: leave_types; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.leave_types (id, type_name, is_pns, is_pppk) FROM stdin;
2	Cuti Besar	t	f
3	Cuti Sakit	t	t
4	Cuti Melahirkan	t	t
5	Cuti Karena Alasan Penting	t	f
6	Cuti Bersama	t	t
1	Cuti Tahunan	t	t
7	Cuti Diluar Tanggungan Negara	t	f
\.


--
-- TOC entry 3399 (class 0 OID 24677)
-- Dependencies: 218
-- Data for Name: roles; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.roles (id, role_name) FROM stdin;
1	PPK
2	STAFF
\.


--
-- TOC entry 3414 (class 0 OID 0)
-- Dependencies: 215
-- Name: employees_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.employees_id_seq', 2, true);


--
-- TOC entry 3415 (class 0 OID 0)
-- Dependencies: 221
-- Name: leave_requests_column1_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.leave_requests_column1_seq', 12, true);


--
-- TOC entry 3416 (class 0 OID 0)
-- Dependencies: 219
-- Name: leave_types_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.leave_types_id_seq', 7, true);


--
-- TOC entry 3417 (class 0 OID 0)
-- Dependencies: 217
-- Name: roles_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.roles_id_seq', 2, true);


--
-- TOC entry 3249 (class 2606 OID 24725)
-- Name: leave_requests leave_requests_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.leave_requests
    ADD CONSTRAINT leave_requests_pk PRIMARY KEY (id);


--
-- TOC entry 3245 (class 2606 OID 24702)
-- Name: leave_types leave_types_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.leave_types
    ADD CONSTRAINT leave_types_pk PRIMARY KEY (id);


--
-- TOC entry 3247 (class 2606 OID 24704)
-- Name: leave_types leave_types_unique; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.leave_types
    ADD CONSTRAINT leave_types_unique UNIQUE (type_name);


--
-- TOC entry 3239 (class 2606 OID 24675)
-- Name: employees newtable_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.employees
    ADD CONSTRAINT newtable_pk PRIMARY KEY (id);


--
-- TOC entry 3241 (class 2606 OID 24684)
-- Name: roles roles_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_pk PRIMARY KEY (id);


--
-- TOC entry 3243 (class 2606 OID 24686)
-- Name: roles roles_unique; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_unique UNIQUE (role_name);


--
-- TOC entry 3250 (class 2606 OID 24687)
-- Name: employees employees_roles_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.employees
    ADD CONSTRAINT employees_roles_fk FOREIGN KEY (role_id) REFERENCES public.roles(id);


--
-- TOC entry 3251 (class 2606 OID 24726)
-- Name: leave_requests leave_requests_employees_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.leave_requests
    ADD CONSTRAINT leave_requests_employees_fk FOREIGN KEY (employee_id) REFERENCES public.employees(id);


--
-- TOC entry 3252 (class 2606 OID 24731)
-- Name: leave_requests leave_requests_leave_types_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.leave_requests
    ADD CONSTRAINT leave_requests_leave_types_fk FOREIGN KEY (leave_type) REFERENCES public.leave_types(id);


-- Completed on 2025-02-27 19:39:00

--
-- PostgreSQL database dump complete
--

