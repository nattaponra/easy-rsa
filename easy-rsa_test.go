package rsa

import "testing"

func TestTimeConsuming(t *testing.T) {
	GetPublicKeyFromPEMFile("testdata/public.pem")

}
