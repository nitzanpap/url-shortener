import axios from 'axios';
import { configurations } from "../configs/config"

const serverApi = axios.create({
    baseURL: configurations.envVars.serverBaseUrl
});

export const isServerAvailable = async () => {
    try {
        const response = await serverApi.get("/");
        return response.status === 200 && response.data.status === "ok";
    } catch (error) {
        console.error("Error checking server status:", error);
        return false;
    }
}

export const generateShortUrl = async (url: string) => {
};

// Add more API methods as needed

export default serverApi;
