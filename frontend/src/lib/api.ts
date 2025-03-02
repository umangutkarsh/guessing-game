import axios from 'axios';

const api = axios.create({
	baseURL: `https://guessing-game-server-3yoxob47z-umangutkarshs-projects.vercel.app/api/v1`,
});

export default api;
