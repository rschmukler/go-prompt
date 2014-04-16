package prompt

import (
  "fmt"
  "github.com/howeyc/gopass"
)

func Password(prompt string) string {
  fmt.Printf(prompt + ": ")
  return string(gopass.GetPasswd())
}

func PasswordOrKeep(prompt, val string) string {
  if len(val) > 0 {
    fmt.Printf(prompt + ": (enter to keep) ")
  } else {
    fmt.Printf(prompt + ": ")
  }

  tmpPassword := string(gopass.GetPasswd())
  if len(tmpPassword) > 0 {
    return tmpPassword
  } else {
    return val
  }
}

func String(prompt string) string {
  var result string;

  fmt.Printf(prompt + ": ")
  fmt.Scanln(&result)
  return result
}

func StringOrKeep(prompt, val string) string {
  var result string

  if len(val) > 0 {
    fmt.Printf(prompt + ": (%s) ", val)
  } else {
    fmt.Printf(prompt + ": ")
  }

  fmt.Scanln(&result)

  if len(result) > 0 {
    return result
  } else {
    return val
  }
}
