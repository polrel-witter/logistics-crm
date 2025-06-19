-- Companies table
CREATE TABLE companies (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    domain TEXT UNIQUE NOT NULL, -- Key for email association
    cg_code TEXT,
    note TEXT,
    industry TEXT,
    revenue INTEGER,
    locations TEXT, -- JSON array as TEXT
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Ports table
CREATE TABLE ports (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT UNIQUE NOT NULL,
    kind TEXT NOT NULL CHECK (kind IN ('ocean', 'rail'))
);

-- Trade data table
CREATE TABLE trade_data (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    company_id INTEGER REFERENCES companies(id),
    port_id INTEGER REFERENCES ports(id),
    direction TEXT NOT NULL CHECK (direction IN ('import', 'export')),
    volume INTEGER NOT NULL,
    period TEXT NOT NULL CHECK (period IN ('monthly', 'quarterly', 'yearly')),
    top_commodities TEXT, -- JSON array as TEXT
    trend TEXT NOT NULL CHECK (trend IN ('increasing', 'decreasing', 'stable')),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Contacts table
CREATE TABLE contacts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    company_id INTEGER REFERENCES companies(id),
    first_name TEXT,
    last_name TEXT,
    email TEXT,
    phone TEXT,
    title TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Timeline entries
CREATE TABLE timeline_entries (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    company_id INTEGER REFERENCES companies(id),
    note TEXT,
    reminder DATETIME,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Indexes for better performance
CREATE INDEX idx_companies_domain ON companies(domain);
CREATE INDEX idx_trade_data_company_id ON trade_data(company_id);
CREATE INDEX idx_trade_data_port_id ON trade_data(port_id);
CREATE INDEX idx_contacts_company_id ON contacts(company_id);
CREATE INDEX idx_timeline_entries_company_id ON timeline_entries(company_id);