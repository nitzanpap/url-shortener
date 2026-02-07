import axios from "axios"
import { configurations } from "../configs/config"
import type { IGetOriginalUrlResponse } from "./serverApi.model"

const serverApi = axios.create({
  baseURL: `${configurations.envVars.serverBaseUrl}`,
  withCredentials: true,
})

const shortUrlEndpoint = "s/"

export const getOriginalUrl = async (
  shortUrlHash: string
): Promise<IGetOriginalUrlResponse | null> => {
  try {
    const response = await serverApi.get(`${shortUrlEndpoint}${shortUrlHash}`)
    return response.data
  } catch (error) {
    console.error("Error getting original URL:", error)
    return null
  }
}

export default serverApi
