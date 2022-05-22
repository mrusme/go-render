package render

import (
  "os"
  "testing"
  "fmt"
  "encoding/json"
)

func TestListServices(t *testing.T) {
  r, err := New(os.Getenv("RENDER_API_TOKEN"))
  if err != nil {
    t.Fatalf("%s", err)
  }

  services, err := r.ListServices()
  if err != nil {
    t.Fatalf("%s", err)
  }

  for _, service := range services {
    b, _ := json.MarshalIndent(service, "", "    ")
    fmt.Printf("%s\n\n", string(b))
  }
}
