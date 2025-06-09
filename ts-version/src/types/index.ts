// Enums
export type PortType = 'ocean' | 'rail';
export type DirectionType = 'import' | 'export';
export type VolumeTrend = 'increasing' | 'decreasing' | 'stable';
export type Period = 'monthly' | 'quarterly' | 'yearly';

// Base interfaces (database entities)
export interface Company {
  id: number;
  name: string;
  domain: string;
  cg_code?: string;
  lead?: string;
  note?: string;
  revenue?: number;
  locations?: string[];
  employee_count?: number;
  created_at: Date;
  updated_at: Date;
}

export interface Port {
  id: number;
  name: string;
  kind: PortType;
}

export interface TradeData {
  id: number;
  company_id: number;
  port_id: number;
  direction: DirectionType;
  volume: number;
  period: Period;
  top_commodities?: string[];
  trend: VolumeTrend;
  created_at: Date;
  updated_at: Date;
}

export interface Contact {
  id: number;
  company_id: number;
  first_name: string;
  last_name: string;
  email?: string;
  phone?: string;
  title?: string;
  created_at: Date;
}

export interface TimelineEntry {
  id: number;
  company_id: number;
  note?: string;
  reminder?: Date;
  created_at: Date;
}

// Input types (for creating/updating - without auto-generated fields)
export interface CreateCompanyInput {
  name: string;
  domain: string;
  cg_code?: string;
  lead?: string;
  note?: string;
  revenue?: number;
  locations?: string[];
  employee_count?: number;
}

export interface UpdateCompanyInput {
  name?: string;
  domain?: string;
  cg_code?: string;
  lead?: string;
  note?: string;
  revenue?: number;
  locations?: string[];
  employee_count?: number;
}

export interface CreatePortInput {
  name: string;
  kind: PortType;
}

export interface CreateTradeDataInput {
  company_id: number;
  port_id: number;
  direction: DirectionType;
  volume: number;
  period: Period;
  top_commodities?: string[];
  trend: VolumeTrend;
}

export interface CreateContactInput {
  company_id: number;
  first_name: string;
  last_name: string;
  email?: string;
  phone?: string;
  title?: string;
}

export interface CreateTimelineEntryInput {
  company_id: number;
  note?: string;
  reminder?: Date;
}

// Response types (for API responses with joined data)
export interface CompanyWithRelations extends Company {
  contacts?: Contact[];
  trade_data?: TradeDataWithPort[];
  timeline_entries?: TimelineEntry[];
}

export interface TradeDataWithCompanyAndPort extends TradeData {
  port?: Port;
  company?: Company;
}

// API standard response wrappers
export interface ApiResponse<T> {
  success: boolean;
  data: T;
  message?: string;
}

export interface PaginatedResponse<T> {
  success: boolean;
  data: T[];
  pagination: {
    page: number;
    limit: number;
    total: number;
    totalPages: number;
  };
}

// Query parameter types
// TODO modify/extend
export interface CompanyQueryParams {
  page?: number;
  limit?: number;
  search?: string;
  lead?: string;
  revenue_min?: number;
  revenue_max?: number;
}

export interface TradeDataQueryParams {
  page?: number;
  limit?: number;
  company_id?: number;
  port_id?: number;
  direction?: DirectionType;
  period?: Period;
  trend?: VolumeTrend;
}