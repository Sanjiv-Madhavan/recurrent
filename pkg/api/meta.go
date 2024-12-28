package api

type ServiceInstanceContext struct {
	Platform        string `json:"platform,omitempty"`
	GlobalAccountID string `json:"global_account_id,omitempty"`
	SubaccountID    string `json:"subaccount_id,omitempty"`
	Subdomain       string `json:"subdomain,omitempty"`
	InstanceName    string `json:"instance_name,omitempty"`
	Region          string `json:"region,omitempty"`
	Origin          string `json:"origin,omitempty"`
	SpaceID         string `json:"space_id,omitempty"`
	OrganizationID  string `json:"organization_id,omitempty"`
	ClusterID       string `json:"clusterid,omitempty"`
	Namespace       string `json:"namespace,omitempty"`
}
