import express from 'express';
import cors from 'cors';
import helmet from 'helmet';
import routes from './routes';

const app = express();

// Middleware setup
app.use(helmet());
app.use(cors());
app.use(express.json());

// All routes under /
app.use('/api', routes);

export default app;