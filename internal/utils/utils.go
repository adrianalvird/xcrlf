package utils

import "os"

func WriteToFile(filepath, content string) error {
 file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
 if err != nil {
  return err
 }
 defer file.Close()

 _, err = file.WriteString(content + "\n")
 return err
}
