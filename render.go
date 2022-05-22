package render

import (
  "encoding/json"
  "fmt"
  "net/http"
  "time"
)

type RenderClient struct {
  httpClient        http.Client
  token             string
}

type Deploy struct {
  ID                          string `json:"id"`
  Commit                      string `json:"commit"`
  Status                      string `json:"status"`
  CreatedAt                   time.Time `json:"createdAt"`
  UpdatedAt                   time.Time `json:"updatedAt"`
  FinishedAt                  time.Time `json:"finishedAt"`
}

type ServiceDetailsEnv struct {
  BuildCommand                string `json:"buildCommand"`
  StartCommand                string `json:"startCommand"`
  DockerCommand               string `json:"dockerCommand"`
  DockerContext               string `json:"dockerContext"`
  DockerfilePath              string `json:"dockerfilePath"`
}

type ServiceDetails struct {
  Disk                        map[string]interface{} `json:"disk"`
  Env                         string `json:"env"`
  EnvSpecificDetails          ServiceDetailsEnv `json:"envSpecificDetails"`
  HealthCheckPath             string `json:"healthCheckPath"`
  NumInstances                int `json:"numInstances"`
  OpenPorts                   map[string]interface{} `json:"openPorts"`
  PublishPath                 string `json:"publishPath"`
  ParentServer                map[string]interface{} `json:"parentServer"`
  Plan                        string `json:"plan"`
  PullRequestPreviewsEnabled  bool `json:"pullRequestPreviewsEnabled"`
  Region                      string `json:"region"`
  URL                         string `json:"url"`
  BuildCommand                string `json:"buildCommand"`
}

type Service struct {
  ID                          string `json:"id"`
  Type                        string `json:"type"`
  Repo                        string `json:"repo"`
  Name                        string `json:"name"`
  AutoDeploy                  bool `json:"autoDeploy"`
  Branch                      string `json:"branch"`
  CreatedAt                   time.Time `json:"createdAt"`
  UpdatedAt                   time.Time `json:"updatedAt"`
  NotifyOnFail                string `json:"notifyOnFail"`
  OwnerID                     string `json:"ownerId"`
  Slug                        string `json:"slug"`
  Suspenders                  []string `json:"suspenders"`
  Schedule                    string `json:"schedule"`
  LastSuccessfulRunAt         time.Time `json:"lastSuccessfulRunAt"`
  ServiceDetails              ServiceDetails `json:"serviceDetails"`
  Deploys                     []Deploy `json:"deploys"`
}

type ResponseItem struct {
  Cursor                      string `json:"cursor"`
  Service                     Service `json:"service,omitempty"`
  Deploy                      Deploy `json:"deploy,omitempty"`
}
type Response []ResponseItem

func New(token string) (*RenderClient, error) {
  renderClient := new(RenderClient)
  renderClient.httpClient = http.Client{Timeout: 10 * time.Second}
  renderClient.token = token
  

  return renderClient, nil
}

func (r *RenderClient) ListServices() ([]Service, error) {
  req, err := http.NewRequest(
    "GET", 
    "https://api.render.com/v1/services",
    nil,
  )
  if err != nil {
    return []Service{}, err
  }

  req.Header.Add(
    "Authorization", 
    fmt.Sprintf("Bearer %s", r.token),
  )
  resp, err := r.httpClient.Do(req)
  if err != nil {
    return []Service{}, err
  }

  defer resp.Body.Close()
  var response Response
  json.NewDecoder(resp.Body).Decode(&response)

  var services []Service
  for _, responseItem := range response {
    deploys, _ := r.ListDeploys(responseItem.Service.ID)

    responseItem.Service.Deploys = deploys
    services = append(services, responseItem.Service)
  }
  return services, nil
}

func (r *RenderClient) ListDeploys(serviceID string) ([]Deploy, error) {
  req, err := http.NewRequest(
    "GET", 
    fmt.Sprintf("https://api.render.com/v1/services/%s/deploys", serviceID),
    nil,
  )
  if err != nil {
    return []Deploy{}, err
  }

  req.Header.Add(
    "Authorization", 
    fmt.Sprintf("Bearer %s", r.token),
  )
  resp, err := r.httpClient.Do(req)
  if err != nil {
    return []Deploy{}, err
  }

  defer resp.Body.Close()
  var response Response
  json.NewDecoder(resp.Body).Decode(&response)

  var deploys []Deploy
  for _, responseItem := range response {
    deploys = append(deploys, responseItem.Deploy)
  }
  return deploys, nil
}

