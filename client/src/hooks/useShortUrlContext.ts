import { createContext, useContext } from "react";

export const ShortUrlContext = createContext({ shortUrl: "", setShortUrl: (_: string) => {} });

export const useShortUrlContext = () => {
  const context = useContext(ShortUrlContext);
  if (!context) {
    throw new Error("useShortUrlContext must be used within a ShortUrlProvider");
  }
  return context;
};
