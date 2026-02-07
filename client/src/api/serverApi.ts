import axios from "axios"
import { configurations } from "../configs/config"
import type { IGetShortUrlHashResponse } from "./serverApi.model"

const apiVersion = "v1"
const serverApi = axios.create({
  baseURL: `${configurations.envVars.serverBaseUrl}api/${apiVersion}/`,
})

const urlGroupEndpoint = "url/"

export const isServerAvailable = async () => {
  try {
    const response = await serverApi.get("/")
    return response.status === 200 && response.data.status === "ok"
  } catch (error) {
    console.error("Error checking server status:", error)
    return false
  }
}

export const getShortUrlHash = async (url: string): Promise<IGetShortUrlHashResponse | null> => {
  try {
    const response = await serverApi.post(urlGroupEndpoint, { url })
    return response.data
  } catch (error) {
    console.error("Error generating short URL:", error)
    return null
  }
}

export default serverApi
