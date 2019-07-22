package gogpg

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"

	"golang.org/x/crypto/openpgp"
)

// Encrypt encrypts the message using the public key file
func Encrypt(publickey io.Reader, text []byte) ([]byte, error) {
	if publickey == nil {
		return nil, errors.New("invalid argument")
	}

	var (
		buf       = new(bytes.Buffer)
		key, err  = openpgp.ReadArmoredKeyRing(publickey)
		plaintext io.WriteCloser
	)

	if err != nil {
		return nil, err
	}

	if plaintext, err = openpgp.Encrypt(buf, key, nil, nil, nil); err != nil {
		return nil, err
	}

	if _, err = plaintext.Write(text); err != nil {
		return nil, err
	}

	plaintext.Close()

	return ioutil.ReadAll(buf)
}

// Decrypt decrypts the message using secret key file and a passphrase
func Decrypt(secretkey io.Reader, passphrase string, text []byte) ([]byte, error) {
	if secretkey == nil {
		return nil, errors.New("invalid argument")
	}

	var (
		keys, err   = openpgp.ReadArmoredKeyRing(secretkey)
		passphraseb = []byte(passphrase)
		md          *openpgp.MessageDetails
	)

	if err != nil {
		return nil, err
	}

	key := keys[0]
	key.PrivateKey.Decrypt(passphraseb)

	for _, subkey := range key.Subkeys {
		subkey.PrivateKey.Decrypt(passphraseb)
	}

	if md, err = openpgp.ReadMessage(bytes.NewBuffer(text), keys, nil, nil); err != nil {
		return nil, err
	}

	return ioutil.ReadAll(md.UnverifiedBody)
}
