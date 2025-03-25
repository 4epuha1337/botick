package tools

import (
  "strings"
  "strconv"
  "os"
)

func IsAdmin(id int64) bool {
  idAdmStr := os.Getenv("TELEGRAM_IDADM")
  idsAdm := strings.Split(idAdmStr, "?")
  for _, s := range idsAdm {
    idAdm, _ := strconv.Atoi(s)
    if id == int64(idAdm) {
      return true
    }
  }
  return false
}