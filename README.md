# Certify AWS updated Issuer

Intended to use with newest version of `github.com/aws/aws-sdk-go-v2`
which is not supported by aws implementation in `github.com/johanbrandhorst/certify`.

## How to use it

### Example using config

```go
import aws "github.com/mytheresa/certify-aws-issuer/pkg/tls"

func serverTLS(config aws.AWSConfig) (grpc.ServerOption, error) {
  ai, err := aws.NewAWSIssuerFromConfig(config)
  if err != nil {
    return grpc.EmptyServerOption{}, err
  }

  tlsConfig, err := aws.NewAWSTLSConfig(ai)
  if err != nil {
    return grpc.EmptyServerOption{}, err
  }
  
  return grpc.Creds(credentials.NewTLS(tlsConfig)), nil
}
```

`aws.AWSConfig` can be added to config struct (it have _yaml_ and _env_ annotations)

![](https://media.giphy.com/media/AhQev1suy32mWdFNcq/giphy.gif)
