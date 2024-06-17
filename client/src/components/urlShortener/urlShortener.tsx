"use client"
import { getShortUrlHash, isServerAvailable } from "@/api/serverApi"
import { isValidUrl } from "@/utils/utils"
import { Button, TextInput, useMantineTheme } from "@mantine/core"
import { useEffect, useState } from "react"
import styles from "./urlShortener.module.scss"

export const UrlShortener = () => {
  const [urlInput, setUrlInput] = useState<string>("")
  const [isInputReady, setIsInputReady] = useState(false)
  const [shortUrl, setShortUrl] = useState<string>("")
  const inputErrMsg = !isInputReady && urlInput && "Invalid URL"
  const theme = useMantineTheme()

  const handleUrlInputChanged = (e: React.ChangeEvent<HTMLInputElement>) => {
    setUrlInput(e.target.value)
  }

  useEffect(() => {
    if (!isValidUrl(urlInput)) {
      setIsInputReady(false)
      return
    }
    setIsInputReady(true)
  }, [urlInput])

  const handleGenerateButtonClicked = async () => {
    if (!isValidUrl(urlInput)) {
      console.log("Invalid URL")
      return
    }
    // send the URL to the server
    const shortUrlHashResData = await getShortUrlHash(urlInput)
    if (!shortUrlHashResData) {
      console.log("Failed to generate short URL")
      return
    }
    const { obfuscatedShortenedUrl: shortUrlHash } = shortUrlHashResData
    setShortUrl(`${window.location.origin}/s/${shortUrlHash}`)
  }

  useEffect(() => {
    new Promise<void>(async (resolve, reject) => {
      console.log("Is server available?", await isServerAvailable())
    })
  }, [])

  return (
    <section className={styles.urlShortenerContainer}>
      <TextInput
        type="url"
        name="url"
        id="url"
        className={styles.urlInput}
        placeholder="Enter URL"
        required
        error={inputErrMsg}
        value={urlInput}
        onChange={handleUrlInputChanged}
      />
      <Button
        className={styles.urlButton}
        onClick={handleGenerateButtonClicked}
        disabled={!isInputReady}
        variant="gradient"
        gradient={{ from: theme.primaryColor, to: theme.colors.red[8] }}
      >
        Generate
      </Button>
      {shortUrl && (
        <a href={shortUrl} className={styles.shortUrl}>
          {shortUrl}
        </a>
      )}
    </section>
  )
}
