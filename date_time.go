package go_threema_bcast

import (
  "strings"
  "time"
)

const dateTimeLayout = "2006-01-02T15:04:05-0700"
// 2006-01-02T15:04:05Z07:00

type DateTime struct {
  time.Time
}

func (dt *DateTime) UnmarshalJSON(b []byte) (err error) {
  s := strings.Trim(string(b), `"`)
  dt.Time, err = time.Parse(dateTimeLayout, s)
  if err != nil {
    return err
  }
  return
}
