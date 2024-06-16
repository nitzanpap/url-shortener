import axios from "axios"
import { configurations } from "../configs/config"

const apiVersion = "v1"
const serverApi = axios.create({
  baseURL: `${configurations.envVars.serverBaseUrl}api/${apiVersion}/`,
})

const urlsEndpoint = "urls/"

export const isServerAvailable = async () => {
  try {
    const response = await serverApi.get("/")
    return response.status === 200 && response.data.status === "ok"
  } catch (error) {
    console.error("Error checking server status:", error)
    return false
  }
}

export const generateShortUrl = async (url: string) => {
  try {
    const response = await serverApi.post(urlsEndpoint, { url })
    return response.data
  } catch (error) {
    console.error("Error generating short URL:", error)
    return null
  }
}

// Add more API methods as needed

export default serverApi
