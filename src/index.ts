import app from './app';
import pool from './database/connection';

const PORT = process.env.PORT || 3000;

class Server {
  private server: any;

  async start() {
    try {
      await this.connectDatabase();
      await this.startHttpServer();
      this.setupGracefulShutdown();
    } catch (error) {
      console.error('Failed to start server:', error);
      await this.cleanup();
      process.exit(1);
    }
  }

  private async connectDatabase() {
    console.log('Connecting to database...');
    const client = await pool.connect();
    await client.query('SELECT NOW()');
    client.release();
    console.log('Database connected successfully!');
  }

  private async startHttpServer() {
    return new Promise<void>((resolve) => {
      this.server = app.listen(PORT, () => {
        console.log(`Server running on http://localhost:${PORT}`);
        resolve();
      });
    });
  }

  private setupGracefulShutdown() {
    process.on('SIGTERM', () => this.shutdown());
    process.on('SIGINT', () => this.shutdown());
  }

  private async shutdown() {
    console.log('Shutting down gracefully...');
    
    if (this.server) {
      this.server.close(async () => {
        console.log('HTTP server closed');
        await this.cleanup();
        process.exit(0);
      });
    }

    // Force shutdown after 10 seconds
    setTimeout(() => {
      console.error('Could not close connections in time, forcefully shutting down');
      process.exit(1);
    }, 10000);
  }

  private async cleanup() {
    try {
      await pool.end();
      console.log('Database connections closed');
    } catch (error) {
      console.error('Error closing database connections:', error);
    }
  }
}

const server = new Server();
server.start();