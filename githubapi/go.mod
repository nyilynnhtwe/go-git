module example.com/githubapi

go 1.20

replace example.com/getpass => ../getpass

require (
	example.com/getpass v0.0.0-00010101000000-000000000000
	github.com/google/go-github v17.0.0+incompatible
	golang.org/x/oauth2 v0.8.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	golang.org/x/crypto v0.9.0 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/term v0.8.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)
