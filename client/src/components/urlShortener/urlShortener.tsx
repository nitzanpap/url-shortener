"use client"
import { generateShortUrl, isServerAvailable } from "@/api/serverApi"
import { isValidUrl } from "@/utils/utils"
import { useState, useEffect } from "react"
import styles from "./urlShortener.module.scss"

export const UrlShortener = () => {
  const [url, setUrl] = useState<string>("")

  const handleUrlInputChanged = (e: React.ChangeEvent<HTMLInputElement>) => {
    setUrl(e.target.value)
  }

  const handleGenerateButtonClicked = async () => {
    // check if the URL is valid

    if (!isValidUrl(url)) {
      console.log("Invalid URL")
      return
    }

    // send the URL to the server
    const generatedUrl = await generateShortUrl(url)
    if (generatedUrl) {
      console.log("Generated URL:", generatedUrl)
    } else {
      console.log("Failed to generate short URL")
    }
  }

  useEffect(() => {
    new Promise<void>(async (resolve, reject) => {
      console.log("Is server available?", await isServerAvailable())
    })
  }, [])

  return (
    <>
      <input
        type="url"
        name="url"
        id="url"
        className={styles.urlInput}
        placeholder="Enter URL"
        value={url}
        onChange={handleUrlInputChanged}
      />
      <button className={styles.urlButton} onClick={handleGenerateButtonClicked}>
        Generate
      </button>
    </>
  )
}
