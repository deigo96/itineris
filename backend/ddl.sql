-- public.leave_types definition

-- Drop table

-- DROP TABLE public.leave_types;

CREATE TABLE public.leave_types (
	id smallserial NOT NULL,
	type_name varchar NOT NULL,
	is_pns bool DEFAULT true NOT NULL,
	is_pppk bool DEFAULT false NOT NULL,
	CONSTRAINT leave_types_pk PRIMARY KEY (id),
	CONSTRAINT leave_types_unique UNIQUE (type_name)
);


-- public.roles definition

-- Drop table

-- DROP TABLE public.roles;

CREATE TABLE public.roles (
	id smallserial NOT NULL,
	role_name varchar NOT NULL,
	CONSTRAINT roles_pk PRIMARY KEY (id),
	CONSTRAINT roles_unique UNIQUE (role_name)
);


-- public.employees definition

-- Drop table

-- DROP TABLE public.employees;

CREATE TABLE public.employees (
	id serial4 NOT NULL,
	nip varchar NOT NULL,
	"name" varchar NOT NULL,
	"password" varchar NOT NULL,
	role_id int2 NOT NULL,
	leave_balance int2 DEFAULT 12 NOT NULL,
	is_pns bool DEFAULT true NOT NULL,
	created_at timestamptz DEFAULT now() NOT NULL,
	created_by varchar DEFAULT CURRENT_USER NOT NULL,
	updated_at timestamptz DEFAULT now() NOT NULL,
	updated_by varchar DEFAULT CURRENT_USER NOT NULL,
	CONSTRAINT newtable_pk PRIMARY KEY (id),
	CONSTRAINT employees_roles_fk FOREIGN KEY (role_id) REFERENCES public.roles(id)
);


-- public.leave_requests definition

-- Drop table

-- DROP TABLE public.leave_requests;

CREATE TABLE public.leave_requests (
	id int4 DEFAULT nextval('leave_requests_column1_seq'::regclass) NOT NULL,
	employee_id int4 NOT NULL,
	leave_type int2 NOT NULL,
	start_date date NOT NULL,
	end_date date NOT NULL,
	reason varchar NULL,
	status public."request_status" NOT NULL,
	rejection_note varchar NULL,
	created_at timestamptz DEFAULT now() NOT NULL,
	created_by varchar DEFAULT CURRENT_USER NOT NULL,
	updated_at timestamptz DEFAULT now() NOT NULL,
	updated_by varchar DEFAULT CURRENT_USER NOT NULL,
	total_days int2 DEFAULT 0 NOT NULL,
	CONSTRAINT leave_requests_pk PRIMARY KEY (id),
	CONSTRAINT leave_requests_employees_fk FOREIGN KEY (employee_id) REFERENCES public.employees(id),
	CONSTRAINT leave_requests_leave_types_fk FOREIGN KEY (leave_type) REFERENCES public.leave_types(id)
);