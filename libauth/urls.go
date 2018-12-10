package libauth

type UrlProvider struct {
	protocol, host string
}

func NewUrlProvider(host string) *UrlProvider {
	u := new(UrlProvider)
	u.host = host
	u.protocol = "https://"
	return u
}

func (u *UrlProvider) LoginUriBase() string {
	return u.protocol + u.host + "/cgi-bin/srun_portal"
}

func (u *UrlProvider) OnlineCheckUriBase() string {
	return u.protocol + u.host + "/srun_portal_pc.php"
}

func (u *UrlProvider) ChallengeUriBase() string {
	return u.protocol + u.host + "/cgi-bin/get_challenge"
}
