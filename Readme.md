# go-prompt

Prompt for user data, quickly

## Available Methods

#### String(prompt string) string

Get a string

#### StringOrKeep(prompt string, val string) string

Get a string. If nothing is entered, return val instead

#### Password(prompt string) string

Get a password (string). Don't echo the output

#### PasswordOrKeep(prompt string, val string) string

Get a string. If nothing is entered, return val instead
