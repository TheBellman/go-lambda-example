# go-lambda-example

Very simple Go code used to demonstrate the lambda pipeline. This was built and tested with go1.15:

```
% go version
go version go1.15 darwin/amd64
```

## Testing
Testing is straight forward:

```
% go test
PASS
ok  	parttimepolymath.net/lambda	0.127s
```

## Building
A simple build to the current directory is also straightforward, although it's not much use running locally:

```
% go build
% ./lambda
2020/11/14 10:51:51 expected AWS Lambda environment variables [_LAMBDA_SERVER_PORT AWS_LAMBDA_RUNTIME_API] are not defined
```

## License
Copyright 2020 Little Dog Digital

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
