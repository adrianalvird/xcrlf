package scanner

import (
 "os"
)

func ResumeScan(resumeFile string, targets []string) []string {
 resumeSet := make(map[string]bool)
 file, err := os.Open(resumeFile)
 if err != nil {
  fmt.Printf("Error opening resume file: %v\n", err)
  return targets
 }
 defer file.Close()

 var scannedTarget string
 for {
  _, err := fmt.Fscanln(file, &scannedTarget)
  if err != nil {
   break
  }
  resumeSet[scannedTarget] = true
 }

 var remainingTargets []string
 for _, target := range targets {
  if !resumeSet[target] {
   remainingTargets = append(remainingTargets, target)
  }
 }
 return remainingTargets
}
