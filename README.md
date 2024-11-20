# xcrlf
CRLF Injector Testing , Fast Efficient , written in go with more than 100 payloads and advance detection mechanism .

# FileStructure


crlfi/
├── cmd/
│   └── crlfi/
│       └── main.go
├── internal/
│   ├── scanner/
│   │   ├── scanner.go           # Core scanning logic
│   │   ├── payloads.go          # Manage payloads
│   │   ├── resume.go            # Resume functionality
│   │   ├── delay.go             # Randomized delays
│   │   ├── useragents.go        # User-Agent rotation
│   └── utils/
│       └── utils.go             # Utilities for file handling
├── payloads/
│   └── default.txt              # Placeholder for your payloads
├── README.md
├── go.mod
└── go.sum
