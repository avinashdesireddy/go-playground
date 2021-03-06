package main

import (
  "fmt"
  "go.etcd.io/etcd/clientv3"
  "context"
  "time"
)

func main() {
  cli, err := clientv3.New(clientv3.Config{
      Endpoints: []string{"localhost:2379"},
      DialTimeout: 5 * time.Second,
  })
  if err != nil {
    fmt.Println(err)
  }
  defer cli.Close()

  ctx, cancel := context.WithTimeout(context.Background(), 10)
  response, err := cli.Put(ctx, "sample_key", "sample_value")
  cancel()
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(response)
}
