import app from './app';
import pool from './database/connection';

const PORT = process.env.PORT || 3000;

// Test database connection
pool.connect()
  .then(() => console.log('Database connected'))
  .catch(err => console.error('Database connection error:', err));

app.listen(PORT, () => {
  console.log(`Server running on port ${PORT}`);
});