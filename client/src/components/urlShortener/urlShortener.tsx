"use client"
import { Button, useMantineTheme } from "@mantine/core"
import { useEffect, useRef, useState } from "react"
import { type Id, toast } from "react-toastify"

import { getShortUrlHash, isServerAvailable } from "@/api/serverApi"
import { useShortUrlContext } from "@/hooks/useShortUrlContext"
import { errorToast, isValidUrl, updateLoadingToast } from "@/utils/utils"
import TextInputField from "../TextInputField/TextInputField"
import styles from "./urlShortener.module.scss"

export const UrlShortener = () => {
  const [urlInput, setUrlInput] = useState<string>("")
  const [serverStatus, setServerStatus] = useState<"loading" | "success" | "error">("loading")
  const serverLoadingToastIdRef = useRef<Id | null>(null)
  const [isInputReady, setIsInputReady] = useState(false)
  const { shortUrl, setShortUrl } = useShortUrlContext()
  const theme = useMantineTheme()

  const inputErrMsg = urlInput && !isInputReady ? "Invalid URL" : ""

  useEffect(() => {
    const validateUrl = () => setIsInputReady(isValidUrl(urlInput))
    validateUrl()
  }, [urlInput])

  useEffect(() => {
    let isMounted = true
    const toastId = toast.loading("Connecting to server...")
    serverLoadingToastIdRef.current = toastId

    const timeoutId = setTimeout(() => {
      if (isMounted && serverLoadingToastIdRef.current) {
        toast.update(toastId, {
          render: "Connecting to server... this may take a while (using free serverless tier)",
        })
      }
    }, 4000)

    isServerAvailable().then((isAvailable) => {
      if (!isMounted) return

      const status = isAvailable ? "success" : "error"
      setServerStatus(status)

      toast.update(toastId, {
        render: isAvailable ? "Connected to server" : "Failed to connect to server",
        type: status,
        autoClose: 2000,
        isLoading: false,
      })
    })

    return () => {
      isMounted = false
      clearTimeout(timeoutId)
    }
  }, [])

  const handleUrlInputChanged = (e: React.ChangeEvent<HTMLInputElement>) =>
    setUrlInput(e.target.value)

  const handleGenerateButtonClicked = async () => {
    if (!isValidUrl(urlInput)) {
      errorToast("Invalid URL")
      return
    }

    setShortUrl("")
    const shortUrlHashResData = await getShortUrlHash(urlInput)
    const generatingUrlToastId = toast.loading("Generating short URL...")

    if (!shortUrlHashResData) {
      updateLoadingToast(generatingUrlToastId, "Failed to generate short URL", "error", 2000)
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
            disabled={!isInputReady || serverStatus !== "success"}
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
            disabled={(!urlInput && !shortUrl) || serverStatus !== "success"}
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
