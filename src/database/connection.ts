import { Pool } from 'pg';
import dotenv from 'dotenv';

dotenv.config();

// Connection pooling configuration
const pool = new Pool({
  connectionString: process.env.DATABASE_URL,
});

// Export db instance for models/controllers
export default pool;