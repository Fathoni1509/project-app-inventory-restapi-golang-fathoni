--
-- PostgreSQL database dump
--

\restrict ALo0siRcg4YOlfog53rDgGIGlQYvRWQIiRpAej3YImS9sOdA4SbbpQEIPaGsMgi

-- Dumped from database version 18.1
-- Dumped by pg_dump version 18.1

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: categories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.categories (
    category_id integer NOT NULL,
    name character varying(100) NOT NULL,
    description character varying(256) NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE public.categories OWNER TO postgres;

--
-- Name: categories_category_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.categories_category_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.categories_category_id_seq OWNER TO postgres;

--
-- Name: categories_category_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.categories_category_id_seq OWNED BY public.categories.category_id;


--
-- Name: products; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.products (
    product_id integer NOT NULL,
    name character varying(100) NOT NULL,
    category_id integer NOT NULL,
    purchase_price numeric,
    sell_price numeric,
    updated_at timestamp with time zone DEFAULT now() CONSTRAINT products_last_updated_not_null NOT NULL,
    updated_by integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    shelve_id integer,
    quantity integer DEFAULT 0
);


ALTER TABLE public.products OWNER TO postgres;

--
-- Name: products_product_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.products_product_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.products_product_id_seq OWNER TO postgres;

--
-- Name: products_product_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.products_product_id_seq OWNED BY public.products.product_id;


--
-- Name: sales; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sales (
    sale_id integer NOT NULL,
    user_id integer NOT NULL,
    product_id integer NOT NULL,
    items integer NOT NULL,
    price numeric,
    total numeric,
    created_at timestamp with time zone CONSTRAINT sales_sale_date_not_null NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE public.sales OWNER TO postgres;

--
-- Name: sales_sale_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sales_sale_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.sales_sale_id_seq OWNER TO postgres;

--
-- Name: sales_sale_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sales_sale_id_seq OWNED BY public.sales.sale_id;


--
-- Name: shelves; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.shelves (
    shelve_id integer NOT NULL,
    warehouse_id integer NOT NULL,
    name character varying(100) NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE public.shelves OWNER TO postgres;

--
-- Name: shelves_shelve_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.shelves_shelve_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.shelves_shelve_id_seq OWNER TO postgres;

--
-- Name: shelves_shelve_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.shelves_shelve_id_seq OWNED BY public.shelves.shelve_id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    user_id integer NOT NULL,
    username character varying(100) NOT NULL,
    email character varying(100) NOT NULL,
    password character varying(255) NOT NULL,
    role character varying(20) NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_user_id_seq OWNER TO postgres;

--
-- Name: users_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_user_id_seq OWNED BY public.users.user_id;


--
-- Name: warehouses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.warehouses (
    warehouse_id integer NOT NULL,
    name character varying(100) NOT NULL,
    location character varying(100) NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE public.warehouses OWNER TO postgres;

--
-- Name: warehouses_warehouse_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.warehouses_warehouse_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.warehouses_warehouse_id_seq OWNER TO postgres;

--
-- Name: warehouses_warehouse_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.warehouses_warehouse_id_seq OWNED BY public.warehouses.warehouse_id;


--
-- Name: categories category_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories ALTER COLUMN category_id SET DEFAULT nextval('public.categories_category_id_seq'::regclass);


--
-- Name: products product_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products ALTER COLUMN product_id SET DEFAULT nextval('public.products_product_id_seq'::regclass);


--
-- Name: sales sale_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sales ALTER COLUMN sale_id SET DEFAULT nextval('public.sales_sale_id_seq'::regclass);


--
-- Name: shelves shelve_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shelves ALTER COLUMN shelve_id SET DEFAULT nextval('public.shelves_shelve_id_seq'::regclass);


--
-- Name: users user_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN user_id SET DEFAULT nextval('public.users_user_id_seq'::regclass);


--
-- Name: warehouses warehouse_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.warehouses ALTER COLUMN warehouse_id SET DEFAULT nextval('public.warehouses_warehouse_id_seq'::regclass);


--
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.categories (category_id, name, description, created_at, updated_at, deleted_at) FROM stdin;
1	kategori 1 baru	ini deskripsi kategori 1 baru	2026-01-04 13:30:44.166008+07	2026-01-04 13:53:38.357562+07	2026-01-04 14:46:34.636215+07
2	kategori 2	ini deskripsi kategori 2	2026-01-04 15:03:27.836123+07	2026-01-04 15:03:27.836123+07	\N
3	kategori 3	ini deskripsi kategori 3	2026-01-04 15:04:47.058851+07	2026-01-04 15:04:47.058851+07	\N
4	kategori 4 baru	ini deskripsi kategori 4	2026-01-04 18:56:02.923454+07	2026-01-04 18:56:43.337609+07	\N
5	kategori 5 baru	ini deskripsi kategori 5	2026-01-07 11:10:18.231226+07	2026-01-07 11:10:35.984586+07	\N
6	kategori 6 baru	ini deskripsi kategori 6	2026-01-07 12:46:31.734825+07	2026-01-07 12:46:49.00499+07	2026-01-07 12:47:04.73471+07
\.


--
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.products (product_id, name, category_id, purchase_price, sell_price, updated_at, updated_by, created_at, deleted_at, shelve_id, quantity) FROM stdin;
4	produk 4	4	10000	15000	2026-01-06 10:49:58.309054+07	4	2026-01-06 10:49:05.435094+07	\N	2	10
3	produk 3	4	28000	35000	2026-01-07 00:11:03.006763+07	4	2026-01-05 00:11:43.257113+07	\N	2	350
2	produk 2 baru	3	27000	35000	2026-01-05 00:04:21.55556+07	1	2026-01-04 23:53:13.123071+07	2026-01-05 00:05:11.049125+07	2	150
1	produk 1 baru sekali	2	28000	35000	2026-01-07 00:24:48.255007+07	4	2026-01-04 23:46:55.384047+07	\N	2	63
5	produk 5	4	10000	15000	2026-01-07 00:58:02.457521+07	2	2026-01-07 00:58:02.457521+07	\N	2	4
6	produk 6	4	10000	15000	2026-01-07 12:53:39.978895+07	2	2026-01-07 12:53:39.978895+07	\N	2	4
\.


--
-- Data for Name: sales; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sales (sale_id, user_id, product_id, items, price, total, created_at, updated_at, deleted_at) FROM stdin;
1	5	1	100	35000	3500000	2026-01-06 22:42:14.786161+07	2026-01-06 23:28:34.086388+07	2026-01-06 23:30:57.23332+07
2	4	3	20	35000	700000	2026-01-06 23:39:13.694783+07	2026-01-06 23:39:13.694783+07	\N
3	5	3	50	35000	1750000	2026-01-06 23:39:22.186243+07	2026-01-06 23:39:22.186243+07	\N
4	5	1	25	35000	875000	2026-01-07 00:24:21.327079+07	2026-01-07 00:24:21.327079+07	\N
5	3	1	12	35000	420000	2026-01-07 00:24:48.255007+07	2026-01-07 00:24:48.255007+07	\N
\.


--
-- Data for Name: shelves; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.shelves (shelve_id, warehouse_id, name, created_at, updated_at, deleted_at) FROM stdin;
2	2	rak 2 baru	2026-01-04 20:49:45.238626+07	2026-01-04 20:51:18.395107+07	\N
1	2	rak 1	2026-01-04 19:17:47.903984+07	2026-01-04 19:17:47.903984+07	2026-01-04 20:52:12.855886+07
3	2	rak 3 baru	2026-01-07 12:49:49.793287+07	2026-01-07 12:50:00.526074+07	2026-01-07 12:50:25.771174+07
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (user_id, username, email, password, role, created_at, updated_at, deleted_at) FROM stdin;
2	admin_andi	adminandi@sistem.com	andi123	ADMIN	2025-12-31 13:18:30.764772+07	2026-01-04 22:19:46.143139+07	\N
3	admin_cika	admincika@sistem.com	cika123	ADMIN	2025-12-31 13:18:30.764772+07	2026-01-04 22:19:58.220531+07	\N
4	staff_budi	budi@sistem.com	budi123	STAFF	2025-12-31 13:18:30.764772+07	2026-01-04 22:20:05.540007+07	\N
5	staff_siti	siti@sistem.com	siti123	STAFF	2025-12-31 13:18:30.764772+07	2026-01-04 22:20:13.119974+07	\N
1	super_admin	owner@sistem.com	super123admin	SUPER_ADMIN	2025-12-31 13:18:30.764772+07	2026-01-04 22:26:35.920733+07	\N
6	farah	farah@gmail.com	farah123	STAFF	2026-01-04 22:25:56.87552+07	2026-01-04 22:25:56.87552+07	2026-01-04 22:27:01.683253+07
7	lumos	lumos@gmail.com	super123admin	STAFF	2026-01-07 12:51:19.731767+07	2026-01-07 12:51:45.786276+07	2026-01-07 12:52:11.835838+07
\.


--
-- Data for Name: warehouses; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.warehouses (warehouse_id, name, location, created_at, updated_at, deleted_at) FROM stdin;
1		jl.nin aja dulu, Jakarta Selatan	2026-01-04 17:00:39.208926+07	2026-01-04 17:57:57.6843+07	2026-01-04 17:03:58.487876+07
2	gudang 2 baru	jl.nin aja dulu, Jakarta Pusat	2026-01-04 18:20:10.411535+07	2026-01-04 20:50:55.298605+07	\N
3	gudang 3	jl.jalan malam	2026-01-07 12:48:12.925169+07	2026-01-07 12:48:32.43495+07	2026-01-07 12:48:49.40681+07
\.


--
-- Name: categories_category_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.categories_category_id_seq', 6, true);


--
-- Name: products_product_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.products_product_id_seq', 6, true);


--
-- Name: sales_sale_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sales_sale_id_seq', 5, true);


--
-- Name: shelves_shelve_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.shelves_shelve_id_seq', 3, true);


--
-- Name: users_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_user_id_seq', 7, true);


--
-- Name: warehouses_warehouse_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.warehouses_warehouse_id_seq', 3, true);


--
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (category_id);


--
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (product_id);


--
-- Name: sales sales_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sales
    ADD CONSTRAINT sales_pkey PRIMARY KEY (sale_id);


--
-- Name: shelves shelves_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shelves
    ADD CONSTRAINT shelves_pkey PRIMARY KEY (shelve_id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (user_id);


--
-- Name: warehouses warehouses_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.warehouses
    ADD CONSTRAINT warehouses_pkey PRIMARY KEY (warehouse_id);


--
-- Name: products fk_category; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES public.categories(category_id);


--
-- Name: sales fk_produt; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sales
    ADD CONSTRAINT fk_produt FOREIGN KEY (product_id) REFERENCES public.products(product_id);


--
-- Name: products fk_shelve; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT fk_shelve FOREIGN KEY (shelve_id) REFERENCES public.shelves(shelve_id);


--
-- Name: products fk_updated_by; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT fk_updated_by FOREIGN KEY (updated_by) REFERENCES public.users(user_id);


--
-- Name: sales fk_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sales
    ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES public.users(user_id);


--
-- Name: shelves fk_warehouse; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shelves
    ADD CONSTRAINT fk_warehouse FOREIGN KEY (warehouse_id) REFERENCES public.warehouses(warehouse_id);


--
-- PostgreSQL database dump complete
--

\unrestrict ALo0siRcg4YOlfog53rDgGIGlQYvRWQIiRpAej3YImS9sOdA4SbbpQEIPaGsMgi

