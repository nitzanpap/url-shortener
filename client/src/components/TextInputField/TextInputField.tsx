import { TextInput } from "@mantine/core"
import { FC } from "react"
import styles from "./TextInputField.module.scss"

interface TextInputFieldProps {
  type: string
  name: string
  id: string
  className: string
  placeholder: string
  required: boolean
  error: string
  value: string
  onChange: (e: React.ChangeEvent<HTMLInputElement>) => void
}

const TextInputField: FC<TextInputFieldProps> = ({
  type,
  name,
  id,
  className,
  placeholder,
  required,
  error,
  value,
  onChange,
}) => {
  return (
    <TextInput
      type={type}
      name={name}
      id={id}
      className={`${styles.textInput} ${className}`}
      placeholder={placeholder}
      required={required}
      error={error}
      value={value}
      onChange={onChange}
      width={"100%"}
    />
  )
}

export default TextInputField
