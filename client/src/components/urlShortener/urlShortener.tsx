"use client"
import { Button, useMantineTheme } from "@mantine/core"
import { useCallback, useEffect, useRef, useState } from "react"
import { Id, toast } from "react-toastify"

import { getShortUrlHash, isServerAvailable } from "@/api/serverApi"
import { useShortUrlContext } from "@/hooks/useShortUrlContext"
import { errorToast, isValidUrl, updateLoadingToast } from "@/utils/utils"
import TextInputField from "../TextInputField/TextInputField"
import styles from "./urlShortener.module.scss"

export const UrlShortener = () => {
  const [urlInput, setUrlInput] = useState<string>("")
  const [serverLoadingToastStatus, setServerLoadingToastStatus] = useState<
    "loading" | "success" | "error"
  >()
  const [serverLoadingToastId, setServerLoadingToastId] = useState<Id | null>(null)
  const serverLoadingToastStatusRef = useRef(serverLoadingToastStatus)
  const [isInputReady, setIsInputReady] = useState(false)
  const { shortUrl, setShortUrl } = useShortUrlContext()
  const theme = useMantineTheme()

  const inputErrMsg = urlInput && !isInputReady ? "Invalid URL" : ""

  useEffect(() => {
    const validateUrl = () => setIsInputReady(isValidUrl(urlInput))
    validateUrl()
  }, [urlInput])

  useEffect(() => {
    serverLoadingToastStatusRef.current = serverLoadingToastStatus
  }, [serverLoadingToastStatus])

  const handleServerCheck = useCallback(async () => {
    if (serverLoadingToastStatus === "loading") {
      setTimeout(() => {
        if (serverLoadingToastStatusRef.current !== "loading" || !serverLoadingToastId) return
        toast.update(serverLoadingToastId, {
          render: "Connecting to server... this may take a while (using free serverless tier)",
        })
      }, 4000)
    }

    const isAvailable = await isServerAvailable()
    setServerLoadingToastStatus(isAvailable ? "success" : "error")
  }, [serverLoadingToastId, serverLoadingToastStatus])

  useEffect(() => {
    if (!serverLoadingToastId) {
      setServerLoadingToastStatus("loading")
      setServerLoadingToastId(toast.loading("Connecting to server..."))
      return
    }

    handleServerCheck()
    if (serverLoadingToastStatus === "loading") return

    const toastConfig = {
      autoClose: 2000,
      isLoading: false,
    }
    if (serverLoadingToastStatus === "success") {
      toast.update(serverLoadingToastId, {
        render: "Connected to server",
        type: "success",
        ...toastConfig,
      })
      return
    }
    toast.update(serverLoadingToastId, {
      render: "Failed to connect to server",
      type: "error",
      ...toastConfig,
    })
  }, [handleServerCheck, serverLoadingToastId, serverLoadingToastStatus])

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
            disabled={!isInputReady || serverLoadingToastStatus !== "success"}
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
            disabled={(!urlInput && !shortUrl) || serverLoadingToastStatus !== "success"}
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
