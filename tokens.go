package auricvault

// Token Management Methods
//
// https://docs.auricvault.com/api-docs/#_token_management_methods

// Delete returns the same message for both a not-found token and a token that exists,
// but to which you do not have permission. This ensures the existence of the token does not
// leakto a third party that should not have access to the data.
func (v *Vault) Delete(token string) error {
	v.request.Method = "delete_token"
	v.request.Params[0].Token = token
	_, err := v.doRequest()
	if err != nil {
		return err
	}
	return nil
}

// Info retrieves information about a token. Useful for finding out if a token exists
// in the system without needing to retrieve the actual data.
func (v *Vault) Info(token string) (*Result, error) {
	v.request.Method = "token_info"
	v.request.Params[0].Token = token
	res, err := v.doRequest()
	if err != nil {
		return nil, err
	}
	return &res.Result, nil
}

// Touch method is similar to the TokenInfo method except that it does update
// the token last accessed date time stamp. This method is used to reset the start of
// the retention period to the current date/time.
//
// Touching a non-existing token results in an error message and lastActionSucceeded of 0.
func (v *Vault) Touch(token string) (*Result, error) {
	v.request.Method = "touch_token"
	v.request.Params[0].Token = token
	res, err := v.doRequest()
	if err != nil {
		return nil, err
	}
	return &res.Result, nil
}
