package scanner

// Hardcoded payloads for CRLF injection testing
func GetPayloads() []string {
 return []string{
  "%0d%0aTest",
  "%0d%0aHeader-Test: Injected",
  "%0d%0aContent-Length: 0",
  "%0d%0aX-Injected: True",
  "%0d%0aSet-Cookie: test=1",
  "%0d%0aLocation: /redirect",
 }
}
