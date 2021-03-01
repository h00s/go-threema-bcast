package go_threema_bcast

import (
  "fmt"
  "os"
  "testing"
)

func TestGetProfile(t *testing.T) {
  c := NewClient(os.Getenv("API_KEY"))
  p, err := c.GetProfile()
  if err != nil {
    t.Errorf("Error while getting profile: %v", err)
  }
  fmt.Printf("%+v\n", p)
}