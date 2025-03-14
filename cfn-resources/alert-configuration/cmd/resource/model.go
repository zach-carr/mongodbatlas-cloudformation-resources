// Code generated by 'cfn generate', changes will be undone by the next invocation. DO NOT EDIT.
// Updates to this type are made my editing the schema file and executing the 'generate' command.
package resource

// Model is autogenerated from the json schema
type Model struct {
	Id              *string               `json:",omitempty"`
	Profile         *string               `json:",omitempty"`
	Created         *string               `json:",omitempty"`
	Enabled         *bool                 `json:",omitempty"`
	EventTypeName   *string               `json:",omitempty"`
	ProjectId       *string               `json:",omitempty"`
	Matchers        []Matcher             `json:",omitempty"`
	MetricThreshold *MetricThresholdView  `json:",omitempty"`
	Notifications   []NotificationView    `json:",omitempty"`
	Threshold       *IntegerThresholdView `json:",omitempty"`
	TypeName        *string               `json:",omitempty"`
	Updated         *string               `json:",omitempty"`
}

// Matcher is autogenerated from the json schema
type Matcher struct {
	FieldName *string `json:",omitempty"`
	Operator  *string `json:",omitempty"`
	Value     *string `json:",omitempty"`
}

// MetricThresholdView is autogenerated from the json schema
type MetricThresholdView struct {
	MetricName *string  `json:",omitempty"`
	Mode       *string  `json:",omitempty"`
	Operator   *string  `json:",omitempty"`
	Threshold  *float64 `json:",omitempty"`
	Units      *string  `json:",omitempty"`
}

// NotificationView is autogenerated from the json schema
type NotificationView struct {
	ApiToken                 *string  `json:",omitempty"`
	ChannelName              *string  `json:",omitempty"`
	DatadogApiKey            *string  `json:",omitempty"`
	DatadogRegion            *string  `json:",omitempty"`
	DelayMin                 *float64 `json:",omitempty"`
	EmailAddress             *string  `json:",omitempty"`
	EmailEnabled             *bool    `json:",omitempty"`
	FlowName                 *string  `json:",omitempty"`
	FlowdockApiToken         *string  `json:",omitempty"`
	IntervalMin              *float64 `json:",omitempty"`
	MicrosoftTeamsWebhookUrl *string  `json:",omitempty"`
	MobileNumber             *string  `json:",omitempty"`
	NotificationToken        *string  `json:",omitempty"`
	OpsGenieApiKey           *string  `json:",omitempty"`
	OpsGenieRegion           *string  `json:",omitempty"`
	OrgName                  *string  `json:",omitempty"`
	Roles                    []string `json:",omitempty"`
	RoomName                 *string  `json:",omitempty"`
	ServiceKey               *string  `json:",omitempty"`
	Severity                 *string  `json:",omitempty"`
	SmsEnabled               *bool    `json:",omitempty"`
	TeamId                   *string  `json:",omitempty"`
	TeamName                 *string  `json:",omitempty"`
	TypeName                 *string  `json:",omitempty"`
	Username                 *string  `json:",omitempty"`
	VictorOpsApiKey          *string  `json:",omitempty"`
	VictorOpsRoutingKey      *string  `json:",omitempty"`
	WebhookSecret            *string  `json:",omitempty"`
	WebhookUrl               *string  `json:",omitempty"`
}

// IntegerThresholdView is autogenerated from the json schema
type IntegerThresholdView struct {
	Operator  *string  `json:",omitempty"`
	Threshold *float64 `json:",omitempty"`
	Units     *string  `json:",omitempty"`
}
