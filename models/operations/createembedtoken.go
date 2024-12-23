// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type CreateEmbedTokenRequestBody struct {
	LinkID string `json:"linkId"`
}

func (o *CreateEmbedTokenRequestBody) GetLinkID() string {
	if o == nil {
		return ""
	}
	return o.LinkID
}

// CreateEmbedTokenResponseBody - The created public embed token.
type CreateEmbedTokenResponseBody struct {
	PublicToken string `json:"publicToken"`
	Expires     string `json:"expires"`
}

func (o *CreateEmbedTokenResponseBody) GetPublicToken() string {
	if o == nil {
		return ""
	}
	return o.PublicToken
}

func (o *CreateEmbedTokenResponseBody) GetExpires() string {
	if o == nil {
		return ""
	}
	return o.Expires
}
