package patrick

// TODO: optimize cbor storage
type Meta struct {
	Ref        string     `json:"ref" cbor:"4,keyasint"`
	Before     string     `json:"before" cbor:"8,keyasint"`
	After      string     `json:"after" cbor:"16,keyasint"`
	HeadCommit HeadCommit `json:"head_commit" cbor:"32,keyasint"`
	Repository Repository `cbor:"64,keyasint"`
}

type HeadCommit struct {
	ID string `json:"id" cbor:"33,keyasint"`
}

type Repository struct {
	ID         int    `json:"id" cbor:"65,keyasint"`
	Provider   string `json:"provider" cbor:"66,keyasint"`
	SSHURL     string `json:"ssh_url" cbor:"67,keyasint"`
	Branch     string `json:"default_branch" cbor:"68,keyasint"`
	MainBranch string `json:"master_branch" cbor:"69,keyasint"`
}
