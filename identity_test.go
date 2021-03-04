package go_threema_bcast

import (
  "fmt"
  "os"
  "testing"
)

func TestGetIdentities(t *testing.T) {
  c := NewClient(os.Getenv("API_KEY"))
  p, err := c.GetIdentities()
  if err != nil {
    t.Errorf("Error while getting identities: %v", err)
  }
  fmt.Printf("%+v\n", p)
}
