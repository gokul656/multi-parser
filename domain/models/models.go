package models

// API
type PlotRequest []struct {
	PlotType string `json:"plot_type"`
	XAxis    string `json:"x-axis,omitempty"`
	YAxis    string `json:"y-axis,omitempty"`
	ZAxis    string `json:"z-axis,omitempty"`
}

type PlotRespose []struct {
	PlotType string   `json:"plot_type"`
	XAxis    string   `json:"x_axis,omitempty"`
	YAxis    string   `json:"y_axis,omitempty"`
	ZAxis    string   `json:"z_axis,omitempty"`
	Metadata Metadata `json:"metadata,omitempty"`
	Values   []string `json:"values,omitempty"`
}

type Metadata struct {
	Limit     int    `json:"limit,omitempty"`
	Offset    int    `json:"offset,omitempty"`
	OrderedBy string `json:"ordered_by,omitempty"`
	SortedBy  string `json:"sorted_by,omitempty"`
}
