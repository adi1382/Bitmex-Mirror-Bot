package swagger

// Account Operations
type User struct {
	Id           NullInt          `json:"id,omitempty"`
	OwnerId      NullInt          `json:"ownerId,omitempty"`
	Firstname    NullString       `json:"firstname,omitempty"`
	Lastname     NullString       `json:"lastname,omitempty"`
	Username     NullString       `json:"username"`
	Email        NullString       `json:"email"`
	Phone        NullString       `json:"phone,omitempty"`
	Created      NullTime         `json:"created,omitempty"`
	LastUpdated  NullTime         `json:"lastUpdated,omitempty"`
	Preferences  *UserPreferences `json:"preferences,omitempty"`
	TFAEnabled   NullString       `json:"TFAEnabled,omitempty"`
	AffiliateID  NullString       `json:"affiliateID,omitempty"`
	PgpPubKey    NullString       `json:"pgpPubKey,omitempty"`
	Country      NullString       `json:"country,omitempty"`
	GeoipCountry NullString       `json:"geoipCountry,omitempty"`
	GeoipRegion  NullString       `json:"geoipRegion,omitempty"`
	Typ          NullString       `json:"typ,omitempty"`
}
