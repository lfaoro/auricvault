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
	if token != "" {
		v.request.Params[0].Token = token
	} else {
		v.request.Params[0].Last4 = ""
	}
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
	v.request.Params[0].Token = token
	res, err := v.doRequest()
	if err != nil {
		return "", err
	}
	log.Debug("response: ", res)
	return res.Result.PlaintextValue, nil
}
