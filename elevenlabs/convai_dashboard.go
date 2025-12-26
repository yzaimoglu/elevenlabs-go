package elevenlabs

import (
	"context"
	"errors"
)

type ConvaiDashboardAPI interface {
	GetDashboardSettings(ctx context.Context) (DashboardSettingsResp, error)
	UpdateDashboardSettings(ctx context.Context, req *UpdateDashboardSettingsReq) (DashboardSettingsResp, error)
}

// DashboardChartType represents the type of dashboard chart
type DashboardChartType string

const (
	DashboardChartTypeCallSuccess    DashboardChartType = "call_success"
	DashboardChartTypeCriteria       DashboardChartType = "criteria"
	DashboardChartTypeDataCollection DashboardChartType = "data_collection"
)

// DashboardChart represents a dashboard chart configuration
type DashboardChart struct {
	Name             string             `json:"name"`
	Type             DashboardChartType `json:"type"`
	CriteriaId       string             `json:"criteria_id,omitempty"`
	DataCollectionId string             `json:"data_collection_id,omitempty"`
}

// DashboardSettingsResp represents the response for dashboard settings
type DashboardSettingsResp struct {
	Charts []DashboardChart `json:"charts,omitempty"`
}

// GetDashboardSettings retrieves Convai dashboard settings for the workspace.
// https://elevenlabs.io/docs/api-reference/workspace/dashboard/get
func (c *Client) GetDashboardSettings(ctx context.Context) (DashboardSettingsResp, error) {
	body, err := c.get(ctx, "/convai/settings/dashboard")
	if err != nil {
		return DashboardSettingsResp{}, err
	}

	var resp DashboardSettingsResp
	if err := c.parseResponse(body, &resp); err != nil {
		return DashboardSettingsResp{}, err
	}

	return resp, nil
}

// UpdateDashboardSettingsReq represents the request for updating dashboard settings
type UpdateDashboardSettingsReq struct {
	Charts []DashboardChart `json:"charts,omitempty"`
}

func NewUpdateDashboardSettingsReq() *UpdateDashboardSettingsReq {
	return &UpdateDashboardSettingsReq{}
}

// UpdateDashboardSettings updates Convai dashboard settings for the workspace.
// https://elevenlabs.io/docs/api-reference/workspace/dashboard/update
func (c *Client) UpdateDashboardSettings(ctx context.Context, req *UpdateDashboardSettingsReq) (DashboardSettingsResp, error) {
	if req == nil {
		return DashboardSettingsResp{}, errors.New("request is nil")
	}

	body, err := c.patch(ctx, "/convai/settings/dashboard", req)
	if err != nil {
		return DashboardSettingsResp{}, err
	}

	var resp DashboardSettingsResp
	if err := c.parseResponse(body, &resp); err != nil {
		return DashboardSettingsResp{}, err
	}

	return resp, nil
}
