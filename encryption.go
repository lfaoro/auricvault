package auricvault

// Vault-Managed Encryption Methods
//
// https://docs.auricvault.com/api-docs/#_vault_managed_encryption_methods

/*
Encrypt a plaintext value and store it using the passed-in token identifier. This allows you to migrate tokens you already have to the AuricVault service and maintain the same token identifier in your databases.

If no token is specified, one will be auto-generated.
*/
func (v *Vault) Encrypt(data, token string) (string, error) {
	v.request.Method = "encrypt"
	v.request.ID = 1
	if token != "" {
		v.request.Params[0].Token = token
	} else {
		v.request.Params[0].Last4 = ""
	}
	v.request.Params[0].UtcTimestamp = getTime()
	v.request.Params[0].PlaintextValue = data
	res, err := v.doRequest()
	if err != nil {
		return "", err
	}
	return res.Result.Token, nil
}

// ReEncrypt submit new plaintext data to be encrypted for an existing token.
func (v *Vault) ReEncrypt(data, token string) (string, error) {
	v.request.Method = "reencrypt"
	v.request.ID = 1
	v.request.Params[0].Token = token
	v.request.Params[0].UtcTimestamp = getTime()
	v.request.Params[0].PlaintextValue = data
	res, err := v.doRequest()
	if err != nil {
		return "", err
	}
	return res.Result.Token, nil
}

// Decrypt given a token retrieves the decrypted plaintext.
func (v *Vault) Decrypt(token string) (data string, err error) {
	v.request.Method = "decrypt"
	v.request.ID = 1
	v.request.Params[0].UtcTimestamp = getTime()
	v.request.Params[0].Token = token
	res, err := v.doRequest()
	if err != nil {
		return "", err
	}
	return res.Result.PlaintextValue, nil
}