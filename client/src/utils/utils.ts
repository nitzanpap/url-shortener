import { Id, toast, ToastOptions } from "react-toastify"

export const isValidUrl = (str: string): boolean => {
  var pattern = new RegExp(
    "^(https?:\\/\\/)?" + // protocol
      "((([a-z\\d]([a-z\\d-]*[a-z\\d])*)\\.)+[a-z]{2,}|" + // domain name
      "((\\d{1,3}\\.){3}\\d{1,3}))" + // OR ip (v4) address
      "(\\:\\d+)?(\\/[-a-z\\d%_.~+]*)*" + // port and path
      "(\\?[;&a-z\\d%_.~+=-]*)?" + // query string
      "(\\#[-a-z\\d_]*)?$",
    "i"
  ) // fragment locator
  return !!pattern.test(str)
}

export const errorToast = (message: string) => {
  console.error(message)
  toast.error(message)
}

export const updateLoadingToast = (
  toastId: Id,
  message: string,
  toastType: ToastOptions["type"],
  autoClose?: ToastOptions["autoClose"]
) => {
  toast.update(toastId, {
    render: message,
    type: toastType,
    autoClose: autoClose,
    isLoading: false,
  })
}
