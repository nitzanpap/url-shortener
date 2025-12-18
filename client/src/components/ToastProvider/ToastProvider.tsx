"use client";

import { ToastContainer } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

interface ToastProviderProps {
  children: React.ReactNode;
}

export default function ToastProvider({ children }: ToastProviderProps) {
  return (
    <>
      {children}
      <ToastContainer
        theme="dark"
        position="bottom-right"
        autoClose={2000}
        closeOnClick
        pauseOnHover
        draggable
      />
    </>
  );
}
