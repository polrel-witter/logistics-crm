-- Base types
CREATE TYPE port_type AS ENUM ('ocean', 'rail');
CREATE TYPE direction_type AS ENUM ('import', 'export');
CREATE TYPE volume_trend AS ENUM ('increasing', 'decreasing', 'stable');
CREATE TYPE period AS ENUM ('monthly', 'quarterly', 'yearly');

-- Companies table
CREATE TABLE companies (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    domain VARCHAR(255) UNIQUE NOT NULL, -- Key for email association
    cg_code VARCHAR(50),
    lead VARCHAR(50), -- User who owns the account
    note TEXT,
    revenue INTEGER,
    locations TEXT[],
    employee_count INTEGER,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Ports table
CREATE TABLE ports (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    kind port_type NOT NULL
);

-- Trade data table
CREATE TABLE trade_data (
    id SERIAL PRIMARY KEY,
    company_id INTEGER REFERENCES companies(id),
    port_id INTEGER REFERENCES ports(id),
    direction direction_type NOT NULL,
    volume INTEGER NOT NULL,
    period period NOT NULL,
    top_commodities TEXT[],
    trend volume_trend NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Contacts table
CREATE TABLE contacts (
    id SERIAL PRIMARY KEY,
    company_id INTEGER REFERENCES companies(id),
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    email VARCHAR(255),
    phone VARCHAR(50),
    title VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW()
);

-- Timeline entries
CREATE TABLE timeline_entries (
    id SERIAL PRIMARY KEY,
    company_id INTEGER REFERENCES companies(id),
    note TEXT,
    reminder TIMESTAMP, 
    created_at TIMESTAMP DEFAULT NOW()
);