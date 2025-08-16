package elevenlabs

type AccessInfo struct {
	IsCreator    bool   `json:"is_creator"`
	CreatorName  string `json:"creator_name"`
	CreatorEmail string `json:"creator_email"`
	Role         Role   `json:"role"`
}

type Role string

const (
	RoleAdmin  = "admin"
	RoleEditor = "editor"
	RoleViewer = "viewer"
)
