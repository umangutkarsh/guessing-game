import axios from 'axios';

const api = axios.create({
	baseURL: `https://guessing-game-server-3yoxob47z-umangutkarshs-projects.vercel.app`,
});

export default api;
