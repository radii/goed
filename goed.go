package main

import "encoding/base64"
import "fmt"
import "golang.org/x/crypto/ed25519"

func b64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func main() {
	pubkey, privkey, err := ed25519.GenerateKey(nil)
	if err != nil {
		panic(err)
	}

	msg := "{\"timestamp\": 1510588145, \"payload\": \"Two if by sea\"}"
	msgbytes := []byte(msg)

	sig := ed25519.Sign(privkey, msgbytes)
	fmt.Println("  privkey:", len(privkey), b64(privkey))
	fmt.Println("   pubkey:", len(pubkey), b64(pubkey))
	fmt.Println("signature:", len(sig), b64(sig))

	result := ed25519.Verify(pubkey, msgbytes, sig)
	fmt.Println("verification result:", result)
}
