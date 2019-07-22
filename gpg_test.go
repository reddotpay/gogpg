package gogpg_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/reddotpay/gogpg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func MustFileOpen(path string) *os.File {
	f, err := os.Open(path)

	if err != nil {
		panic(err)
	}
	return f
}

func TestEncyptDecrypt(t *testing.T) {
	var (
		require = require.New(t)
		assert  = assert.New(t)

		message = "hello world"

		passphrase = "password123"
		publickey  = MustFileOpen("testdata/demo-test.pub")
		secretkey  = MustFileOpen("testdata/demo-test.pvt")

		encryptedMessage []byte
		decryptedMessage []byte
		err              error
	)

	defer func() {
		publickey.Close()
		secretkey.Close()
	}()

	encryptedMessage, err = gogpg.Encrypt(publickey, []byte(message))
	require.NoError(err)

	decryptedMessage, err = gogpg.Decrypt(secretkey, passphrase, encryptedMessage)
	require.NoError(err)
	assert.Equal(string(decryptedMessage), message)

	_, err = gogpg.Encrypt(nil, []byte(message))
	assert.EqualError(err, "invalid argument")

	_, err = gogpg.Encrypt(bytes.NewReader([]byte{}), []byte(message))
	assert.EqualError(err, "openpgp: invalid argument: no armored data found")

	_, err = gogpg.Decrypt(secretkey, "wrongpassphrase", encryptedMessage)
	assert.EqualError(err, "openpgp: invalid argument: no armored data found")
}
