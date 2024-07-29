"use client"
import { getShortUrlHash, isServerAvailable } from "@/api/serverApi"
import { useShortUrlContext } from "@/hooks/useShortUrlContext"
import { errorToast, isValidUrl, updateLoadingToast } from "@/utils/utils"
import { Button, useMantineTheme } from "@mantine/core"
import { useEffect, useState } from "react"
import { toast } from "react-toastify"
import TextInputField from "../TextInputField/TextInputField"
import styles from "./urlShortener.module.scss"

export const UrlShortener = () => {
  const [urlInput, setUrlInput] = useState<string>("")
  const [serverLoadingToastStatus, setServerLoadingToastStatus] = useState<
    "loading" | "success" | "error"
  >("loading")
  const [isInputReady, setIsInputReady] = useState(false)
  const { shortUrl, setShortUrl } = useShortUrlContext()
  const inputErrMsg = !isInputReady && urlInput ? "Invalid URL" : ""
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
      errorToast("Invalid URL")
      return
    }
    // send the URL to the server
    const shortUrlHashResData = await getShortUrlHash(urlInput)
    const generatingUrlToastId = toast.loading("Generating short URL...")
    if (!shortUrlHashResData) {
      errorToast("Failed to generate short URL")
      return
    }
    toast.update(generatingUrlToastId, {
      render: "Short URL generated",
      type: "success",
      autoClose: 2000,
      isLoading: false,
    })
    const { obfuscatedShortenedUrl: shortUrlHash } = shortUrlHashResData
    setShortUrl(`${window.location.origin}/s/${shortUrlHash}`)
  }

  useEffect(() => {
    const serverLoadingToastId = toast.loading("Connecting to server...")
    setTimeout(() => {
      if (!toast.isActive(serverLoadingToastId) || serverLoadingToastStatus !== "loading") return
      toast.update(serverLoadingToastId, {
        render: "Connecting to server... this may take a while (using free serverless tier)",
      })
    }, 4000)
    isServerAvailable().then((isAvailable) => {
      if (!isAvailable) {
        updateLoadingToast(serverLoadingToastId, "Server is not available", "error", false)
        setServerLoadingToastStatus("error")
        return
      }
      updateLoadingToast(serverLoadingToastId, "Server connected", "success", 2000)
      setServerLoadingToastStatus("success")
    })
  }, [])

  return (
    <section className={styles.urlShortenerSection}>
      <div className={styles.urlShortenerContainer}>
        <TextInputField
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
        <div className={styles.buttonsContainer}>
          <Button
            className={styles.urlButton}
            onClick={handleGenerateButtonClicked}
            disabled={!isInputReady}
            variant="gradient"
            gradient={{ from: theme.primaryColor, to: theme.colors.blue[4] }}
          >
            Generate
          </Button>
          <Button
            className={styles.clearButton}
            onClick={() => {
              setUrlInput("")
              setShortUrl("")
            }}
            disabled={!urlInput && !shortUrl}
            variant="gradient"
            gradient={{ from: theme.colors.red[8], to: theme.colors.red[4] }}
          >
            Clear
          </Button>
        </div>
      </div>
    </section>
  )
}
